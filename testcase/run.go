package testcase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	kaia "github.com/kaiachain/kaia"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/testcase/tcutil"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"github.com/kaiachain/kaia/blockchain"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/common"
	"github.com/kaiachain/kaia/networks/rpc"
	"github.com/myzhan/boomer"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/sha3"
)

func transferSignedRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferSignedTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "signedtransfer"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "signedtransfer"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func valueTransferRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)
	defer config.CliPool.Free(cli)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewValueTransferTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewValueTransferTx"+" to "+config.EndPoint, elapsed, int64(10))
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewValueTransferTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newValueTransferWithCancelRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewValueTransferWithCancelTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewValueTransferWithCancelTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewValueTransferWithCancelTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedValueTransferRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedValueTransferTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedValueTransferTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedValueTransferTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedValueTransferWithRatioRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedValueTransferWithRatioTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedValueTransferWithRatioTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedValueTransferWithRatioTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newValueTransferMemoRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewValueTransferMemoTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewValueTransferMemoTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewValueTransferMemoTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newValueTransferLargeMemoRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewValueTransferLargeMemoTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewValueTransferMemoTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewValueTransferMemoTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newValueTransferSmallMemoRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewValueTransferSmallMemoTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewValueTransferMemoTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewValueTransferMemoTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedValueTransferMemoRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedValueTransferMemoTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedValueTransferMemoTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedValueTransferMemoTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedValueTransferMemoWithRatioRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedValueTransferMemoWithRatioTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedValueTransferMemoWithRatioTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedValueTransferMemoWithRatioTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newAccountCreationRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := account.NewKaiaAccount(0)
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewAccountCreationTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewAccountCreationTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewAccountCreationTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newAccountUpdateRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewAccountUpdateTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewAccountUpdateTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewAccountUpdateTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedAccountUpdateRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedAccountUpdateTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedAccountUpdateWithRatioRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedAccountUpdateWithRatioTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newSmartContractDeployRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := account.NewKaiaAccount(0)
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, _, err := from.TransferNewSmartContractDeployTx(cli, to, value, account.TestContractInfos[account.ContractGeneral].Bytecode)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewSmartContractDeployTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewSmartContractDeployTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedSmartContractDeployRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedSmartContractDeployTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedSmartContractDeployWithRatioRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := account.NewKaiaAccount(0)
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedSmartContractDeployWithRatioTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedSmartContractDeployWithRatioTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedSmartContractDeployWithRatioTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newSmartContractExecutionRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.SmartContractAccount
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, to, value, account.TestContractInfos[account.ContractGeneral].GenData(from.GetAddress(), nil))
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewSmartContractExecutionTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewSmartContractExecutionTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedSmartContractExecutionRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedSmartContractExecutionTx(cli, config.SmartContractAccount, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedSmartContractExecutionTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedSmartContractExecutionTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedSmartContractExecutionWithRatioRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.SmartContractAccount
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedSmartContractExecutionWithRatioTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedSmartContractExecutionWithRatioTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedSmartContractExecutionWithRatioTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newCancelRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewCancelTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewCancelTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewCancelTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedCancelRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedCancelTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func newFeeDelegatedCancelWithRatioRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	from := config.AccGrp[rand.Int()%config.NAcc]
	to := config.AccGrp[rand.Int()%config.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.TransferNewFeeDelegatedCancelWithRatioTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "transferNewFeeDelegatedAccountUpdateTx"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func storageTrieWriteRun(tcConfig *tcutil.TcConfig) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	from := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	to := tcConfig.SmartContractAccount
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, _, err := from.ExecuteStorageTrieStore(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "storageTrieWrite"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "storageTrieWrite"+" to "+tcConfig.EndPoint, elapsed, err.Error())
	}
}

