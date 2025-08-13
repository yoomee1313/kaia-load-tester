package refactored_closure

import (
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/client"
)

const NewValueTransferMemoTCName = "newValueTransferMemoTC"

var newValueTransferMemoConfig *TCConfig

func InitNewValueTransferMemoTC(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	newValueTransferMemoConfig = Init(accs, contractsParam, endpoint, gp)
}

func RunValueTransferMemo() {
	RunBaseValueTransfer(newValueTransferMemoConfig, func(cli *client.Client, from *account.Account, to *account.Account) (interface{}, *big.Int, error) {
		value := big.NewInt(int64(rand.Int() % 3))
		return from.TransferNewValueTransferMemoTx(cli, to, value)
	}, "transferNewValueTransferMemoTx")
}
