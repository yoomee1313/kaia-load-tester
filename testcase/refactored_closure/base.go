package refactored_closure

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

// TCConfig holds common configuration for test cases
type TCConfig struct {
	Name                 string
	EndPoint             string
	NAcc                 int
	AccGrp               []*account.Account
	CliPool              clipool.ClientPool
	SmartContractAccount *account.Account
}

// Init initializes common configuration for test cases
func Init(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) *TCConfig {
	config := &TCConfig{
		EndPoint: endpoint,
	}

	cliCreate := func() interface{} {
		c, err := client.Dial(config.EndPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	config.CliPool.Init(20, 300, cliCreate)

	for _, acc := range accs {
		config.AccGrp = append(config.AccGrp, acc)
	}

	config.NAcc = len(config.AccGrp)
	return config
}

// RunBaseValueTransfer executes a test case with common logic
func RunBaseValueTransfer(config *TCConfig, txFunc func(*client.Client, *account.Account, *account.Account) (interface{}, *big.Int, error), txName string) {
	cli := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(cli)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]

	start := boomer.Now()
	_, _, err := txFunc(cli, from, to)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", txName+" to "+config.EndPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", txName+" to "+config.EndPoint, elapsed, err.Error())
	}
}

// RunBaseWithContract executes a test case with contract account
func RunBaseWithContract(config *TCConfig, txFunc func(*client.Client, *account.Account, *account.Account) (interface{}, *big.Int, error), txName string) {
	cli := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(cli)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.SmartContractAccount

	start := boomer.Now()
	_, _, err := txFunc(cli, from, to)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", txName+" to "+config.EndPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", txName+" to "+config.EndPoint, elapsed, err.Error())
	}
}