func cpuHeavyRun(config *tcutil.TcConfig) {
	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	from := fromAccount.GetKey()

	fromAccount.Lock()
	auth := bind.NewKeyedTransactor(from)

	nonce := fromAccount.GetNonce(conn)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice

	start := boomer.Now()
	_, err := config.GCPUHeavy.SortSingle(auth, big.NewInt(1))
	elapsed := boomer.Now() - start

	if err != nil {
		if err.Error() == blockchain.ErrNonceTooLow.Error() ||
			err.Error() == blockchain.ErrReplaceUnderpriced.Error() {
			fmt.Printf("Account(%v) nonce(%v) : Failed to sendTransaction: %v\n", fromAccount.GetAddress().String(), nonce, err)
			fmt.Printf("Account(%v) nonce is added to %v\n", fromAccount.GetAddress().String(), nonce+1)
			fromAccount.UpdateNonce()
		} else {
			fmt.Printf("Account(%v) nonce(%v) : Failed to sendTransaction: %v\n", fromAccount.GetAddress().String(), nonce, err)
		}

		fmt.Printf("Failed to retrieve pending tx: %v\n", err)
		//fromAccount.GetNonceFromBlock(conn)
	} else {
		//fmt.Println("Pending tx:", res.Hash().String())
		fromAccount.UpdateNonce()
	}

	fromAccount.UnLock()

	if err == nil {
		boomer.Events.Publish("request_success", "contract", "cpuHeavy"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", "cpuHeavy"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func largeMemoRun(config *tcutil.TcConfig) {
	funcName := "largeMemo"

	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	from := fromAccount.GetKey()

	auth := bind.NewKeyedTransactor(from)
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice

	fromAccount.Lock()

	nonce := fromAccount.GetNonce(conn)
	auth.Nonce = big.NewInt(int64(nonce))

	var err error
	var min, max int = 50, 2000

	const Letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// generate random string
	randInt := min + rand.Intn(max-min)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, randInt)
	for i := range b {
		b[i] = Letters[r.Intn(len(Letters))]
	}
	str := string(b)

	start := boomer.Now()
	_, err = config.GLargeMemo.SetName(auth, str)
	elapsed := boomer.Now() - start

	if err != nil {
		log.Printf("[LargeMemo] Failed to call %s(), err=%v\n", funcName, err)
		fromAccount.GetNonceFromBlock(conn)
	} else {
		fromAccount.UpdateNonce()
	}

	fromAccount.UnLock()

	// Uncomment the below for debugging
	//if err == nil {
	//	utils.CheckReceipt(conn, tx.Hash())
	//}

	msg := "LargeMemo" + " to " + config.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		log.Printf("[LargeMemo] request_failure of msg %s, err=%v\n", msg, err)

		conn.Close()
	}
}

// internalTxRun transfers txs calling sendRewards function of mainContract.
// During the execution of the function, four internal transactions are triggered also:
// Mint KIP17 token, Transfer KIP17 token, send KLAY to inviteeAccount and hostAccount.
func internalTxRun(config *tcutil.TcConfig) {
	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	fromAccount.Lock()
	nonce := fromAccount.GetNonce(conn)

	auth := bind.NewKeyedTransactor(fromAccount.GetKey())
	auth.GasLimit = 9999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(100)

	inviteeAccount := config.AccGrp[rand.Int()%config.NAcc]
	hostAccount := config.AccGrp[rand.Int()%config.NAcc]

	start := boomer.Now()
	tx, err := config.MainContract.Transact(auth, "sendRewards", inviteeAccount.GetAddress(), hostAccount.GetAddress())
	if err != nil {
		log.Printf("[internalTxTC] Failed to execute contract, from=%s nonce=%d err=%v\n",
			fromAccount.GetAddress().String(), nonce, err)
		fromAccount.GetNonceFromBlock(conn)
	} else {
		log.Printf("[internalTxTC] tx=%s\n", tx.Hash().String())
		fromAccount.UpdateNonce()
	}
	elapsed := boomer.Now() - start
	fromAccount.UnLock()

	msg := "internalTxTC/" + " to " + config.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		conn.Close()
	}
}

