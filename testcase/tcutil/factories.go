package tcutil

import (
	"log"
	"math/big"
	"time"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/cpuHeavyTC"
	"github.com/kaiachain/kaia/common"
	"github.com/kaiachain/kaia/core/types"
	"github.com/kaiachain/kaia/client"
)

// NewCpuHeavyTestConfig creates and initializes a new CpuHeavyConfig
func NewCpuHeavyTestConfig(accs []*account.Account, endpoint string, gp *big.Int) *CpuHeavyConfig {
	base := InitBaseConfig(accs, endpoint, gp)
	config := NewCpuHeavyConfig(base)

	// Deploy CPU-heavy contract
	conn := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)

	coinbase := accs[0]
	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonce(conn)))

	var address common.Address
	var cpuHeavy *cpuHeavyTC.CpuHeavyTC
	var err error

	// Keep retrying until successful
	for {
		var tx *types.Transaction
		address, tx, cpuHeavy, err = cpuHeavyTC.DeployCpuHeavyTC(auth, conn)
		_ = tx // Explicitly mark as used
		if err == nil {
			break
		}

		log.Printf("Failed to deploy contract: %v, retrying...\n", err)
		time.Sleep(1 * time.Second)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonce(conn)))
	}

	config.SetCpuHeavyContract(cpuHeavy)
	log.Printf("CPU Heavy contract deployed at: %s\n", address.Hex())

	// Wait for contract deployment to be mined
	time.Sleep(2 * time.Second)

	return config
}

// NewAnalyticTestConfig creates and initializes a new AnalyticConfig
func NewAnalyticTestConfig(accs []*account.Account, endpoint string, gp *big.Int) *AnalyticConfig {
	base := InitBaseConfig(accs, endpoint, gp)
	return NewAnalyticConfig(base)
}

// NewReadAPITestConfig creates and initializes a new ReadAPIConfig
func NewReadAPITestConfig(accs []*account.Account, endpoint string, gp *big.Int) *ReadAPIConfig {
	base := InitBaseConfig(accs, endpoint, gp)
	config := NewReadAPIConfig(base)
	// Initialize any read API specific configurations here
	return config
}

// NewBlockbenchTestConfig creates and initializes a new BlockbenchConfig
func NewBlockbenchTestConfig(accs []*account.Account, endpoint string, gp *big.Int) *BlockbenchConfig {
	base := InitBaseConfig(accs, endpoint, gp)
	config := NewBlockbenchConfig(base)
	// Default values for blockbench tests
	config.MaxNumUsers = 1000
	config.MaxNumKeys = 10000
	return config
}

// NewReceiptCheckTestConfig creates and initializes a new ReceiptCheckConfig
func NewReceiptCheckTestConfig(accs []*account.Account, endpoint string, gp *big.Int) *ReceiptCheckConfig {
	base := InitBaseConfig(accs, endpoint, gp)
	hashPoolSize := 100 * 5 * 60 // for 5 minutes at 100 TPS
	return NewReceiptCheckConfig(base, hashPoolSize)
}
