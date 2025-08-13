package refactored_closure

import (
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia/client"
)

const CpuHeavyTCName = "cpuHeavyTC"

var cpuHeavyConfig *TCConfig

func InitCpuHeavyTC(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	cpuHeavyConfig = Init(accs, contractsParam, endpoint, gp)
	cpuHeavyConfig.SmartContractAccount = contractsParam[account.ContractCPUHeavy]
}

func RunCpuHeavy() {
	RunBaseWithContract(cpuHeavyConfig, func(cli *client.Client, from *account.Account, to *account.Account) (interface{}, *big.Int, error) {
		value := big.NewInt(int64(rand.Int() % 3))
		return from.TransferNewSmartContractExecutionTx(cli, to, value, account.TestContractInfos[account.ContractCPUHeavy].GenData(from.GetAddress(), nil))
	}, "transferCpuHeavyTx")
}
