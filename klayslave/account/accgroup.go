package account

import (
	"fmt"
	"math/big"
)

// AccList defines the enum for accList
type AccList int

const (
	AccListForSignedTx AccList = iota
	AccListForUnsignedTx
	AccListForNewAccounts
	AccListEnd
)

// TestContract defines the enum for TestContract
type TestContract int

const (
	ContractErc20 TestContract = iota
	ContractErc721
	ContractStorageTrie
	ContractGeneral
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
func (a *AccGroup) CreateAccountsPerAccGrp(nUserForSignedTx int, nUserForUnsignedTx int, nUserForNewAccounts int, tcStrList []string, gEndpoint string) {
	for idx, nUser := range []int{nUserForSignedTx, nUserForUnsignedTx, nUserForNewAccounts} {
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
	for _, accGrp := range a.accLists {
		numActiveAccGrpForSignedTx := len(accGrp) * activeUserPercent / 100
		// Not to assign 0 account for some cases.
		if numActiveAccGrpForSignedTx == 0 {
			numActiveAccGrpForSignedTx = 1
		}

		accGrp = accGrp[:numActiveAccGrpForSignedTx]
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

//func (a *AccGroup) ChargeTestTokens() {
//	numChargedAcc := 0
//	lastFailedNum := 0
//	for _, recipientAccount := range a.cfg.GetAllAccGrp() {
//		for {
//			_, _, err := tokenChargeFn(false, cfg.GetGCli(), tokenContractAddr, recipientAccount, tokenChargeAmount)
//			if err == nil {
//				break // Success, move to next account.
//			}
//			numChargedAcc, lastFailedNum = estimateRemainingTime(cfg.GetAllAccGrp(), numChargedAcc, lastFailedNum)
//		}
//		numChargedAcc++
//	}
//}
//func (a *AccGroup) PrepareTestContract(t TestContract) {
//	if !a.cfg.InTheTcList(t) {
//		return
//	}
//
//	// Dedicated and fixed private key used to deploy a smart contract for ERC20 and ERC721 value transfer and storage trie writeperformance test.
//	privateKeyStr := []string{
//		"eb2c84d41c639178ff26a81f488c196584d678bb1390cc20a3aeb536f3969a98", // Erc20 deployer private key
//		"45c40d95c9b7898a21e073b5bf952bcb05f2e70072e239a8bbd87bb74a53355e", // Erc721 deployer private key
//		"3737c381633deaaa4c0bdbc64728f6ef7d381b17e1d30bbb74665839cec942b8", // StorageTrie deployer private key
//	}
//	deployerAcc := GetAccountFromKey(0, privateKeyStr[t])
//	log.Printf("prepare%sTransfer, deployer addr:%s", t, deployerAcc.GetAddress().String())
//
//	a.cfg.GetLocalReservoir().TransferSignedTxWithGuaranteeRetry(a.cfg.GetGCli(), deployerAcc, a.cfg.GetChargeValue())
//	a.contracts[t] = deployerAcc.DeploySingleSmartContract(a.cfg.GetGCli(), int(t))
//
//	a.chargeReservoir()
//	a.chargeAccList()
//
//	//transferTcSca := deploySingleSmartContract(erc20DeployAcc, erc20DeployAcc.DeployERC20, "ERC20 Performance Test Contract", cfg)
//	//firstChargeTokenToTestAccounts(erc20TransferTcSca.GetAddress(), erc20DeployAcc.TransferERC20, big.NewInt(1e11), cfg)
//	//chargeTokenToTestAccounts(erc20TransferTcSca.GetAddress(), cfg.GetLocalReservoir().TransferERC20, big.NewInt(1e4), cfg)
//	contracts.SetErc20TransferTcContract(erc20TransferTcSca)
//
//}
