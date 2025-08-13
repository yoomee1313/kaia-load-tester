//go:generate abigen --sol ReadApiCallContract.sol --pkg readApiCallContract --out ReadApiCallContract.go
package readApiCallContractTC

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"sync"

	kaia "github.com/kaiachain/kaia"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/accounts/abi"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
	"github.com/myzhan/boomer"
	"golang.org/x/crypto/sha3"
)

var (
	endPoint string
	cliPool  clipool.ClientPool

	mutex       sync.Mutex
	initialized = false

	nAcc   int
	accGrp []*account.Account

	gasPrice *big.Int

	SmartContractAccount *account.Account

	retValOfCall        *big.Int
	retValOfStorageAt   *big.Int
	retValOfEstimateGas uint64
)

func Init(accs []*account.Account, contractsParam []*account.Account, ep string, gp *big.Int) {
	mutex.Lock()
	defer mutex.Unlock()

	if initialized {
		return
	}
	initialized = true

	gasPrice = gp
	endPoint = ep
	SmartContractAccount = contractsParam[account.ContractReadApiCallContract]
	cliCreate := func() interface{} {
		c, err := client.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}
	cliPool.Init(20, 300, cliCreate)

	for _, acc := range accs {
		accGrp = append(accGrp, acc)
	}
	nAcc = len(accGrp)
	fmt.Println("setAnswerVariables")
	setAnswerVariables()
}

func getMethodId(str string) []byte {
	transferFnSignature := []byte(str)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	return methodID
}

// TODO-kaia-load-tester: deleting for loop
func setAnswerVariables() {
	retValOfCall = big.NewInt(4)
	retValOfStorageAt = big.NewInt(4)

	// Check if SmartContractAccount is available
	if SmartContractAccount == nil {
		log.Printf("[TC] readApiCallContract: SmartContractAccount is nil, skipping setAnswerVariables")
		return
	}

	fmt.Printf("[TC] readApiCallContract: SmartContractAccount address: %s\n", SmartContractAccount.GetAddress().String())

	for {
		ctx := context.Background()
		cli := cliPool.Alloc().(*client.Client)

		fromAccount := accGrp[rand.Int()%nAcc]
		contractAddr := SmartContractAccount.GetAddress()
		data := account.TestContractInfos[account.ContractReadApiCallContract].GenData(fromAccount.GetAddress(), big.NewInt(1))

		callMsg := kaia.CallMsg{
			From:     fromAccount.GetAddress(),
			To:       &contractAddr,
			Gas:      1100000,
			GasPrice: gasPrice,
			Value:    big.NewInt(0),
			Data:     data,
		}
		ret, err := cli.EstimateGas(ctx, callMsg)
		cliPool.Free(cli)
		if err == nil {
			retValOfEstimateGas = ret
			break
		}
	}
}

func sendBoomerEvent(tcName string, logString string, elapsed int64, cli *client.Client, err error) {
	if err == nil {
		boomer.Events.Publish("request_success", "http", tcName+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", tcName+" to "+endPoint, elapsed, err.Error())
	}
}

func GetStorageAt() {
	ctx := context.Background()
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	contractAddr := SmartContractAccount.GetAddress()
	start := boomer.Now()
	ret, err := cli.StorageAt(ctx, contractAddr, common.Hash{}, nil)
	elapsed := boomer.Now() - start

	if err == nil && new(big.Int).SetBytes(ret).Cmp(retValOfStorageAt) != 0 {
		err = errors.New("wrong storage value: " + string(ret) + ", answer: " + retValOfStorageAt.String())
	}
	sendBoomerEvent("readGetStorageAt", "Failure to call klay_getStorageAt", elapsed, cli, err)
}

func Call() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	// Check if SmartContractAccount is available
	if SmartContractAccount == nil {
		sendBoomerEvent("readCall", "SmartContractAccount is nil", 0, cli, errors.New("SmartContractAccount is nil"))
		return
	}

	fromAccount := accGrp[rand.Int()%nAcc]
	contractAddr := SmartContractAccount.GetAddress()

	// Use GenData to get get() function data (value = 0 for get function)
	data := account.TestContractInfos[account.ContractReadApiCallContract].GenData(fromAccount.GetAddress(), big.NewInt(0))

	callMsg := kaia.CallMsg{
		From: fromAccount.GetAddress(),
		To:   &contractAddr,
		Data: data,
	}

	start := boomer.Now()
	result, err := cli.CallContract(context.Background(), callMsg, nil)
	elapsed := boomer.Now() - start

	if err == nil {
		// Parse the result using the same ABI
		abiStr := `[{"inputs":[],"name":"get","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
		parsedABI, err := abi.JSON(strings.NewReader(abiStr))
		if err == nil {
			var ret *big.Int
			err = parsedABI.UnpackIntoInterface(&ret, "get", result)
			if err == nil && ret.Cmp(retValOfCall) != 0 {
				err = errors.New("wrong call: " + ret.String() + ", answer: " + retValOfCall.String())
			}
		}
	}
	sendBoomerEvent("readCall", "Failed to call klay_call", elapsed, cli, err)
}

func EstimateGas() {
	ctx := context.Background()
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	// Check if SmartContractAccount is available
	if SmartContractAccount == nil {
		sendBoomerEvent("readEstimateGas", "SmartContractAccount is nil", 0, cli, errors.New("SmartContractAccount is nil"))
		return
	}

	fromAccount := accGrp[rand.Int()%nAcc]
	contractAddr := SmartContractAccount.GetAddress()
	data := account.TestContractInfos[account.ContractReadApiCallContract].GenData(fromAccount.GetAddress(), big.NewInt(1))

	callMsg := kaia.CallMsg{
		From:     fromAccount.GetAddress(),
		To:       &contractAddr,
		Gas:      1100000,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     data,
	}
	start := boomer.Now()
	ret, err := cli.EstimateGas(ctx, callMsg)
	elapsed := boomer.Now() - start

	if err == nil && ret != retValOfEstimateGas {
		err = errors.New("wrong estimate gas: " + strconv.Itoa(int(ret)) + ", answer: " + strconv.Itoa(int(retValOfEstimateGas)))
	}
	sendBoomerEvent("readEstimateGas", "Failed to call klay_estimateGas", elapsed, cli, err)
}
