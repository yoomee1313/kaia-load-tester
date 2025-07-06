package testcase

import (
	"context"
	"errors"
	"fmt"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"sync/atomic"

	"github.com/kaiachain/kaia-load-tester/testcase/tcutil"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

// `analyticTC` is based on the Analytic benchmark in
// [BlockBench](https://github.com/ooibc88/blockbench).

// For your reference, the original description of the Analytic benchmark in the
// SIGMOD paper is like the below:
// > This workload considers the performance of blockchain system in answering
// > analytical queries about the historical data. Similar to an OLAP benchmark,
// > this workload evaluates how the system implements scan-like and aggregate
// > queries, which are determined by its data model.

// `analyticTC` implements three analytic operations using Kaia's JSON RPC API.
// The table below describes the three operations.

// | Function | Description                                                                                                                                           |
// | -------- |-------------------------------------------------------------------------------------------------------------------------------------------------------|
// | `QueryTotalTxVal` | Calculate the sum of transaction's values in the latest 30 blocks. It internally calls `klay_getBlockByNumber` through Kaia's Client interface.       |
// | `QueryLargestTxVal` | Find the largest transaction value in the latest 30 blocks. It internally calls `klay_getBlockByNumber` through Kaia's Client interface.              |
// | `QueryLargestAccBal` | Find the largest balance of a randomly chosen account in the latest 30 blocks. It internally calls `klay_getBalance` through Kaia's Client interface. |
// | `Run` | Randomly invoke one test function among `QueryTotalTxVal`, `QueryLargestTxVal`, and `QueryLargestAccBal`.

func queryTotalTxValRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	conn := config.CliPool.Alloc().(*client.Client)
	msg := "analytic/QueryTotalTxVal to " + config.EndPoint

	// Get the latest block
	start := boomer.Now()
	block, err := conn.BlockByNumber(ctx, nil)
	if err != nil {
		elapsed := boomer.Now() - start
		log.Printf("[Analytic/QueryTotalTxVal] Failed to call BlockByNumber(), err=%v\n", err)
		boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
		conn.Close()
		return
	}

	blockCnt := 1
	txCnt := len(block.Transactions())
	totalValue := sumUpValues(block.Transactions())

	blockNum := block.Number()

	// Do not continue if the blockchain doesn't have 30 blocks
	if blockNum.Int64() < 30 {
		elapsed := boomer.Now() - start
		log.Printf("[Analytic/QueryTotalTxVal] TC needs 30 blocks, but the blockchain has only %v blocks.\n", blockNum)
		boomer.Events.Publish("request_failure", "http", msg, elapsed, "not enough blocks")
		config.CliPool.Free(conn)
		return
	}

	// Read 29 more previous blocks from the latest block
	startNum := new(big.Int).Set(blockNum)
	startNum.Sub(startNum, big.NewInt(29))
	if startNum.Cmp(big.NewInt(0)) == -1 {
		startNum = big.NewInt(0)
	}
	for blockNum.Cmp(startNum) > 0 {
		block, err := conn.BlockByNumber(ctx, blockNum)
		if err != nil {
			elapsed := boomer.Now() - start
			log.Printf("[Analytic/QueryTotalTxVal] Failed to call BlockByNumber(%v), err=%v\n", blockNum, err)
			boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
			conn.Close()
			return
		}

		txCnt += len(block.Transactions())
		totalValue.Add(totalValue, sumUpValues(block.Transactions()))
		blockCnt++

		blockNum.Sub(blockNum, big.NewInt(1))
	}
	elapsed := boomer.Now() - start

	fmt.Printf("[Analytic/QueryTotalTxVal] The total value in %d txs from %d latest blocks: %v (%v ms)\n", txCnt, blockCnt, totalValue, elapsed)
	boomer.Events.Publish("request_success", "http", msg, elapsed, int64(10))
	config.CliPool.Free(conn)
}

func queryLargestTxValRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	conn := config.CliPool.Alloc().(*client.Client)
	msg := "analytic/QueryLargestTxVal to " + config.EndPoint

	// Get the latest block
	start := boomer.Now()
	block, err := conn.BlockByNumber(ctx, nil)
	if err != nil {
		elapsed := boomer.Now() - start
		log.Printf("[Analytic/QueryLargestTxVal] Failed to call BlockByNumber(), err=%v\n", err)
		boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
		conn.Close()
		return
	}

	blockCnt := 1
	txCnt := len(block.Transactions())
	largestValue := findLargestValue(block.Transactions())

	blockNum := block.Number()

	// Do not continue if the blockchain doesn't have 30 blocks
	if blockNum.Int64() < 30 {
		elapsed := boomer.Now() - start
		log.Printf("[Analytic/QueryLargestTxVal] TC needs 30 blocks, but the blockchain has only %v blocks.\n", blockNum)
		boomer.Events.Publish("request_failure", "http", msg, elapsed, "not enough blocks")
		config.CliPool.Free(conn)
		return
	}

	// Read 29 more previous blocks from the latest block
	startNum := new(big.Int).Set(blockNum)
	startNum.Sub(startNum, big.NewInt(29))
	if startNum.Cmp(big.NewInt(0)) == -1 {
		startNum = big.NewInt(0)
	}
	for blockNum.Cmp(startNum) > 0 {
		block, err := conn.BlockByNumber(ctx, blockNum)
		if err != nil {
			elapsed := boomer.Now() - start
			log.Printf("[Analytic/QueryLargestTxVal] Failed to call BlockByNumber(%v), err=%v\n", blockNum, err)
			boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
			conn.Close()
			return
		}

		txCnt += len(block.Transactions())
		val := findLargestValue(block.Transactions())
		if largestValue.Uint64() < val.Uint64() {
			largestValue.Set(val)
		}
		blockCnt++

		blockNum.Sub(blockNum, big.NewInt(1))
	}
	elapsed := boomer.Now() - start

	fmt.Printf("[Analytic/QueryLargestTxVal] The largest value in %d txs from %d latest blocks: %v (%v ms)\n", txCnt, blockCnt, largestValue, elapsed)
	boomer.Events.Publish("request_success", "http", msg, elapsed, int64(10))
	config.CliPool.Free(conn)
}

func queryLargestAccBalRun(config *tcutil.TcConfig) {
	msg := "analytic/QueryLargestAccBal to " + config.EndPoint
	targetAddr := config.AccGrp[rand.Int()%config.NAcc].GetAddress()

	ctx := context.Background()
	conn := config.CliPool.Alloc().(*client.Client)

	// Get the latest block to obtain the block number
	start := boomer.Now()
	block, err := conn.BlockByNumber(ctx, nil)
	if err != nil {
		elapsed := boomer.Now() - start
		log.Printf("[Analytic/QueryLargestAccBal] Failed to call BlockByNumber(), err=%v\n", err)
		boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
		conn.Close()
		return
	}

	blockCnt := 0
	largestBal := big.NewInt(0)

	blockNum := block.Number()

	// Do not continue if the blockchain doesn't have 30 blocks
	if blockNum.Int64() < 30 {
		elapsed := boomer.Now() - start
		log.Printf("[Analytic/QueryLargestAccBal] TC needs 30 blocks, but the blockchain has only %v blocks.\n", blockNum)
		boomer.Events.Publish("request_failure", "http", msg, elapsed, "not enough blocks")
		config.CliPool.Free(conn)
		return
	}

	// Find targetAddr's largest balance in the 30 latest blocks
	startNum := new(big.Int).Set(blockNum)
	startNum.Sub(startNum, big.NewInt(30))
	if startNum.Cmp(big.NewInt(0)) == -1 {
		startNum = big.NewInt(0)
	}
	for blockNum.Cmp(startNum) > 0 {
		bal, err := conn.BalanceAt(ctx, targetAddr, blockNum)
		if err != nil {
			elapsed := boomer.Now() - start
			log.Printf("[Analytic/QueryLargestAccBal] Failed to call BalanceAt(%v), err=%v\n", blockNum, err)
			boomer.Events.Publish("request_failure", "http", msg, elapsed, err.Error())
			conn.Close()
			return
		}

		if largestBal.Uint64() < bal.Uint64() {
			largestBal.Set(bal)
		}

		blockCnt++
		blockNum.Sub(blockNum, big.NewInt(1))
	}
	elapsed := boomer.Now() - start

	fmt.Printf("[Analytic/QueryLargestAccBal] The largest balance of account %s in %d latest blocks: %v (%v ms)\n", targetAddr.String(), blockCnt, largestBal, elapsed)
	boomer.Events.Publish("request_success", "http", msg, elapsed, int64(10))
	config.CliPool.Free(conn)
}

