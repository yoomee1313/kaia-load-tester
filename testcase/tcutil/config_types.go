package tcutil

import (
	"math/big"
	"sync"

	"github.com/kaiachain/kaia-load-tester/testcase/contracts/cpuHeavyTC"
	"github.com/kaiachain/kaia/common"
)

// CpuHeavyConfig contains configuration specific to CPU-heavy test cases
type CpuHeavyConfig struct {
	*BaseConfig
	CpuHeavyContract *cpuHeavyTC.CpuHeavyTC
}

// NewCpuHeavyConfig creates a new CpuHeavyConfig
func NewCpuHeavyConfig(base *BaseConfig) *CpuHeavyConfig {
	return &CpuHeavyConfig{
		BaseConfig: base,
	}
}

// SetCpuHeavyContract sets the CPU heavy contract
func (c *CpuHeavyConfig) SetCpuHeavyContract(contract *cpuHeavyTC.CpuHeavyTC) {
	c.CpuHeavyContract = contract
}

// GetCpuHeavyContractAddress returns the address of the CPU heavy contract
func (c *CpuHeavyConfig) GetCpuHeavyContractAddress() common.Address {
	if c.CpuHeavyContract == nil {
		return common.Address{}
	}
	// The contract's address is stored in the CpuHeavyTCCaller's contract field
	// which is a *bind.BoundContract containing the address
	return c.CpuHeavyContract.CpuHeavyTCCaller.Contract.Address()
}

// AnalyticConfig contains configuration for analytic test cases
type AnalyticConfig struct {
	*BaseConfig
	// Add analytic-specific fields here
}

// NewAnalyticConfig creates a new AnalyticConfig
func NewAnalyticConfig(base *BaseConfig) *AnalyticConfig {
	return &AnalyticConfig{
		BaseConfig: base,
	}
}

// ReadAPIConfig contains configuration for read API test cases
type ReadAPIConfig struct {
	*BaseConfig
	LatestBlockNumber *big.Int
	Count            uint64
}

// NewReadAPIConfig creates a new ReadAPIConfig
func NewReadAPIConfig(base *BaseConfig) *ReadAPIConfig {
	return &ReadAPIConfig{
		BaseConfig: base,
	}
}

// BlockbenchConfig contains configuration for blockbench test cases
type BlockbenchConfig struct {
	*BaseConfig
	GSig        int64 // should be updated atomically
	MaxNumUsers int   // maxNumUsers determines the maximum number of users used in tests
	MaxNumKeys  int   // maxNumKeys determines the maximum number of keys used in tests
}

// NewBlockbenchConfig creates a new BlockbenchConfig
func NewBlockbenchConfig(base *BaseConfig) *BlockbenchConfig {
	return &BlockbenchConfig{
		BaseConfig: base,
	}
}

// ReceiptCheckConfig contains configuration for receipt check test cases
type ReceiptCheckConfig struct {
	*BaseConfig
	HashPool []common.Hash
	rwMutex  sync.RWMutex
	tail     int
	isFull   bool
}

// NewReceiptCheckConfig creates a new ReceiptCheckConfig
func NewReceiptCheckConfig(base *BaseConfig, hashPoolSize int) *ReceiptCheckConfig {
	return &ReceiptCheckConfig{
		BaseConfig: base,
		HashPool:   make([]common.Hash, hashPoolSize),
	}
}
