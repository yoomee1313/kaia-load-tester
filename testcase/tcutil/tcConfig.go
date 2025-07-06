package tcutil

import (
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/IOHeavyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/doNothingTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/smallBankTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/ycsbTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/cpuHeavyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/largeMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/readApiCallContractTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/userStorageTC"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
	"log"
	"math/big"
	"math/rand"
	"sync"
)

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
func NewTcConfig() *TcConfig {
	return &TcConfig{}
}

// InitTcConfig initializes a TcConfig with common settings
func InitTcConfig(accs []*account.Account, endpoint string, gp *big.Int) *TcConfig {
	config := NewTcConfig()
	config.GasPrice = gp
	config.EndPoint = endpoint

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
