package testcase

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/kaiachain/kaia"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/cpuHeavyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/internalTxTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/largeMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/readApiCallContractTC"
	"github.com/kaiachain/kaia-load-tester/testcase/contracts/userStorageTC"
	"github.com/kaiachain/kaia-load-tester/testcase/tcutil"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
)

func cpuHeavyInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// deploy heavy cpu contract
	conn := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)

	coinbase := accs[0]
	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonce(conn)))

	var address common.Address
	var tx *types.Transaction
	println("Deploying new smart contract")

	for {
		addr, tTx, cpuHeavy, err := cpuHeavyTC.DeployCpuHeavyTC(auth, conn)
		address = addr
		tx = tTx
		if err != nil {
			coinbase.UpdateNonce()
		}
		config.GCPUHeavy = cpuHeavy

		if err != nil {
			//log.Printf("Failed to deploy new contract: %v\n", err)
		} else {
			break
		}
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second) // Avoiding Nonce corruption
	}
	fmt.Printf("=> Contract pending deploy: 0x%x\n", address)

	fmt.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond) // Allow it to be processed by the local node :P
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			//fmt.Printf("Failed to check receipt: %v\n", err)
			continue
		}
		fmt.Printf("=> Contract Receipt Status: %v\n", receipt.Status)
		break
	}
	return config
}

func largeMemoInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// deploy large memo contract
	conn, ok := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)
	if !ok {
		log.Fatal("[LargeMemo] conn is not client.Client")
		return nil
	}

	coinbase := accs[0]
	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))

	var tx *types.Transaction

	for {
		var err error
		_, tx, config.GLargeMemo, err = largeMemoTC.DeployLargeMemoTC(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[LargeMemo] Failed to deploy the contract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}

	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Printf("[LargeMemo] Failed to check receipt: %v\n", err)
			continue
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			break
		} else {
			log.Fatalf("[LargeMemo] Contract Receipt Status: %v\n", receipt.Status)
		}
	}

	return config
}

func internalTxInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// deploy internal tx contract
	coinbase := accs[0]

	conn, ok := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)
	if !ok {
		log.Fatal("[internalTxTC] conn is not client.Client")
		return nil
	}

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 9999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
	log.Println("[internalTxTC] Deploying two smart contracts")

	ctx := context.Background()
	defer ctx.Done()

	// Deploy Token Contract
	var (
		KIP17Tx      *types.Transaction
		KIP17Address common.Address
	)

	for {
		var err error
		KIP17Address, KIP17Tx, config.KIP17Contract, err = internalTxTC.DeployKIP17TokenContract(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[internalTxTC] Failed to deploy the KIP17 token mainContract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}
	log.Printf("[internalTxTC] KIP17 token contract address: 0x%x\n", KIP17Address)
	log.Printf("[internalTxTC] Transaction waiting to be mined: 0x%x\n", KIP17Tx.Hash())

	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, KIP17Tx.Hash())
		if err != nil {
			log.Printf("[internalTxTC] Failed to check receipt: %v\n", err)
			continue
		}
		log.Println("[internalTxTC] Received the receipt")
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Println("[internalTxTC] Contract deployment was successful")
			break
		} else {
			log.Fatalf("[internalTxTC] Contract Receipt Status: %v\n", receipt.Status)
		}
	}

	// Deploy Main Contract
	var (
		mainAddress common.Address
		mainTx      *types.Transaction
	)
	auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
	for {
		var err error
		mainAddress, mainTx, config.MainContract, err = internalTxTC.DeployMainContract(auth, conn, KIP17Address)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}

		log.Printf("[internalTxTC] Failed to deploy the main contract, err=%v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second)
	}
	log.Printf("[internalTxTC] Main contract address : 0x%x\n", mainAddress)
	log.Printf("[internalTxTC] Transaction waiting to be mined: 0x%x\n", mainTx.Hash())

	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := conn.TransactionReceipt(ctx, mainTx.Hash())
		if err != nil {
			log.Printf("[internalTxTC] Failed to check receipt: %v\n", err)
			continue
		}
		log.Println("[internalTxTC] Received the receipt")
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Println("[internalTxTC] Contract deployment was successful")
			break
		} else {
			log.Fatalf("[internalTxTC] Contract Receipt Status: %v\n", receipt.Status)
		}
	}

	return config
}

func readApiCallContractInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// deploy read api call contract
	coinbase := accs[0]
	conn := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonce(conn)))

	var (
		tx           *types.Transaction
		contractAddr common.Address
	)
	log.Println("[TC] readApiCallContract: Deploying new smart contract")

	for {
		var err error
		contractAddr, tx, config.ReadApiCallContract, err = readApiCallContractTC.DeployReadApiCallContractTC(auth, conn)
		if err == nil {
			coinbase.UpdateNonce()
			break
		}
		log.Printf("[TC] readApiCallContract: Failed to deploy new contract: %v\n", err)
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second) // Avoiding Nonce corruption
	}
	log.Printf("[TC] readApiCallContract: Contract address: 0x%x\n", contractAddr)
	log.Printf("[TC] readApiCallContract: Transaction waiting to be mined: 0x%x\n", tx.Hash())

	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond) // Allow it to be processed by the local node :P
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			//fmt.Printf("Failed to check receipt: %v\n", err)
			continue
		}
		log.Printf("=> Contract Receipt Status: %v\n", receipt.Status)
		break
	}

	// set answer variables
	config.RetValOfCall = big.NewInt(4)
	config.RetValOfStorageAt = big.NewInt(4)
	for {
		ctx := context.Background()
		cli := config.CliPool.Alloc().(*client.Client)

		fromAccount := config.AccGrp[rand.Int()%config.NAcc]
		callMsg := kaia.CallMsg{
			From:     fromAccount.GetAddress(),
			To:       &contractAddr,
			Gas:      1100000,
			GasPrice: config.GasPrice,
			Value:    big.NewInt(0),
			Data:     getMethodId("set()"),
		}
		ret, err := cli.EstimateGas(ctx, callMsg)

		if err == nil {
			config.RetValOfEstimateGas = ret
			config.CliPool.Free(cli)
			break
		} else {
			cli.Close()
		}
	}

	return config
}

func storageTrieWriteInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// deploy storage trie write contract
	coinbase := accs[0]
	conn := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(conn)

	auth := bind.NewKeyedTransactor(coinbase.GetKey())
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(coinbase.GetNonce(conn)))

	var (
		address common.Address
		tx      *types.Transaction
		err     error
	)
	println("Deploying a new smart contract")

	for {
		address, tx, config.GUserStorage, err = userStorageTC.DeployUserStorageTC(auth, conn)
		if err != nil {
			coinbase.UpdateNonce()
		}

		if err != nil {
			//log.Printf("Failed to deploy new contract: %v\n", err)
		} else {
			break
		}
		auth.Nonce = big.NewInt(int64(coinbase.GetNonceFromBlock(conn)))
		time.Sleep(1 * time.Second) // Avoiding Nonce corruption
	}
	fmt.Printf("=> Contract pending deploy: 0x%x\n", address)

	fmt.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
	ctx := context.Background()
	defer ctx.Done()
	for {
		time.Sleep(500 * time.Millisecond) // Allow it to be processed by the local node :P
		receipt, err := conn.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			//fmt.Printf("Failed to check receipt: %v\n", err)
			continue
		}
		fmt.Printf("=> Contract Receipt Status: %v\n", receipt.Status)
		break
	}

	return config
}

func ethereumTxInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	// Path to executable file that generates ethereum tx.
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("exPath: ", exPath)

	config.ExecutablePath = exPath + "/ethTxGenerator"
	log.Println("executablePath: ", config.ExecutablePath)

	return config
}

func receiptCheckInit(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig {
	config := tcutil.InitTcConfig(accs, endpoint, gp)

	var hashPoolSize = 100 * 5 * 60 // for init 5min, if input send tps is 100Txs/Sec

	config.HashPool = make([]common.Hash, hashPoolSize, hashPoolSize)
	return config
}