// utils
func sumUpValues(txs types.Transactions) *big.Int {
	totalValue := big.NewInt(0)
	for _, tx := range txs {
		totalValue.Add(totalValue, tx.Value())
	}
	return totalValue
}

func findLargestValue(txs types.Transactions) *big.Int {
	largestValue := big.NewInt(0)
	for _, tx := range txs {
		if largestValue.Uint64() < tx.Value().Uint64() {
			largestValue.Set(tx.Value())
		}
	}
	return largestValue
}

// ## Test Case Description
//
// `doNothingTC` is based on the DoNothing benchmark in
// [BlockBench](https://github.com/ooibc88/blockbench).
//
// For your reference, the original description of DoNothing benchmark in the SIGMOD
// paper is like the below:
// > This contract accepts transaction as input and simply returns. In other
// > words, it involves minimal number of operations at the execution layer and
// > data model layer, thus the overall performance will be mainly determined by
// > the consensus layer. Previous works on performance of blockchain consensus
// > protocol use time to consensus to measure its performance. In BLOCKBENCH,
// > this metric is directly reflected in the transaction latency.
//
// `doNothingTC` implements a Go test function that calls the `nothing` function,
// whose body is empty (i.e., does nothing), in the `Nothing` contract implemented
// in `DoNothing.sol`.
//
// | Function | Description |
// | -------- | ----------- |
// | `Run` | Call the `nothing` function, which does nothing, in the `DoNothing` contract |
func doNothingRun(tcConfig *tcutil.TcConfig) {
	funcName := "Nothing"

	conn := tcConfig.CliPool.Alloc().(*client.Client)

	fromAccount := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	from := fromAccount.GetKey()

	auth := bind.NewKeyedTransactor(from)
	auth.GasLimit = 999999
	auth.GasPrice = tcConfig.GasPrice

	fromAccount.Lock()

	nonce := fromAccount.GetNonce(conn)
	auth.Nonce = big.NewInt(int64(nonce))

	log.Printf("[DoNothing] from=%s nonce=%d %s()\n", fromAccount.GetAddress().String(), nonce, funcName)

	var tx *types.Transaction
	var err error

	start := boomer.Now()
	tx, err = tcConfig.GDoNothing.Nothing(auth)
	elapsed := boomer.Now() - start

	if err != nil {
		log.Printf("[DoNothing] Failed to call %s(), err=%v\n", funcName, err)
		fromAccount.GetNonceFromBlock(conn)
	} else {
		log.Printf("[DoNothing] %s tx=%s\n", funcName, tx.Hash().String())
		fromAccount.UpdateNonce()
	}

	fromAccount.UnLock()

	// Uncomment the below for debugging
	//if err == nil {
	//	utils.CheckReceipt(conn, tx.Hash())
	//}

	msg := "doNothing" + " to " + tcConfig.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		tcConfig.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		conn.Close()
	}
}

//`ioHeavyTC` is based on the IOHeavy benchmark in
//[BlockBench](https://github.com/ooibc88/blockbench).
//
//For your reference, the original description of IOHeavy benchmark in the SIGMOD
//paper is like the below:
//> Current blockchain systems rely on key-value storage to persist blockchain
//> transactions and states. Each storage system may perform differently under
//> different workloads. This workload is designed to evaluate the IO performance
//> by invoking a contract that performs a large number of random writes and
//> random reads to the contract’s states. The I/O bandwidth can be estimated via
//> the observed transaction latency.
//
//Similarly to BlockBench's IOHeavy benchmark, `ioHeavyTC` evaluates the
//performance of writing and reading Kaia's states, especially state variables
//declared in a smart contract.  The `IOHeavy` smart contract in `IOHeavy.sol` is
//used for testing, and its functions are tested through the following interface
//implemented in `ioHeavyTC.go`.
//
//| Function | Description |
//| -------- | ----------- |
//| `Write` | Call the `write` function, which performs multiple write operations against storage variables (actually a map), in the `IOHeavy` contract with a randomly chosen but a range-limited integer key, the number of keys, and a signature (or identifier) recorded in the event log |
//| `Scan` | Call the `scan` function, which reads multiple storage variables (actually a map), in the `IOHeavy` contract with a range-limited integer key, the number of keys, and a signature (or identifier) recorded in the event log |
//| `Run` | Randomly invoke one test function between `Write` and `Scan` |

