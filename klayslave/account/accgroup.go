package account

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"sync"

	"github.com/kaiachain/kaia/client"
)

// AccList defines the enum for accList
type AccList int

const (
	AccListForSignedTx AccList = iota
	AccListForUnsignedTx
	AccListForNewAccounts
	AccListForGaslessRevertTx
	AccListForGaslessApproveTx
	AccListEnd
)

// TestContract defines the enum for TestContract
type TestContract int

const (
	ContractErc20 TestContract = iota
	ContractErc721
	ContractStorageTrie
	ContractGeneral
	ContractGaslessToken
	ContractWKaia
	ContractUniswapV2Factory
	ContractUniswapV2Router
	ContractGaslessSwapRouter
	ContractCounterForTestAuction
	ContractAuctionFeeVault
	ContractAuctionDepositVault
	ContractAuctionEntryPoint
	ContractEnd
)

type AccLoader func(*AccGroup)

type AccGroup struct {
	containsUnsignedAccGrp bool

	accLists  [][]*Account
	contracts []*Account
}

func NewAccGroup(chainId *big.Int, gasPrice *big.Int, baseFee *big.Int, contains bool) *AccGroup {
	SetChainID(chainId)
	SetGasPrice(gasPrice)
	SetBaseFee(baseFee)

	return &AccGroup{
		containsUnsignedAccGrp: contains,
		accLists:               make([][]*Account, AccListEnd),
		contracts:              make([]*Account, ContractEnd),
	}
}
func (a *AccGroup) Load(loader AccLoader) { loader(a) }

func (a *AccGroup) GetTestContractByName(t TestContract) *Account { return a.contracts[t] }
func (a *AccGroup) GetAccListByName(t AccList) []*Account         { return a.accLists[t] }

func (a *AccGroup) SetTestContractByName(c *Account, t TestContract) { a.contracts[t] = c }
func (a *AccGroup) SetAccListByName(accs []*Account, t AccList) {
	for _, acc := range accs {
		a.accLists[t] = append(a.accLists[t], acc)
	}
}
func (a *AccGroup) AddAccToListByName(acc *Account, t AccList) {
	a.accLists[t] = append(a.accLists[t], acc)
}
func (a *AccGroup) CreateAccountsPerAccGrp(nUserForSignedTx int, nUserForUnsignedTx int, nUserForNewAccounts int, nUserForGaslessRevertTx int, nUserForGaslessApproveTx int, tcStrList []string, gEndpoint string) {
	for idx, nUser := range []int{nUserForSignedTx, nUserForUnsignedTx, nUserForNewAccounts, nUserForGaslessRevertTx, nUserForGaslessApproveTx} {
		println(idx, " Account Group Preparation...")
		for i := 0; i < nUser; i++ {
			account := NewAccount(i)
			a.AddAccToListByName(account, AccList(idx))
			fmt.Printf("%v\n", account.address.String())
		}
	}

	// Unlock AccGrpForUnsignedTx if needed
	for _, tcName := range tcStrList {
		if tcName != "transferUnsignedTx" {
			continue
		}
		// If at least one task needs unlocking, unlock the accGrp for unsignedTx.
		for _, acc := range a.GetAccListByName(AccListForUnsignedTx) {
			acc.ImportUnLockAccount(gEndpoint)
		}
		break
	}
}

func (a *AccGroup) SetAccGrpByActivePercent(activeUserPercent int) {
	for i, accGrp := range a.accLists {
		numActiveAccGrpForSignedTx := len(accGrp) * activeUserPercent / 100
		if numActiveAccGrpForSignedTx == 0 {
			a.accLists[i] = nil
			continue
		}

		a.accLists[i] = accGrp[:numActiveAccGrpForSignedTx]
	}
}

func (a *AccGroup) GetValidAccGrp() []*Account {
	var accGrp []*Account
	for _, acc := range a.GetAccListByName(AccListForSignedTx) {
		accGrp = append(accGrp, acc)
	}
	//if !a.cfg.InTheTcList("transferUnsignedTx") {
	if !a.containsUnsignedAccGrp {
		return accGrp
	}
	for _, acc := range a.GetAccListByName(AccListForUnsignedTx) {
		accGrp = append(accGrp, acc)
	}
	return accGrp
}

