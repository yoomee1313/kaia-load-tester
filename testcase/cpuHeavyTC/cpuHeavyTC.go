package cpuHeavyTC

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "cpuHeavyTx"

var (
	endPoint string

	nAcc   int
	accGrp []*account.Account

	cliPool clipool.ClientPool

	gasPrice *big.Int

	// multinode tester
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
	defer cliPool.Free(cli)

	from := accGrp[rand.Int()%nAcc]
	value := big.NewInt(100) // size parameter
	data := account.TestContractInfos[account.ContractCPUHeavy].GenData(from.GetAddress(), value)

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, SmartContractAccount, nil, data)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", Name+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, elapsed, err.Error())
	}
}