const maxKey = 100000
const writeSize = 100 // TODO: fixed size vs. random size
const scanSize = 100  // TODO: fixed size vs. random size

const (
	testWrite = iota
	testScan
)

// Run randomly calls Write() or Scan().
func ioHeavyRun(config *tcutil.TcConfig) {
	target := rand.Int() % testLast
	log.Printf("[IOHeavy] calling %s()...\n", toString(target))
	ioHeavyCallFunc(target, config)
}

func ioHeavyScanRun(config *tcutil.TcConfig) {
	ioHeavyCallFunc(testScan, config)
}

func ioHeavyWriteRun(config *tcutil.TcConfig) {
	ioHeavyCallFunc(testWrite, config)
}

func ioHeavyCallFunc(target int, config *tcutil.TcConfig) {
	var size int64

	// Check if target is valid
	switch target {
	case testWrite:
		size = writeSize
	case testScan:
		size = scanSize
	default:
		log.Printf("[IOHeavy] Unknown target: %d\n", target)
		boomer.Events.Publish("request_failure", "contract", "ioHeavy/"+string(target)+" to "+config.EndPoint, 0, "Unknown target")
		return
	}

	// Get the function name as a string
	funcName := toString(target)

	// Choose the start key randomly
	startKey := rand.Int63() % maxKey

	// Signature to distinguish txs
	sig := atomic.AddInt64(&config.GSig, 1)

	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	from := fromAccount.GetKey()

	auth := bind.NewKeyedTransactor(from)
	auth.GasLimit = 99999999
	auth.GasPrice = config.GasPrice

	fromAccount.Lock()

	nonce := fromAccount.GetNonce(conn)
	auth.Nonce = big.NewInt(int64(nonce))

	log.Printf("[IOHeavy] from=%s nonce=%d %s(startKey=%d, size=%d, sig=%d)\n",
		fromAccount.GetAddress().String(), nonce, funcName, startKey, size, sig)

	var tx *types.Transaction
	var err error

	start := boomer.Now()
	switch target {
	case testWrite:
		tx, err = config.GIOHeavy.Write(auth, big.NewInt(startKey), big.NewInt(size), big.NewInt(sig))
	case testScan:
		tx, err = config.GIOHeavy.Scan(auth, big.NewInt(startKey), big.NewInt(size), big.NewInt(sig))
	default:
		log.Printf("[IOHeavy] target %d (%s) is not handled.\n", target, funcName)
		err = errors.New("unhandled target")
	}
	elapsed := boomer.Now() - start

	if err != nil {
		log.Printf("[IOHeavy] Failed to call %s(), err=%v\n", funcName, err)
		fromAccount.GetNonceFromBlock(conn)
	} else {
		log.Printf("[IOHeavy] %s tx=%s\n", funcName, tx.Hash().String())
		fromAccount.UpdateNonce()
	}

	fromAccount.UnLock()

	// Uncomment the below for debugging
	//if err == nil {
	//	utils.CheckReceipt(conn, tx.Hash())
	//}

	msg := "ioHeavy/" + funcName + " to " + config.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		conn.Close()
	}
}

