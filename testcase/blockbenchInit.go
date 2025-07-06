package testcase

import (
	"context"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/IOHeavyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/doNothingTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/smallBankTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/blockbench/ycsbTC"
	"github.com/kaiachain/kaia-load-tester/testcase/tcutil"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
)

func doNothingInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// deploy do nothing contract
	coinbase := accs[0]
	conn, ok := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)
	if !ok {
		log.Fatal("[DoNothing] conn is not client.Client")
		return nil
	}

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))

	var address common.Address
	var tx *types.Transaction
	log.Println("[DoNothing] Deploying a new smart contract")

	for {
		var err error
		address, tx, config.GDoNothing, err = doNothingTC.DeployDoNothingTC(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[DoNothing] Failed to deploy the contract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}
	log.Printf("[DoNothing] Contract address : 0x%x\n", address)
	log.Printf("[DoNothing] Transaction waiting to be mined: 0x%x\n", tx.Hash())

	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Printf("[DoNothing] Failed to check receipt: %v\n", err)
			continue
		}
		log.Println("[DoNothing] Received the receipt")
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Println("[DoNothing] Contract deployment was successful")
			break
		} else {
			log.Fatalf("[DoNothing] Contract Receipt Status: %v\n", receipt.Status)
		}
	}
	return config
}

func ioHeavyInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	config.GSig = 0

	// deploy io heavy contract
	coinbase := accs[0]
	conn, ok := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)
	if !ok {
		log.Fatal("[IOHeavy] conn is not client.Client")
		return nil
	}

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 9999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))

	var address common.Address
	var tx *types.Transaction
	log.Println("[IOHeavy] Deploying a new smart contract")

	for {
		var err error
		address, tx, config.GIOHeavy, err = IOHeavyTC.DeployIoHeavyTC(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[IOHeavy] Failed to deploy the contract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}
	log.Printf("[IOHeavy] Contract address : 0x%x\n", address)
	log.Printf("[IOHeavy] Transaction waiting to be mined: 0x%x\n", tx.Hash())

	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Printf("[IOHeavy] Failed to check receipt: %v\n", err)
			continue
		}
		log.Println("[IOHeavy] Received the receipt")
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Println("[IOHeavy] Contract deployment was successful")
			break
		} else {
			log.Fatalf("[IOHeavy] Contract Receipt Status: %v\n", receipt.Status)
		}
	}
	return config
}

func smallBankInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// Change the maximum number of users if the environment variable SMALLBANK_MAX_NUM_USERS has been set
	config.MaxNumUsers = 100000
	if v := os.Getenv("SMALLBANK_MAX_NUM_USERS"); v != "" {
		if envVal, err := strconv.Atoi(v); err == nil {
			config.MaxNumUsers = envVal
		}
	}
	log.Printf("[SmallBank] maxNumUsers=%d\n", config.MaxNumUsers)

	// deploy small back contract
	coinbase := accs[0]

	conn, ok := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)
	if !ok {
		log.Fatal("[SmallBank] conn is not client.Client")
		return nil
	}

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 9999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))

	var address common.Address
	var tx *types.Transaction
	log.Println("[SmallBank] Deploying a new smart contract")

	for {
		var err error
		address, tx, config.GSmallBank, err = smallBankTC.DeploySmallBankTC(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[SmallBank] Failed to deploy the contract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}
	log.Printf("[SmallBank] Contract address : 0x%x\n", address)
	log.Printf("[SmallBank] Transaction waiting to be mined: 0x%x\n", tx.Hash())

	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Printf("[SmallBank] Failed to check receipt: %v\n", err)
			continue
		}
		log.Println("[SmallBank] Received the receipt")
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Println("[SmallBank] Contract deployment was successful")
			break
		} else {
			log.Fatalf("[SmallBank] Contract Receipt Status: %v\n", receipt.Status)
		}
	}
	return config
}

func ycsbInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// Change the maximum number of keys if the environment variable YCSB_MAX_NUM_KEYS has been set
	config.MaxNumKeys = 100000
	if v := os.Getenv("YCSB_MAX_NUM_KEYS"); v != "" {
		if envVal, err := strconv.Atoi(v); err == nil {
			config.MaxNumKeys = envVal
		}
	}
	log.Printf("[YCSB] maxNumKeys=%d\n", config.MaxNumKeys)

	// deploy ycsb contract
	coinbase := accs[0]
	conn, ok := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)
	if !ok {
		log.Fatal("[YCSB] conn is not client.Client")
		return nil
	}

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))

	var address common.Address
	var tx *types.Transaction
	log.Println("[YCSB] Deploying a new smart contract")

	for {
		var err error
		address, tx, config.GKVstore, err = ycsbTC.DeployYcsbTC(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[YCSB] Failed to deploy the contract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}
	log.Printf("[YCSB] Contract address : 0x%x\n", address)
	log.Printf("[YCSB] Transaction waiting to be mined: 0x%x\n", tx.Hash())

	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Printf("[YCSB] Failed to check receipt: %v\n", err)
			continue
		}
		log.Println("[YCSB] Received the receipt")
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Println("[YCSB] Contract deployment was successful")
			break
		} else {
			log.Fatalf("[YCSB] Contract Receipt Status: %v\n", receipt.Status)
		}
	}

	return config
}
