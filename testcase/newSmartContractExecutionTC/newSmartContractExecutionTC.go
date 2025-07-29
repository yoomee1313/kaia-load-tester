package newSmartContractExecutionTC

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
	"github.com/kaiachain/kaia/params"
	"github.com/myzhan/boomer"
)

const Name = "newSmartContractExecutionTC"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool
	gasPrice *big.Int

	// multinode tester
	expectedFee *big.Int

	SmartContractAccount *account.Account
)

func Init(accs []*account.Account, endpoint string, gp *big.Int) {
	gasPrice = gp

	endPoint = endpoint

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

func Run() {
	cli := cliPool.Alloc().(*client.Client)

	from := accGrp[rand.Int()%nAcc]
	to := SmartContractAccount

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, to, nil, account.TestContractInfos[account.ContractGeneral].GenData(from.GetAddress(), nil))
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewSmartContractExecutionTx"+" to "+endPoint, elapsed, int64(10))
		cliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewSmartContractExecutionTx"+" to "+endPoint, elapsed, err.Error())
	}
}

func RunSingle() (txHash common.Hash, err error) {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	fromIdx := rand.Int() % nAcc

	from := accGrp[fromIdx]
	to := SmartContractAccount
	expectedFee = big.NewInt(0).Mul(big.NewInt(25*params.Gkei), big.NewInt(21000))

	tx, _, err := from.TransferNewSmartContractExecutionTx(cli, to, nil, account.TestContractInfos[account.ContractGeneral].GenData(from.GetAddress(), nil))
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), err
}