// `smallBankTC` is based on the SmallBank benchmark in
// [BlockBench](https://github.com/ooibc88/blockbench).
//
// For your reference, the original description of SmallBank benchmark in the
// SIGMOD paper is like the below:
// > Unlike YCSB which does not consider transactions, Smallbank is a popular
// > benchmark for OLTP workload. It consists of three tables and four basic
// > procedures simulating basic operations on bank accounts. Smallbank is
// > implemented as a smart contract which simply transfers money from one account
// > to another.
//
// `smallBankTC` tests basic bank operations using the `SmallBank` smart contract.
// The test Go functions in `smallBankTC.go` and their corresponding function in
// the `SmallBank` smart contract are described in the table below.
//
// | Function | Description |
// | -------- | ----------- |
// | `Almagate` | Call the `almagate` function in the `SmallBank` contract, which moves the entire balance in one's checking account to the other's saving account |
// | `GetBalance` | Call the `getBalance` function in the `SmallBank` contract, which returns the balance sum of one's checking and saving accounts |
// | `UpdateBalance` | Call the `updateBalance` function in the `SmallBank` contract, which adds a given value to one's checking account |
// | `UpdateSaving` | Call the `updateSaving` function in the `SmallBank` contract, which adds a given value to one's saving account |
// | `SendPayment` | Call the `sendPayment` function in the `SmallBank` contract, which pays (or moves) a requested value from one's checking account to the other's |
// | `WriteCheck` | Call the `writeCheck` function in the `SmallBank` contract, which tries to mimic the action of issuing a check (however, the implementation logic looks weird) |
// | `Run` | Randomly invoke one test function among `Almagate`, `GetBalance`, `UpdateBalance`, `UpdateSaving`, `SendPayment`, and `WriteCheck` |
//
// ## How to Set the Maximum Number of Users
//
// The environment variable `SMALLBANK_MAX_NUM_USERS` can be used to change the
// maximum number of users, whose default value is 100000.  For example,
// ```shell
// $ SMALLBANK_MAX_NUM_USERS=200000 ./klayslave ...
// ```
//
// Note that the number of users can affect the performance of this test case
// because the smart contract's state could be increased depending on the number
// of different users.
const (
	testAlmagate = iota + 10
	testGetBalance
	testUpdateBalance
	testUpdateSaving
	testSendPayment
	testWriteCheck
	testLast
)

// smallBankRun randomly calls one test case.
func smallBankRun(tcConfig *tcutil.TcConfig) {
	target := rand.Int() % testLast
	log.Printf("[SmallBank] calling %s()...\n", toString(target))
	smallBankCallFunc(target, tcConfig)
}

// smallBankAlmagateRun tests the almagate function in the SmallBank contract.
func smallBankAlmagateRun(tcConfig *tcutil.TcConfig) {
	smallBankCallFunc(testAlmagate, tcConfig)
}

// smallBankGetBalanceRun tests the getBalance function in the SmallBank contract.
func smallBankGetBalanceRun(tcConfig *tcutil.TcConfig) {
	smallBankCallFunc(testGetBalance, tcConfig)
}

// smallBankUpdateBalanceRun tests the updateBalance function in the SmallBank contract.
func smallBankUpdateBalanceRun(tcConfig *tcutil.TcConfig) {
	smallBankCallFunc(testUpdateBalance, tcConfig)
}

// smallBankUpdateSavingRun tests the updateSaving function in the SmallBank contract.
func smallBankUpdateSavingRun(tcConfig *tcutil.TcConfig) {
	smallBankCallFunc(testUpdateSaving, tcConfig)
}

// smallBankSendPaymentRun tests the sendPayment function in the SmallBank contract.
func smallBankSendPaymentRun(tcConfig *tcutil.TcConfig) {
	smallBankCallFunc(testSendPayment, tcConfig)
}

// smallBankWriteCheckRun tests the writeCheck function in the SmallBank contract.
func smallBankWriteCheckRun(tcConfig *tcutil.TcConfig) {
	smallBankCallFunc(testWriteCheck, tcConfig)
}