// internalTxMintNFTRun transfers txs calling mintCard function of KIP17Contract.
// The function mints a KIP17 token, NFT, for the sender.
func internalTxMintNFTRun(config *tcutil.TcConfig) {
	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	fromAccount.Lock()
	nonce := fromAccount.GetNonce(conn)

	auth := bind.NewKeyedTransactor(fromAccount.GetKey())
	auth.GasLimit = 9999999
	auth.GasPrice = config.GasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	start := boomer.Now()
	tx, err := config.KIP17Contract.Transact(auth, "mintCard")
	if err != nil {
		log.Printf("[mintNFTTC] Failed to execute contract, from=%s nonce=%d err=%v\n",
			fromAccount.GetAddress().String(), nonce, err)
		fromAccount.GetNonceFromBlock(conn)
	} else {
		log.Printf("[mintNFTTC] tx=%s\n", tx.Hash().String())
		fromAccount.UpdateNonce()
	}
	elapsed := boomer.Now() - start
	fromAccount.UnLock()

	msg := "mintNFTTC/" + " to " + config.EndPoint
	if err == nil {
		boomer.Events.Publish("request_success", "contract", msg, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", msg, elapsed, err.Error())
		conn.Close()
	}
}

func storageTrieWriteSetRun(config *tcutil.TcConfig) {
	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	from := fromAccount.GetKey()

	fromAccount.Lock()
	auth := bind.NewKeyedTransactor(from)

	nonce := fromAccount.GetNonce(conn)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice

	start := boomer.Now()
	_, err := config.GUserStorage.Set(auth, big.NewInt(100))
	elapsed := boomer.Now() - start

	if err != nil {
		fmt.Printf("Failed to retrieve pending tx: %v\n", err)
		fromAccount.GetNonceFromBlock(conn)
	} else {
		fromAccount.UpdateNonce()
	}

	fromAccount.UnLock()

	if err == nil {
		boomer.Events.Publish("request_success", "contract", "userStorage/Set"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", "userStorage/Set"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func storageTrieWriteSetGetRun(config *tcutil.TcConfig) {
	conn := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	from := fromAccount.GetKey()

	fromAccount.Lock()

	start := boomer.Now()

	// Call Set()
	auth := bind.NewKeyedTransactor(from)
	auth.Nonce = big.NewInt(int64(fromAccount.GetNonce(conn)))
	auth.GasLimit = 999999
	auth.GasPrice = config.GasPrice

	value := auth.Nonce

	tx, err := config.GUserStorage.Set(auth, value)
	elapsed := boomer.Now() - start

	if err != nil {
		fmt.Printf("Failed to retrieve pending tx: %v\n", err)
		fromAccount.GetNonceFromBlock(conn)

		fromAccount.UnLock()

		boomer.Events.Publish("request_failure", "contract", "userStorage/SetGet"+" to "+config.EndPoint, elapsed, err.Error())

		return
	}

	start = boomer.Now()

	// Increment fromAccount's nonce
	fromAccount.UpdateNonce()

	ctx := context.Background()
	defer ctx.Done()

	// Wait for the receipt to be available
	for {
		time.Sleep(500 * time.Millisecond)
		receipt, receiptErr := conn.TransactionReceipt(ctx, tx.Hash())
		if receiptErr != nil {
			continue
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			err = errors.New(fmt.Sprintf("tx=%v: from=%v failed=%v", tx.Hash().String(), fromAccount.GetAddress().String(), receipt.Status))
			fmt.Println(err.Error())
		}
		break
	}

	if err == nil {
		time.Sleep(1500 * time.Millisecond)

		// Wait until tx is included in the block
		for {
			_, isPending, _ := conn.TransactionByHash(ctx, tx.Hash())
			if isPending {
				time.Sleep(5 * time.Millisecond)
			} else {
				break
			}
		}

		// Call Get() to retrieve the value set by Set()
		var callopts bind.CallOpts
		callopts.Pending = false
		callopts.From = fromAccount.GetAddress()
		result, getErr := config.GUserStorage.Get(&callopts)
		if getErr != nil {
			err = getErr
		} else if result.Cmp(value) != 0 {
			err = errors.New(fmt.Sprintf("tx=%v: from=%v, incorrect value (received=%v vs. expected=%v)", tx.Hash().String(), callopts.From.String(), result, value))
			fmt.Println(err.Error())
		}
	}

	elapsed += boomer.Now() - start

	fromAccount.UnLock()

	if err == nil {
		boomer.Events.Publish("request_success", "contract", "userStorage/SetGet"+" to "+config.EndPoint, elapsed, int64(10))
		config.CliPool.Free(conn)
	} else {
		boomer.Events.Publish("request_failure", "contract", "userStorage/SetGet"+" to "+config.EndPoint, elapsed, err.Error())
	}
}

func readApiCallContractGetStorageAtRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	cli := config.CliPool.Alloc().(*client.Client)

	start := boomer.Now()
	ret, err := cli.StorageAt(ctx, config.SmartContractAccount.GetAddress(), common.Hash{}, nil)
	elapsed := boomer.Now() - start

	if err == nil && new(big.Int).SetBytes(ret).Cmp(config.RetValOfStorageAt) != 0 {
		err = errors.New("wrong storage value: " + string(ret) + ", answer: " + config.RetValOfStorageAt.String())
	}
	sendBoomerEvent("readGetStorageAt", "Failure to call klay_getStorageAt", elapsed, cli, config, err)
}

func readApiCallContractCallRun(config *tcutil.TcConfig) {
	cli := config.CliPool.Alloc().(*client.Client)

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	var callopts bind.CallOpts
	callopts.Pending = false
	callopts.From = fromAccount.GetAddress()

	start := boomer.Now()
	ret, err := config.ReadApiCallContract.Get(&callopts)
	elapsed := boomer.Now() - start

	if err == nil && ret.Cmp(config.RetValOfCall) != 0 {
		err = errors.New("wrong call: " + ret.String() + ", answer: " + config.RetValOfCall.String())
	}
	sendBoomerEvent("readCall", "Failed to call klay_call", elapsed, cli, config, err)
}

func readApiCallContractEstimateGasRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	cli := config.CliPool.Alloc().(*client.Client)
	to := config.SmartContractAccount.GetAddress()

	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	callMsg := kaia.CallMsg{
		From:     fromAccount.GetAddress(),
		To:       &to,
		Gas:      1100000,
		GasPrice: config.GasPrice,
		Value:    big.NewInt(0),
		Data:     getMethodId("set()"),
	}
	start := boomer.Now()
	ret, err := cli.EstimateGas(ctx, callMsg)
	elapsed := boomer.Now() - start

	if err == nil && ret != config.RetValOfEstimateGas {
		err = errors.New("wrong estimate gas: " + strconv.Itoa(int(ret)) + ", answer: " + strconv.Itoa(int(config.RetValOfEstimateGas)))
	}
	sendBoomerEvent("readEstimateGas", "Failed to call klay_estimateGas", elapsed, cli, config, err)
}

func getMethodId(str string) []byte {
	transferFnSignature := []byte(str)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	return methodID
}

func sendBoomerEvent(tcName string, logString string, elapsed int64, cli *client.Client, tcConfig *tcutil.TcConfig, err error) {
	if err == nil {
		boomer.Events.Publish("request_success", "http", tcName+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		log.Printf("[TC] %s: %s, err=%v\n", tcName, logString, err)
		boomer.Events.Publish("request_failure", "http", tcName+" to "+tcConfig.EndPoint, elapsed, err.Error())
		cli.Close()
	}
}

func readApiCallGasPriceRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	cli := config.CliPool.Alloc().(*client.Client)

	start := boomer.Now()
	gp, err := cli.SuggestGasPrice(ctx)
	elapsed := boomer.Now() - start
	if err == nil && gp.Cmp(config.GasPrice) != 0 {
		err = errors.New("wrong gas price: " + gp.String() + ", answer: " + config.GasPrice.String())
	}
	sendBoomerEvent("readGasPrice", "Failed to call klay_gasPrice", elapsed, cli, config, err)
}

func readApiCallBlockNumberRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	cli := config.CliPool.Alloc().(*client.Client)

	start := boomer.Now()

	bn, err := cli.BlockNumber(ctx)
	if err == nil && bn.Cmp(big.NewInt(0)) != 1 {
		err = errors.New("wrong block number: 0x" + bn.Text(16) + ", answer: smaller than 0")
	}

	elapsed := boomer.Now() - start
	sendBoomerEvent("readBlockNumber", "Failed to call klay_blockNumber", elapsed, cli, config, err)
}

func readApiCallGetBlockByNumberRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	cli := config.CliPool.Alloc().(*client.Client)

	ansBN := getRandomBlockNumber(cli, ctx, config)
	start := boomer.Now()

	block, err := cli.BlockByNumber(ctx, ansBN) //read the random block
	if err == nil && block.Header().Number.Cmp(ansBN) != 0 {
		err = errors.New("wrong block: 0x" + block.Header().Number.Text(16) + ", answer: 0x" + ansBN.Text(16))
	}

	elapsed := boomer.Now() - start
	sendBoomerEvent("readGetBlockByNumber", "Failed to call klay_getBlockByNumber", elapsed, cli, config, err)
}

func readApiCallGetAccountRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	rpcCli := config.CliPool.Alloc().(*rpc.Client)
	cli := client.NewClient(rpcCli)

	var j json.RawMessage
	fromAccount := config.AccGrp[rand.Int()%config.NAcc]
	start := boomer.Now()

	err := rpcCli.CallContext(ctx, &j, "klay_getAccount", fromAccount.GetAddress(), "latest")
	if err == nil {
		ret := gjson.Get(string(j), "accType").String()
		if ret != "1" {
			err = errors.New("wrong account type: " + ret + ", answer: 1")
		}
	}

	elapsed := boomer.Now() - start
	sendBoomerEvent("readGetAccount", "Failed to call klay_getAccount", elapsed, cli, config, err)
}

