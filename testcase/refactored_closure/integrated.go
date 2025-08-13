package refactored_closure

import (
	"math/big"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
)

// InitAllTCs initializes all refactored test cases
func InitAllTCs(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	// Value transfer TCs
	InitNewValueTransferTC(accs, contractsParam, endpoint, gp)
	InitNewValueTransferMemoTC(accs, contractsParam, endpoint, gp)

	// Smart contract execution TCs
	InitNewSmartContractExecutionTC(accs, contractsParam, endpoint, gp)
	InitErc20TransferTC(accs, contractsParam, endpoint, gp)
	InitCpuHeavyTC(accs, contractsParam, endpoint, gp)
}

// RunTCByName runs a specific test case by name
func RunTCByName(tcName string) {
	switch tcName {
	case NewValueTransferTCName:
		RunValueTransfer()
	case NewValueTransferMemoTCName:
		RunValueTransferMemo()
	case NewSmartContractExecutionTCName:
		RunSmartContractExecution()
	case Erc20TransferTCName:
		RunErc20Transfer()
	case CpuHeavyTCName:
		RunCpuHeavy()
	default:
		// Unknown TC name, do nothing
	}
}

// GetAvailableTCNames returns a list of all available test case names
func GetAvailableTCNames() []string {
	return []string{
		NewValueTransferTCName,
		NewValueTransferMemoTCName,
		NewSmartContractExecutionTCName,
		Erc20TransferTCName,
		CpuHeavyTCName,
	}
}
