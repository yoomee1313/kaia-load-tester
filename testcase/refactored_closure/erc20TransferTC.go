package refactored_closure

import (
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/client"
)

const Erc20TransferTCName = "erc20TransferTC"

var erc20TransferConfig *TCConfig

func InitErc20TransferTC(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	erc20TransferConfig = Init(accs, contractsParam, endpoint, gp)
	erc20TransferConfig.SmartContractAccount = contractsParam[account.ContractErc20]
}

func RunErc20Transfer() {
	RunBaseValueTransfer(erc20TransferConfig, func(cli *client.Client, from *account.Account, to *account.Account) (interface{}, *big.Int, error) {
		value := big.NewInt(int64(rand.Int() % 3))
		data := account.TestContractInfos[account.ContractErc20].GenData(to.GetAddress(), value)
		return from.TransferNewSmartContractExecutionTx(cli, erc20TransferConfig.SmartContractAccount, nil, data)
	}, "transferErc20TransferTx")
}
