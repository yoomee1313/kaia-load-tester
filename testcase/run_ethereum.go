package testcase

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	kaia "github.com/kaiachain/kaia"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/accounts/abi"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
	"github.com/kaiachain/kaia/common/hexutil"
	"github.com/myzhan/boomer"
)

// Ethereum transaction related variables
var (
	// Smart contract bytecode for deployment
	ethereumContractCode = "0x608060405234801561001057600080fd5b506101de806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631a39d8ef81146100805780636353586b146100a757806370a08231146100ca578063fd6b7ef8146100f8575b3360009081526001602052604081208054349081019091558154019055005b34801561008c57600080fd5b5061009561010d565b60408051918252519081900360200190f35b6100c873ffffffffffffffffffffffffffffffffffffffff60043516610113565b005b3480156100d657600080fd5b5061009573ffffffffffffffffffffffffffffffffffffffff60043516610147565b34801561010457600080fd5b506100c8610159565b60005481565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604081208054349081019091558154019055565b60016020526000908152604090205481565b336000908152600160205260408120805490829055908111156101af57604051339082156108fc029083906000818181858888f193505050501561019c576101af565b3360009081526001602052604090208190555b505600a165627a7a72305820627ca46bb09478a015762806cc00c431230501118c7c26c30ac58c4e09e51c4f0029"

	// Maximum retry count for getting transaction receipt
	maxRetryCount = 30

	// Executable path for ethereum transaction generator
	executablePath string
)

// initEthereum initializes ethereum related variables
func initEthereum(config *TCConfig) {
	// Path to executable file that generates ethereum tx.
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("exPath: ", exPath)

	executablePath = exPath + "/ethTxGenerator"
	log.Println("executablePath: ", executablePath)
}

// createRandomArguments generates arguments randomly with various cases.
// simple value transfer, smart contract deployment, smart contract execution
func createRandomArguments(config *TCConfig, addr common.Address) (*account.Account, *big.Int, string, int, error) {
	// randomLegacyReqType == 0 : Value transfer
	// randomLegacyReqType == 1 : Smart contract deployment
	// randomLegacyReqType == 2 : Smart contract execution
	randomLegacyReqType := rand.Int() % 3

	var to *account.Account
	var value *big.Int
	input := ""

	var err error
	if randomLegacyReqType == 0 {
		to = config.AccGrp[rand.Int()%config.NAcc]
		value = big.NewInt(int64(rand.Int() % 3))
	} else if randomLegacyReqType == 1 {
		value = big.NewInt(0)
		input = ethereumContractCode
	} else if randomLegacyReqType == 2 {
		to = config.SmartContractAccounts[account.ContractGeneral]
		value = big.NewInt(0)
		input, err = makeFunctionCall(addr)
		if err != nil {
			return nil, nil, "", randomLegacyReqType, err
		}
	} else {
		return nil, nil, "", randomLegacyReqType, err
	}

	return to, value, input, randomLegacyReqType, nil
}

// makeFunctionCall returns a function call to execute smart contract.
func makeFunctionCall(addr common.Address) (string, error) {
	abiStr := `[{"constant":true,"inputs":[],"name":"totalAmount","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"receiver","type":"address"}],"name":"reward","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"safeWithdrawal","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"}]`
	abii, err := abi.JSON(strings.NewReader(string(abiStr)))
	if err != nil {
		log.Fatalf("failed to abi.JSON: %v", err)
		return "", err
	}
	data, err := abii.Pack("reward", addr)
	if err != nil {
		log.Fatalf("failed to abi.Pack: %v", err)
		return "", err
	}
	return hex.EncodeToString(data), nil
}

// getReceipt returns a transaction receipt.
// If receipt is nil, retry until maxRetry.
func getReceipt(cli *client.Client, txHash common.Hash, maxRetry int) map[string]interface{} {
	ctx := context.Background()
	defer ctx.Done()
	retryCount := 0

	for {
		time.Sleep(500 * time.Millisecond)
		receipt, err := cli.TransactionReceiptRpcOutput(ctx, txHash)
		if receipt != nil {
			return receipt
		}
		if err != nil {
			if err.Error() == kaia.NotFound.Error() && retryCount < maxRetry {
				retryCount++
				continue
			}
			fmt.Printf("return nil because receipt(%v) is not notFound: %v, maxRetry: %v \n", txHash.String(), retryCount, maxRetry)
			return nil
		}
	}
}

