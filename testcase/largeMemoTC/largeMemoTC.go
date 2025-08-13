//go:generate abigen --sol LargeMemo.sol --pkg largeMemoTC --out LargeMemo.go

// Package largeMemoTC is used to test required network bandwidth for large block sizes.
// tries to simulate bots which exhausts resource
// See README.md for more details.
package largeMemoTC

import (
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "largeMemoTC"
const Letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	SmartContractAccount *account.Account
)

// Init initializes cliPool and accGrp; and also deploys the smart contract.
func Init(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	endPoint = endpoint
	SmartContractAccount = contractsParam[account.ContractLargeMemo]

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

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Letters[r.Intn(len(Letters))]
	}
	return string(b)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Run() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	from := accGrp[rand.Int()%nAcc]
	value := big.NewInt(int64(randInt(50, 2000))) // memo size
	data := account.TestContractInfos[account.ContractLargeMemo].GenData(from.GetAddress(), value)

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, SmartContractAccount, nil, data)
	elapsed := boomer.Now() - start

	msg := "LargeMemo" + " to " + endPoint
	if err == nil {
		boomer.Events.Publish("request_success", "http", msg, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
	}
}
