package tcutil

import (
	"log"
	"math/big"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
)

// TestCaseConfig defines the interface that all test case configurations must implement
type TestCaseConfig interface {
	// Common methods
	GetEndPoint() string
	GetNAcc() int
	GetAccGrp() []*account.Account
	GetCliPool() clipool.ClientPool
	GetGasPrice() *big.Int

	// Optional methods for specific test cases
	// These will be implemented by specific config types when needed
}

// BaseConfig contains common configuration fields for all test cases
type BaseConfig struct {
	EndPoint string
	NAcc     int
	AccGrp   []*account.Account
	CliPool  clipool.ClientPool
	GasPrice *big.Int
}

// Common getter methods for BaseConfig
func (c *BaseConfig) GetEndPoint() string {
	return c.EndPoint
}

func (c *BaseConfig) GetNAcc() int {
	return c.NAcc
}

func (c *BaseConfig) GetAccGrp() []*account.Account {
	return c.AccGrp
}

func (c *BaseConfig) GetCliPool() clipool.ClientPool {
	return c.CliPool
}

func (c *BaseConfig) GetGasPrice() *big.Int {
	return c.GasPrice
}

// InitBaseConfig initializes a base configuration with common settings
func InitBaseConfig(accs []*account.Account, endpoint string, gp *big.Int) *BaseConfig {
	config := &BaseConfig{
		GasPrice: gp,
		EndPoint: endpoint,
		AccGrp:   make([]*account.Account, 0, len(accs)),
	}

	// Initialize client pool
	cliCreate := func() interface{} {
		c, err := client.Dial(config.EndPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	config.CliPool.Init(20, 300, cliCreate)

	// Copy accounts
	config.AccGrp = append(config.AccGrp, accs...)
	config.NAcc = len(config.AccGrp)

	return config
}