// checkResult returns true and nil error, if expected results are observed, otherwise returns false and error.
func checkResult(cli *client.Client, txHash common.Hash, reqType int, config *TCConfig, expectedTxType types.TxType) (bool, error) {
	receipt := getReceipt(cli, txHash, maxRetryCount)

	if receipt == nil {
		return false, errors.New("failed to get transaction receipt")
	}

	status, _ := receipt["status"].(string)
	if status != hexutil.Uint(types.ReceiptStatusSuccessful).String() {
		fmt.Printf("[FAILED] TxHash=%v, Receipt status=%v, Tx error msg=%v\n", txHash.String(), status, receipt["txError"])
		return false, errors.New("transaction status in receipt is fail")
	}

	// Check smart contract related fields
	if reqType == 1 {
		contractAddress, _ := receipt["contractAddress"]
		_, ok := contractAddress.(string)
		if !ok {
			return false, errors.New("failed to get contract address from the receipt")
		}
	} else if reqType == 2 {
		toFromReceipt, ok := receipt["to"].(string)
		if !ok || strings.ToLower(toFromReceipt) != strings.ToLower(config.SmartContractAccounts[account.ContractGeneral].GetAddress().String()) {
			return false, errors.New("mismatched to address in the receipt and smart contract address")
		}
	}

	// Check transaction type for non-legacy transactions
	if expectedTxType != types.TxTypeLegacyTransaction {
		typeString, ok := receipt["type"].(string)
		if !ok {
			return false, errors.New("failed to get type from the receipt")
		}
		if typeString != expectedTxType.String() {
			return false, errors.New("type mismatched in transaction receipt")
		}

		typeInt, ok := receipt["typeInt"].(float64)
		if !ok {
			return false, errors.New("failed to get typeInt from the receipt")
		}
		if types.TxType(typeInt) != expectedTxType {
			return false, errors.New("typeInt mismatched in transaction receipt")
		}
	}

	return true, nil
}

// RunEthereumTxLegacyTC creates a closure for ethereum legacy transaction test case
func RunEthereumTxLegacyTC(config *TCConfig) func() {
	return func() {
		cli := config.CliPool.Alloc().(*client.Client)
		defer config.CliPool.Free(cli)

		from := config.AccGrp[rand.Int()%config.NAcc]
		to, value, input, reqType, err := createRandomArguments(config, from.GetAddress())
		if err != nil {
			fmt.Printf("Failed to create arguments to send Legacy Tx: %v\n", err.Error())
			return
		}

		start := boomer.Now()

		txHash, _, err := from.TransferNewLegacyTxWithEth(cli, config.EndPoint, to, value, input, executablePath)

		elapsed := boomer.Now() - start

		if err != nil {
			boomer.Events.Publish("request_failure", "http", "TransferNewLegacyTx"+" to "+config.EndPoint, elapsed, err.Error())
			return
		}

		// Check test result with checkResult function
		go func(transactionHash common.Hash) {
			ret, err := checkResult(cli, transactionHash, reqType, config, types.TxTypeLegacyTransaction)
			if ret == false || err != nil {
				boomer.Events.Publish("request_failure", "http", "TransferNewLegacyTx"+" to "+config.EndPoint, elapsed, err.Error())
				return
			}

			boomer.Events.Publish("request_success", "http", "TransferNewLegacyTx"+" to "+config.EndPoint, elapsed, int64(10))
		}(txHash)
	}
}

// RunEthereumTxAccessListTC creates a closure for ethereum access list transaction test case
func RunEthereumTxAccessListTC(config *TCConfig) func() {
	return func() {
		cli := config.CliPool.Alloc().(*client.Client)
		defer config.CliPool.Free(cli)

		from := config.AccGrp[rand.Int()%config.NAcc]
		to, value, input, reqType, err := createRandomArguments(config, from.GetAddress())
		if err != nil {
			fmt.Printf("Failed to create arguments to send Access List Tx: %v\n", err.Error())
			return
		}

		start := boomer.Now()

		txHash, _, err := from.TransferNewEthAccessListTxWithEth(cli, config.EndPoint, to, value, input, executablePath)

		elapsed := boomer.Now() - start

		if err != nil {
			boomer.Events.Publish("request_failure", "http", "TransferNewEthAccessListTx"+" to "+config.EndPoint, elapsed, err.Error())
			return
		}

		// Check test result with checkResult function
		go func(transactionHash common.Hash) {
			ret, err := checkResult(cli, transactionHash, reqType, config, types.TxTypeEthereumAccessList)
			if ret == false || err != nil {
				boomer.Events.Publish("request_failure", "http", "TransferNewEthAccessListTx"+" to "+config.EndPoint, elapsed, err.Error())
				return
			}

			boomer.Events.Publish("request_success", "http", "TransferNewEthAccessListTx"+" to "+config.EndPoint, elapsed, int64(10))
		}(txHash)
	}
}