func smallBankCallFunc(target int, tcConfig *tcutil.TcConfig) {
	// Check if target is valid
	if target < testAlmagate || target >= testLast {
		log.Printf("[SmallBank] Unknown target: %d\n", target)
		boomer.Events.Publish("request_failure", "contract", "smallBank/"+string(target)+" to "+tcConfig.EndPoint, 0, "Unknown target")
		return
	}

	// Get the function name as a string
	funcName := toString(target)

	// Prepare function parameters
	user1 := "user" + strconv.Itoa(rand.Int()%tcConfig.MaxNumUsers)
	user2 := "user" + strconv.Itoa(rand.Int()%tcConfig.MaxNumUsers)
	//log.Printf("[SmallBank] users: %s %s\n", user1, user2)

	conn := tcConfig.CliPool.Alloc().(*client.Client)

	fromAccount := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	from := fromAccount.GetKey()

	var err error
	var elapsed int64

	if target == testGetBalance {
		var balance *big.Int
		callOpts := &bind.CallOpts{Pending: false, From: fromAccount.GetAddress(), Context: nil}

		start := boomer.Now()
		balance, err = tcConfig.GSmallBank.GetBalance(callOpts, user1)
		if err == nil {
			log.Printf("[SmallBank] %s(%s)=%v\n", funcName, user1, balance)
		} else {
			log.Printf("[SmallBank] Failed to call %s(), err=%v\n", funcName, err)
		}
		elapsed = boomer.Now() - start
	} else {
		auth := bind.NewKeyedTransactor(from)
		auth.GasLimit = 9999999
		auth.GasPrice = tcConfig.GasPrice

		fromAccount.Lock()

		nonce := fromAccount.GetNonce(conn)
		auth.Nonce = big.NewInt(int64(nonce))

		log.Printf("[SmallBank] from=%s nonce=%d %s()\n", fromAccount.GetAddress().String(), nonce, funcName)

		var tx *types.Transaction

		start := boomer.Now()
		switch target {
		case testAlmagate:
			tx, err = tcConfig.GSmallBank.Almagate(auth, user1, user2)
		case testUpdateBalance:
			// TODO: use a more meaningful value for the new balance
			tx, err = tcConfig.GSmallBank.UpdateBalance(auth, user1, big.NewInt(0))
		case testUpdateSaving:
			// TODO: use a more meaningful value for the new balance
			tx, err = tcConfig.GSmallBank.UpdateSaving(auth, user1, big.NewInt(0))
		case testSendPayment:
			// TODO: use a more meaningful send value
			tx, err = tcConfig.GSmallBank.SendPayment(auth, user1, user2, big.NewInt(0))
		case testWriteCheck:
			// TODO: use a more meaningful check value
			tx, err = tcConfig.GSmallBank.WriteCheck(auth, user1, big.NewInt(0))
		default:
			log.Printf("[SmallBank] target %d (%s) is not handled.\n", target, funcName)
			err = errors.New("unhandled target")
		}
		elapsed = boomer.Now() - start

		if err != nil {
			log.Printf("[SmallBank] Failed to call %s(), err=%v\n", funcName, err)
			fromAccount.GetNonceFromBlock(conn)
		} else {
			log.Printf("[SmallBank] %s tx=%s\n", funcName, tx.Hash().String())
			fromAccount.UpdateNonce()
		}

		fromAccount.UnLock()

		// Uncomment the below for debugging
		//if err == nil {
		//	utils.CheckReceipt(conn, tx.Hash())
		//}
	}

	msg := "smallBank/" + funcName + " to " + tcConfig.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		tcConfig.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		conn.Close()
	}
}

// `ycsbTC` is based on the YCSB (KVStore) benchmark in
// [BlockBench](https://github.com/ooibc88/blockbench) and is actually a
// simplified version of the original YCSB benchmark.
//
// For your reference, the original description of YCSB in the SIGMOD paper is
// like the below:
// > YCSB is a simple smart contract which functions as a key-value storage.  It
// > preloads each store with a number of records, and supports requests with
// > different ratios of read and write operations. YCSB is widely used for
// > evaluating NoSQL databases.
//
// `ycsbTC` tests a key-value storage, where both key and value are string, by
// calling `set` and `get` functions defined in the `KVstore` smart contract
// (`kvstore.sol`).  It provides the following functions for testing:
//
// | Function | Description |
// | -------- | ----------- |
// | `Set` | Call the `set` function in the `KVstore` contract with a randomly chosen but a range-limited key and a random value |
// | `Get` | Call the `get` function in the `KVstore` contract with a randomly chosen key, which has the same range as does the key used for the `set` function |
// | `Run` | Randomly invoke one test function between `Set` and `Get` |
//
// ## How to Set the Maximum Number of Keys
//
// The environment variable `YCSB_MAX_NUM_KEYS` can be used to change the maximum
// number of keys, whose default value is 100000.  For example,
// ```shell
// $ YCSB_MAX_NUM_KEYS=200000 ./klayslave ...
// ```
//
// Note that the number of keys can affect the performance of this test case
// because the smart contract's state could be increased depending on the number
// of different keys.
const (
	testSet = iota + 20
	testGet
)