func readApiCallGetBlockWithConsensusInfoByNumberRun(config *tcutil.TcConfig) {
	ctx := context.Background()
	rpcCli := config.CliPool.Alloc().(*rpc.Client)
	cli := client.NewClient(rpcCli)

	ansBN := getRandomBlockNumber(cli, ctx, config)
	start := boomer.Now()

	var j json.RawMessage
	err := rpcCli.CallContext(ctx, &j, "klay_getBlockWithConsensusInfoByNumber", "0x"+ansBN.Text(16))
	if err == nil {
		ret := gjson.Get(string(j), "number").String()
		if !strings.Contains(ret, "0x"+ansBN.Text(16)) {
			err = errors.New("wrong block: " + ret + ", answer: " + "0x" + ansBN.Text(16))
		}
	}

	elapsed := boomer.Now() - start
	sendBoomerEvent("readGetBlockWithConsensusInfoByNumber",
		"Failed to call klay_GetBlockWithConsensusInfoByNumber", elapsed, cli, config, err)
}

func getRandomBlockNumber(cli *client.Client, ctx context.Context, tcConfig *tcutil.TcConfig) *big.Int {
	tcConfig.Count %= 10000000
	if tcConfig.Count%10000 == 0 {
		bn, err := cli.BlockNumber(ctx)
		if err != nil {
			log.Printf("Failed to update the current block number. err=%s\n", err)
		} else {
			log.Printf("Update the current block number. blockNumber=0x%s\n", bn.Text(16))
			tcConfig.LatestBlockNumber.Set(bn)
		}
	}
	tcConfig.Count += 1

	return big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), tcConfig.LatestBlockNumber)
}

