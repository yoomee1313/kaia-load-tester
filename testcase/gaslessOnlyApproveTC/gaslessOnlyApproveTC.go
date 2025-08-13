package gaslessOnlyApproveTC

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "gaslessOnlyApproveTC"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	TestTokenAccount *account.Account
	GsrAccount       *account.Account
)

func Init(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	endPoint = endpoint
	TestTokenAccount = contractsParam[account.ContractGaslessToken]
	GsrAccount = contractsParam[account.ContractGaslessSwapRouter]

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
	testRecordName := "TransferNewGaslessApproveTx" + " to " + endPoint

	start := boomer.Now()

	_, _, err := from.TransferNewGaslessApproveTx(cli, endPoint, TestTokenAccount, GsrAccount)

	elapsed := boomer.Now() - start

	if err != nil {
		boomer.RecordFailure("http", testRecordName, elapsed, err.Error())
	} else {
		boomer.RecordSuccess("http", testRecordName, elapsed, int64(10))
	}
}