// Run randomly calls one test case.
func ycsbRun(tcConfig *tcutil.TcConfig) {
	target := rand.Int() % testLast
	log.Printf("[YCSB] calling %s()...\n", toString(target))
	ycsbCallFunc(target, tcConfig)
}

// Set tests the set function in the KVstore contract.
func ycsbSetRun(tcConfig *tcutil.TcConfig) {
	ycsbCallFunc(testSet, tcConfig)
}

// Get tests the get function in the KVstore contract.
func ycsbGetRun(tcConfig *tcutil.TcConfig) {
	ycsbCallFunc(testGet, tcConfig)
}

func ycsbCallFunc(target int, tcConfig *tcutil.TcConfig) {
	// Check if target is valid
	if target < testSet || target >= testLast {
		log.Printf("[YCSB] Unknown target: %d\n", target)
		boomer.Events.Publish("request_failure", "contract", "ycsb/"+string(target)+" to "+tcConfig.EndPoint, 0, "Unknown target")
		return
	}

	// Get the function name as a string
	funcName := toString(target)

	// Prepare function parameters
	user := "user" + strconv.Itoa(rand.Int()%tcConfig.MaxNumKeys)
	val := "val" + strconv.Itoa(rand.Int())

	conn := tcConfig.CliPool.Alloc().(*client.Client)

	fromAccount := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	from := fromAccount.GetKey()

	var err error
	var elapsed int64

	switch target {
	case testSet:
		auth := bind.NewKeyedTransactor(from)
		auth.GasLimit = 9999999
		auth.GasPrice = tcConfig.GasPrice

		fromAccount.Lock()

		nonce := fromAccount.GetNonce(conn)
		auth.Nonce = big.NewInt(int64(nonce))

		log.Printf("[YCSB] from=%s nonce=%d %s()\n", fromAccount.GetAddress().String(), nonce, funcName)
		log.Printf("[YCSB] %s(%s, %s)\n", funcName, user, val)

		var tx *types.Transaction

		start := boomer.Now()
		tx, err = tcConfig.GKVstore.Set(auth, user, val)
		elapsed = boomer.Now() - start

		if err != nil {
			log.Printf("[YCSB] Failed to call %s(), err=%v\n", funcName, err)
			fromAccount.GetNonceFromBlock(conn)
		} else {
			log.Printf("[YCSB] %s tx=%s\n", funcName, tx.Hash().String())
			fromAccount.UpdateNonce()
		}

		fromAccount.UnLock()

		// Uncomment the below for debugging
		//if err == nil {
		//	utils.CheckReceipt(conn, tx.Hash())
		//}

	case testGet:
		var value string
		callOpts := &bind.CallOpts{Pending: false, From: fromAccount.GetAddress(), Context: nil}

		start := boomer.Now()
		value, err = tcConfig.GKVstore.Get(callOpts, user)
		if err == nil {
			log.Printf("[YCSB] %s(%s)=%v\n", funcName, user, value)
		} else {
			log.Printf("[YCSB] Failed to call %s(), err=%v\n", funcName, err)
		}
		elapsed = boomer.Now() - start

	default:
		log.Printf("[YCSB] target %d (%s) is not handled.\n", target, funcName)
		err = errors.New("unhandled target")
	}

	msg := "ycsb/" + funcName + " to " + tcConfig.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		tcConfig.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		conn.Close()
	}
}

func toString(target int) string {
	switch target {
	// ioHeavy
	case testWrite:
		return "Write"
	case testScan:
		return "Scan"
	// smallBank
	case testAlmagate:
		return "Almagate"
	case testGetBalance:
		return "GetBalance"
	case testUpdateBalance:
		return "UpdateBalance"
	case testUpdateSaving:
		return "UpdateSaving"
	case testSendPayment:
		return "SendPayment"
	case testWriteCheck:
		return "WriteCheck"
	// ycsb
	case testSet:
		return "Set"
	case testGet:
		return "Get"
	default:
		return "Unknown"
	}
}