func erc20TransferRun(tcConfig *tcutil.TcConfig) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	from := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	to := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))
	data := account.TestContractInfos[account.ContractErc20].GenData(to.GetAddress(), value)

	start := boomer.Now()
	_, _, err := from.TransferNewSmartContractExecutionTx(cli, tcConfig.SmartContractAccount, nil, data)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "erc20Transfer"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "erc20Transfer"+" to "+tcConfig.EndPoint, elapsed, err.Error())
	}
}

func erc721TransferRun(tcConfig *tcutil.TcConfig) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	fromAcc := tcConfig.AccGrp[rand.Intn(tcConfig.NAcc)]
	toAcc := tcConfig.AccGrp[rand.Intn(tcConfig.NAcc)]

	// Get token ID from the channel
	// Here is an assumption that it won't be blocked by the channel
	// Although this go routine can be blocked, other can send a NFT to this account
	fromNFTs := account.ERC721Ledger[fromAcc.GetAddress()]
	tokenId := <-fromNFTs

	start := boomer.Now()
	_, _, err := fromAcc.TransferERC721(false, cli, tcConfig.SmartContractAccount.GetAddress(), toAcc, tokenId)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "erc721TransferTC"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
		toNFTs := account.ERC721Ledger[toAcc.GetAddress()]
		toNFTs <- tokenId // push the token to the new owner's queue, it it does not fail

	} else {
		boomer.Events.Publish("request_failure", "http", "erc721TransferTC"+" to "+tcConfig.EndPoint, elapsed, err.Error())
		fromNFTs <- tokenId // push back to the original owner, if it fails
	}
}

func doubleLock(to *account.Account, from *account.Account) {
	if from.GetAddress().String() == to.GetAddress().String() {
		from.Lock()
	} else if from.GetAddress().String() > to.GetAddress().String() {
		from.Lock()
		to.Lock()
	} else {
		to.Lock()
		from.Lock()
	}
}

func doubleUnlock(to *account.Account, from *account.Account) {
	if from.GetAddress().String() == to.GetAddress().String() {
		from.UnLock()
	} else if from.GetAddress().String() > to.GetAddress().String() {
		from.UnLock()
		to.UnLock()
	} else {
		to.UnLock()
		from.UnLock()
	}
}

func TransferAndCheck(cli *client.Client, to *account.Account, from *account.Account, value *big.Int) error {
	ctx := context.Background()

	doubleLock(to, from)
	defer doubleUnlock(to, from)
	// The reason of saving balance of current accounts is to comparing with later balance.
	fromFormerBalance, _ := from.GetBalance(cli)
	toFormerBalance, _ := to.GetBalance(cli)

	hash, gasPrice, err := from.TransferSignedTxWithoutLock(cli, to, value)
	if err != nil {
		return err
	}
	startTime := time.Now().Unix()
	var receipt *types.Receipt
	for {
		receipt, _ = cli.TransactionReceipt(ctx, hash)
		if receipt != nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
		if time.Now().Unix()-startTime > 100 {
			return errors.New("Time out : It took more than 100 seconds to make a block ")
		}
	}

	if to.GetAddress() == from.GetAddress() {
		value.SetUint64(0)
	}

	fromFormerBalance.Sub(fromFormerBalance, value)
	gasUsed := big.NewInt((int64)(receipt.GasUsed))
	fee := new(big.Int).Mul(gasUsed, gasPrice)
	fromFormerBalance.Sub(fromFormerBalance, fee)
	toFormerBalance.Add(toFormerBalance, value)

	startTime = time.Now().Unix()
	for {
		errFrom := from.CheckBalance(fromFormerBalance, cli)
		if errFrom != nil {
			log.Printf("from account : %s", errFrom.Error())
			time.Sleep(100 * time.Millisecond)
			if time.Now().Unix()-startTime > 10 {
				return errors.New("Time out (from) : It took more than 10 seconds to retrieve the correct receipt ")
			}
		} else {
			break
		}
	}

	if from.GetAddress() == to.GetAddress() {
		return nil
	}

	startTime = time.Now().Unix()
	for {
		errTo := to.CheckBalance(toFormerBalance, cli)
		if errTo != nil {
			log.Printf("to account : %s", errTo.Error())
			time.Sleep(100 * time.Millisecond)
			if time.Now().Unix()-startTime > 10 {
				return errors.New("Time out (to) : It took more than 10 seconds to retrieve the correct receipt ")
			}
		} else {
			break
		}
	}

	return nil
}

