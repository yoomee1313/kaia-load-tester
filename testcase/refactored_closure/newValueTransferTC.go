package refactored_closure

import (
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/client"
)

const NewValueTransferTCName = "newValueTransferTC"

var newValueTransferConfig *TCConfig

func InitNewValueTransferTC(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	newValueTransferConfig = Init(accs, contractsParam, endpoint, gp)
}

func RunValueTransfer() {
	RunBaseValueTransfer(newValueTransferConfig, func(cli *client.Client, from *account.Account, to *account.Account) (interface{}, *big.Int, error) {
		value := big.NewInt(int64(rand.Int() % 3))
		return from.TransferNewValueTransferTx(cli, to, value)
	}, "transferNewValueTransferTx")
}
