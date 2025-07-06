package tcutil

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

// GaslessTestConfig holds the configuration for gasless transaction tests
type GaslessTestConfig struct {
	EndPoint         string
	NAcc             int
	AccGrp           []*account.Account
	CliPool          clipool.ClientPool
	TestTokenAccount *account.Account
	GsrAccount       *account.Account
	TestType         string // "regular" or "revert"
}

// NewGaslessTestConfig creates a new test configuration
func NewGaslessTestConfig() *GaslessTestConfig {
	return &GaslessTestConfig{}
}

// Init initializes the gasless test configuration
func (cfg *GaslessTestConfig) Init(accs []*account.Account, endpoint string, gp *big.Int, testType string) {
	cfg.EndPoint = endpoint
	cfg.TestType = testType

	cliCreate := func() interface{} {
		c, err := client.Dial(cfg.EndPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	cfg.CliPool.Init(20, 300, cliCreate)

	for _, acc := range accs {
		cfg.AccGrp = append(cfg.AccGrp, acc)
	}

	cfg.NAcc = len(cfg.AccGrp)
}

// Run executes the gasless transaction test
func (cfg *GaslessTestConfig) Run() {
	cli := cfg.CliPool.Alloc().(*client.Client)

	from := cfg.AccGrp[rand.Int()%cfg.NAcc]

	var testRecordName string
	if cfg.TestType == "revert" {
		testRecordName = "TransferNewGaslessRevertTx" + " to " + cfg.EndPoint
	} else {
		testRecordName = "TransferNewGaslessTx" + " to " + cfg.EndPoint
	}

	start := boomer.Now()

	_, _, _, err := from.TransferNewGaslessTx(cli, cfg.EndPoint, cfg.TestTokenAccount, cfg.GsrAccount)

	elapsed := boomer.Now() - start

	cfg.CliPool.Free(cli)

	if err != nil {
		boomer.RecordFailure("http", testRecordName, elapsed, err.Error())
	} else {
		boomer.RecordSuccess("http", testRecordName, elapsed, int64(10))
	}
}
