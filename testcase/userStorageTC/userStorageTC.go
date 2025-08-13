package userStorageTC

import (
	"log"
	"math/big"
	"math/rand"
	"sync"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "userStorageTC"

var (
	endPoint string

	nAcc   int
	accGrp []*account.Account

	SmartContractAccount *account.Account

	cliPool clipool.ClientPool

	mutex       sync.Mutex
	initialized = false

	gasPrice *big.Int
)

func Init(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	mutex.Lock()
	defer mutex.Unlock()

	if initialized {
		return
	}
	initialized = true

	gasPrice = gp

	endPoint = endpoint
	SmartContractAccount = contractsParam[account.ContractUserStorage]

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

func RunSet() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	from := accGrp[rand.Int()%nAcc]
	value := big.NewInt(1) // set function
	data := account.TestContractInfos[account.ContractUserStorage].GenData(from.GetAddress(), value)

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, SmartContractAccount, nil, data)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", Name+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, elapsed, err.Error())
	}
}

func RunSetGet() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	from := accGrp[rand.Int()%nAcc]

	start := boomer.Now()

	// First, call set function
	setValue := big.NewInt(1) // set function
	setData := account.TestContractInfos[account.ContractUserStorage].GenData(from.GetAddress(), setValue)
	_, _, setErr := from.TransferNewSmartContractExecutionTx(cli, SmartContractAccount, nil, setData)

	if setErr != nil {
		elapsed := boomer.Now() - start
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, elapsed, setErr.Error())
		return
	}

	// Then, call get function
	getValue := big.NewInt(0) // get function
	getData := account.TestContractInfos[account.ContractUserStorage].GenData(from.GetAddress(), getValue)
	_, _, getErr := from.TransferNewSmartContractExecutionTx(cli, SmartContractAccount, nil, getData)

	elapsed := boomer.Now() - start

	if getErr == nil {
		boomer.Events.Publish("request_success", "http", Name+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, elapsed, getErr.Error())
	}
}
