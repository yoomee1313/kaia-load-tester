package tcutil

import (
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"math/big"
)

// Deprecated: TcConfig is deprecated. Use specific config types from the config package instead.
// TcConfig holds the common configuration for test cases
type TcConfig struct {
	EndPoint string
	NAcc     int
	AccGrp   []*account.Account
	CliPool  clipool.ClientPool
	GasPrice *big.Int

	// multinode tester
	TransferedValue *big.Int
	ExpectedFee     *big.Int

	FromAccount     *account.Account
	PrevBalanceFrom *big.Int

	ToAccount     *account.Account
	PrevBalanceTo *big.Int

	// smart contract
	SmartContractAccount *account.Account
	GCPUHeavy            *cpuHeavyTC.CpuHeavyTC
	GLargeMemo           *largeMemoTC.LargeMemoTC
	GUserStorage         *userStorageTC.UserStorageTC
	ReadApiCallContract  *readApiCallContractTC.ReadApiCallContractTC
	KIP17Contract        *bind.BoundContract
	MainContract         *bind.BoundContract
	GDoNothing           *doNothingTC.DoNothingTC
	GIOHeavy             *IOHeavyTC.IoHeavyTC
	GSmallBank           *smallBankTC.SmallBankTC
	GKVstore             *ycsbTC.YcsbTC

	// read api call contract
	RetValOfCall        *big.Int
	RetValOfStorageAt   *big.Int
	RetValOfEstimateGas uint64

	// read api call
	LatestBlockNumber *big.Int
	Count             uint64

	// ethereumTx
	ExecutablePath string

	// blockbench
	GSig        int64 // should be updated atomically
	MaxNumUsers int   // maxNumUsers determines the maximum number of users used in tests.
	MaxNumKeys  int   // maxNumKeys determines the maximum number of keys used in tests.

	// receiptCheck
	HashPool []common.Hash
	rwMutex  sync.RWMutex
	tail     int
	isFull   bool
}

// NewTcConfig creates a new TcConfig instance
// Deprecated: Use specific config initializers from the tcutil package instead
func NewTcConfig() *TcConfig {
	return &TcConfig{}
}

// InitTcConfig initializes a TcConfig with common settings
// Deprecated: Use InitBaseConfig or specific config initializers instead
func InitTcConfig(accs []*account.Account, endpoint string, gp *big.Int) *TcConfig {
	base := InitBaseConfig(accs, endpoint, gp)
	
	// Convert to the old TcConfig for backward compatibility
	config := &TcConfig{
		EndPoint: base.EndPoint,
		NAcc:     base.NAcc,
		AccGrp:   base.AccGrp,
		CliPool:  base.CliPool,
		GasPrice: base.GasPrice,
	}

	return config
}

var hashPoolSize = 100 * 5 * 60 // for init 5min, if input send tps is 100Txs/Sec

func (tCfg *TcConfig) AddHash(hash common.Hash) {
	tCfg.rwMutex.Lock()
	defer tCfg.rwMutex.Unlock()

	tCfg.HashPool[tCfg.tail] = hash

	tCfg.tail = (tCfg.tail + 1) % hashPoolSize
	if tCfg.tail == 0 {
		tCfg.isFull = true
	}
}

func (tCfg *TcConfig) GetHash() common.Hash {
	tCfg.rwMutex.RLock()
	defer tCfg.rwMutex.RUnlock()
	if tCfg.isFull {
		return tCfg.HashPool[rand.Int()%hashPoolSize]
	}
	return tCfg.HashPool[rand.Int()%tCfg.tail]
}