func transferSignedWithCheckRun(tcConfig *tcutil.TcConfig) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	from := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	to := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]

	value := big.NewInt(int64(rand.Int() % 3))
	start := boomer.Now()

	err := TransferAndCheck(cli, to, from, value)

	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "signedtransfer_with_check"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "signedtransfer_with_check"+" to "+tcConfig.EndPoint, elapsed, err.Error())
	}
}

func transferUnsignedRun(tcConfig *tcutil.TcConfig) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	from := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	to := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	_, err := from.TransferUnsignedTx(cli, to, value)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", "unsignedtransfer"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", "http", "unsignedtransfer"+" to "+tcConfig.EndPoint, elapsed, err.Error())
	}
}

func receiptCheckRun(tcConfig *tcutil.TcConfig) {
	var (
		defaultInitSendTx = 1000 * 10 // for init 10sec, if input send TPS is 1000Txs/Sec
		ratioReadPerSend  = 9
		name              = "receiptCheckTx"

		cnt      uint32
		initFlag = false
	)

	nc := atomic.AddUint32(&cnt, 1)

	if !initFlag && nc < uint32(defaultInitSendTx) {
		receiptCheckSendTx(tcConfig, name)
	} else {
		initFlag = true

		// following logic can control the ratio between send/read task
		nc = nc % uint32(ratioReadPerSend+1)

		if nc == uint32(ratioReadPerSend) {
			receiptCheckSendTx(tcConfig, name)
		} else {
			receiptCheckReadTx(tcConfig, name)
		}
	}
}

func receiptCheckSendTx(tcConfig *tcutil.TcConfig, name string) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	from := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	to := tcConfig.AccGrp[rand.Int()%tcConfig.NAcc]
	value := big.NewInt(int64(rand.Int() % 3))

	start := boomer.Now()
	hash, _, err := from.TransferSignedTx(cli, to, value)
	tcConfig.AddHash(hash)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", name, "send tx"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", name, "send tx"+" to "+tcConfig.EndPoint, elapsed, err.Error())
	}
}

func receiptCheckReadTx(tcConfig *tcutil.TcConfig, name string) {
	cli := tcConfig.CliPool.Alloc().(*client.Client)

	ctx := context.Background()
	hash := tcConfig.GetHash()

	start := boomer.Now()

	receipt, err := cli.TransactionReceipt(ctx, hash)
	if err == nil {
		if rand.Int()%(1000*60) == 0 {
			log.Printf("pid(%v) : hash(%v) receipt checked\n", os.Getpid(), hash.String())
			log.Printf("%v", receipt)
		}
	} else {
		log.Printf("pid(%v) : hash(%v) receipt check err : %v\n", os.Getpid(), hash.String(), err)
	}

	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", name, "read tx"+" to "+tcConfig.EndPoint, elapsed, int64(10))
		tcConfig.CliPool.Free(cli)
	} else {
		boomer.Events.Publish("request_failure", name, "read tx"+" to "+tcConfig.EndPoint, elapsed, err.Error())
	}
}