func (a *AccGroup) DeployTestContracts(tcList []string, targetTxTypeList []string, localReservoir *Account, gCli *client.Client, chargeValue *big.Int, maxConcurrency int) {
	inTheTcList := func(testNames []string) bool {
		for _, tcName := range tcList {
			for _, target := range testNames {
				if tcName == target {
					return true
				}
			}
		}
		return false
	}

	inTheTargetTxTypeList := func(targetTxTypes []string) bool {
		for _, targetTxType := range targetTxTypeList {
			for _, target := range targetTxTypes {
				if targetTxType == target {
					return true
				}
			}
		}
		return false
	}

	for idx, info := range TestContractInfos {
		testContractType := TestContract(idx)
		if testContractType != ContractGeneral && !inTheTcList(info.testNames) && !inTheTargetTxTypeList(info.auctionTargetTxTypeList) {
			continue
		}

		if info.ShouldDeploy(gCli, info.deployer) {
			if info.deployer == nil {
				info.deployer = localReservoir
			}
			localReservoir.TransferSignedTxWithGuaranteeRetry(gCli, info.deployer, chargeValue)
			a.contracts[idx] = info.deployer.SmartContractDeployWithGuaranteeRetry(gCli, info.GetBytecodeWithConstructorParam(info.Bytecode, a.contracts, info.deployer), info.contractName, true)
		}

		a.contracts[idx] = NewKaiaAccountWithAddr(0, info.GetAddress(gCli, info.deployer))

		// additional work - erc20 token charging or erc721 minting
		if TestContract(idx) == ContractErc20 {
			log.Printf("Start erc20 token charging to the test account group")
			TestContractInfos[ContractErc20].deployer.SmartContractExecutionWithGuaranteeRetry(gCli, a.contracts[ContractErc20], nil, TestContractInfos[ContractErc20].GenData(localReservoir.address, big.NewInt(1e11)))
			ConcurrentTransactionSend(a.GetValidAccGrp(), maxConcurrency, func(acc *Account) {
				localReservoir.SmartContractExecutionWithGuaranteeRetry(gCli, a.contracts[ContractErc20], nil, TestContractInfos[ContractErc20].GenData(acc.address, big.NewInt(1e4)))
			})
		} else if TestContract(idx) == ContractErc721 {
			log.Printf("Start erc721 nft minting to the test account group(similar to erc20 token charging)")
			localReservoir.MintERC721ToTestAccounts(gCli, a.GetValidAccGrp(), a.GetTestContractByName(ContractErc721).GetAddress(), 5)
		} else if TestContract(idx) == ContractGaslessToken && (inTheTcList([]string{"gaslessTransactionTC", "gaslessOnlyApproveTC"}) || inTheTargetTxTypeList([]string{"GAA", "GAS"})) {
			log.Printf("Start gasless test token charging to the test account group")
			lenValidAccGrp := big.NewInt(int64(len(a.GetValidAccGrp())))
			lenGaslessApproveAccGrp := big.NewInt(int64(len(a.GetAccListByName(AccListForGaslessApproveTx))))
			totalChargeValue := new(big.Int).Mul(chargeValue, new(big.Int).Add(lenValidAccGrp, lenGaslessApproveAccGrp))
			// ContractGaslessToken's GenData generate data of approve. So can use ERC20's genData for transfer.
			TestContractInfos[ContractGaslessToken].deployer.SmartContractExecutionWithGuaranteeRetry(gCli, a.contracts[ContractGaslessToken], nil, TestContractInfos[ContractErc20].GenData(localReservoir.address, totalChargeValue))

			// accounts(validAccGrp + gaslessApproveAccGrp) should be charged.
			accounts := a.GetValidAccGrp()
			accounts = append(accounts, a.GetAccListByName(AccListForGaslessApproveTx)...)
			ConcurrentTransactionSend(accounts, maxConcurrency, func(acc *Account) {
				localReservoir.SmartContractExecutionWithGuaranteeRetry(gCli, a.contracts[ContractGaslessToken], nil, TestContractInfos[ContractErc20].GenData(acc.address, chargeValue))
			})
		}
	}
}

type AccountSet struct {
	accounts        []*Account
	mu              sync.Mutex
	roundRobinIndex int
}

func NewAccountSet(accounts []*Account) *AccountSet {
	return &AccountSet{
		accounts:        accounts,
		mu:              sync.Mutex{},
		roundRobinIndex: 0,
	}
}

func (a *AccountSet) Len() int {
	return len(a.accounts)
}

func (a *AccountSet) Add(acc *Account) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.accounts = append(a.accounts, acc)
}

func (a *AccountSet) GetAccountRandomly() *Account {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.accounts[rand.Int()%a.Len()]
}

func (a *AccountSet) GetAccountRoundRobin() *Account {
	a.mu.Lock()
	defer a.mu.Unlock()
	acc := a.accounts[a.roundRobinIndex]
	a.roundRobinIndex++
	if a.roundRobinIndex >= a.Len() {
		a.roundRobinIndex = 0
	}
	return acc
}
