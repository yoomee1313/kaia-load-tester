package task

import (
	"github.com/klaytn/klaytn-load-tester/klayslave/account"
	"github.com/klaytn/klaytn-load-tester/klayslave/erc20TransferTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/erc721TransferTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/ethereumTxAccessListTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/ethereumTxDynamicFeeTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/ethereumTxLegacyTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newEthereumAccessListTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newEthereumDynamicFeeTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedSmartContractExecutionTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedSmartContractExecutionWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newSmartContractExecutionTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/storageTrieWriteTC"
)

func SetTcGeneralSmartContract(generalSmartContract *account.Account) {
	newSmartContractExecutionTC.SmartContractAccount = generalSmartContract
	newFeeDelegatedSmartContractExecutionTC.SmartContractAccount = generalSmartContract
	newFeeDelegatedSmartContractExecutionWithRatioTC.SmartContractAccount = generalSmartContract
	ethereumTxLegacyTC.SmartContractAccount = generalSmartContract
	ethereumTxAccessListTC.SmartContractAccount = generalSmartContract
	ethereumTxDynamicFeeTC.SmartContractAccount = generalSmartContract
	newEthereumAccessListTC.SmartContractAccount = generalSmartContract
	newEthereumDynamicFeeTC.SmartContractAccount = generalSmartContract
}

func SetErc20TransferTcContract(sca *account.Account) {
	erc20TransferTC.SmartContractAccount = sca
}

func SetErc721TransferTcContract(sca *account.Account) {
	erc721TransferTC.SmartContractAccount = sca
}

func SetStorageTrieWriteTcContract(sca *account.Account) {
	storageTrieWriteTC.SmartContractAccount = sca
}