// RunEthereumTxDynamicFeeTC creates a closure for ethereum dynamic fee transaction test case
func RunEthereumTxDynamicFeeTC(config *TCConfig) func() {
	return func() {
		cli := config.CliPool.Alloc().(*client.Client)
		defer config.CliPool.Free(cli)

		from := config.AccGrp[rand.Int()%config.NAcc]
		to, value, input, reqType, err := createRandomArguments(config, from.GetAddress())
		if err != nil {
			fmt.Printf("Failed to create arguments to send Dynamic Fee Tx: %v\n", err.Error())
			return
		}

		start := boomer.Now()

		txHash, _, err := from.TransferNewEthDynamicFeeTxWithEth(cli, config.EndPoint, to, value, input, executablePath)

		elapsed := boomer.Now() - start

		if err != nil {
			boomer.Events.Publish("request_failure", "http", "TransferNewEthDynamicFeeTx"+" to "+config.EndPoint, elapsed, err.Error())
			return
		}

		// Check test result with checkResult function
		go func(transactionHash common.Hash) {
			ret, err := checkResult(cli, transactionHash, reqType, config, types.TxTypeEthereumDynamicFee)
			if ret == false || err != nil {
				boomer.Events.Publish("request_failure", "http", "TransferNewEthDynamicFeeTx"+" to "+config.EndPoint, elapsed, err.Error())
				return
			}

			boomer.Events.Publish("request_success", "http", "TransferNewEthDynamicFeeTx"+" to "+config.EndPoint, elapsed, int64(10))
		}(txHash)
	}
}

// RunNewEthereumAccessListTC creates a closure for new ethereum access list transaction test case
func RunNewEthereumAccessListTC(config *TCConfig) func() {
	return func() {
		cli := config.CliPool.Alloc().(*client.Client)
		defer config.CliPool.Free(cli)

		from := config.AccGrp[rand.Int()%config.NAcc]
		to, value, input, reqType, err := createRandomArguments(config, from.GetAddress())
		if err != nil {
			fmt.Printf("Failed to create arguments to send Access List Tx: %v\n", err.Error())
			return
		}

		start := boomer.Now()
		txHash, _, err := from.TransferNewEthereumAccessListTx(cli, to, value, common.FromHex(input))
		elapsed := boomer.Now() - start

		if err != nil {
			boomer.Events.Publish("request_failure", "http", "transferNewEthereumAccessListTx"+" to "+config.EndPoint, elapsed, err.Error())
			return
		}

		// Check test result with checkResult function
		go func(transactionHash common.Hash) {
			ret, err := checkResult(cli, transactionHash, reqType, config, types.TxTypeEthereumAccessList)
			if ret == false || err != nil {
				boomer.Events.Publish("request_failure", "http", "transferNewEthereumAccessListTx"+" to "+config.EndPoint, elapsed, err.Error())
				return
			}

			boomer.Events.Publish("request_success", "http", "transferNewEthereumAccessListTx"+" to "+config.EndPoint, elapsed, int64(10))
		}(txHash)
	}
}

// RunNewEthereumDynamicFeeTC creates a closure for new ethereum dynamic fee transaction test case
func RunNewEthereumDynamicFeeTC(config *TCConfig) func() {
	return func() {
		cli := config.CliPool.Alloc().(*client.Client)
		defer config.CliPool.Free(cli)

		from := config.AccGrp[rand.Int()%config.NAcc]
		to, value, input, reqType, err := createRandomArguments(config, from.GetAddress())
		if err != nil {
			fmt.Printf("Failed to create arguments to send Dynamic Fee Tx: %v\n", err.Error())
			return
		}

		start := boomer.Now()
		txHash, _, err := from.TransferNewEthereumDynamicFeeTx(cli, to, value, common.FromHex(input))
		elapsed := boomer.Now() - start

		if err != nil {
			boomer.Events.Publish("request_failure", "http", "transferNewEthereumDynamicFeeTx"+" to "+config.EndPoint, elapsed, err.Error())
			return
		}

		// Check test result with checkResult function
		go func(transactionHash common.Hash) {
			ret, err := checkResult(cli, transactionHash, reqType, config, types.TxTypeEthereumDynamicFee)
			if ret == false || err != nil {
				boomer.Events.Publish("request_failure", "http", "transferNewEthereumDynamicFeeTx"+" to "+config.EndPoint, elapsed, err.Error())
				return
			}

			boomer.Events.Publish("request_success", "http", "transferNewEthereumDynamicFeeTx"+" to "+config.EndPoint, elapsed, int64(10))
		}(txHash)
	}
}
