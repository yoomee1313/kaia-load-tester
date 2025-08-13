package refactored_closure

import (
	"math/big"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/client"
)

const NewSmartContractExecutionTCName = "newSmartContractExecutionTC"

var newSmartContractExecutionConfig *TCConfig

func InitNewSmartContractExecutionTC(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	newSmartContractExecutionConfig = Init(accs, contractsParam, endpoint, gp)
	newSmartContractExecutionConfig.SmartContractAccount = contractsParam[account.ContractGeneral]
}

func RunSmartContractExecution() {
	RunBaseWithContract(newSmartContractExecutionConfig, func(cli *client.Client, from *account.Account, to *account.Account) (interface{}, *big.Int, error) {
		return from.TransferNewSmartContractExecutionTx(cli, to, nil, account.TestContractInfos[account.ContractGeneral].GenData(from.GetAddress(), nil))
	}, "transferNewSmartContractExecutionTx")
}
