package internalTxTC

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const (
	Name        = "internalTxTC"
	NameMintNFT = "mintNFTTC"
)

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	// Contract accounts
	KIP17ContractAccount *account.Account
	MainContractAccount  *account.Account
)

func Init(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	endPoint = endpoint
	KIP17ContractAccount = contractsParam[account.ContractInternalTxKIP17]
	MainContractAccount = contractsParam[account.ContractInternalTxMain]

	cliCreate := func() interface{} {
		c, err := client.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	cliPool.Init(20, 300, cliCreate)

	for _, acc := range accs {
		accGrp = append(accGrp, acc)
	}

	nAcc = len(accGrp)
}

// Run transfers txs calling sendRewards function of mainContract.
// During the execution of the function, four internal transactions are triggered also:
// Mint KIP17 token, Transfer KIP17 token, send KLAY to inviteeAccount and hostAccount.
func Run() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	from := accGrp[rand.Int()%nAcc]

	// Generate data for sendRewards function call
	data := account.TestContractInfos[account.ContractInternalTxMain].GenData(from.GetAddress(), big.NewInt(100))

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, MainContractAccount, big.NewInt(100), data)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", Name+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, elapsed, err.Error())
	}
}

// RunMintNFT transfers txs calling mintCard function of KIP17Contract.
// The function mints a KIP17 token, NFT, for the sender.
func RunMintNFT() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	from := accGrp[rand.Int()%nAcc]

	// Generate data for mintCard function call
	data := account.TestContractInfos[account.ContractInternalTxKIP17].GenData(from.GetAddress(), nil)

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, KIP17ContractAccount, nil, data)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", NameMintNFT+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", NameMintNFT+" to "+endPoint, elapsed, err.Error())
	}
}
