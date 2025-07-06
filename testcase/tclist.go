package testcase

import (
	"math/big"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/testcase/tcutil"
)

type ExtendedTask struct {
	Name   string
	Weight int
	Init   func(accs []*account.Account, endpoint string, gp *big.Int) *tcutil.TcConfig
	Run    func(config *tcutil.TcConfig)
	Config *tcutil.TcConfig
}

// TcList holds all available test operations in EVM-style
var TcList = map[string]*ExtendedTask{
	"analyticQueryLargestAccBalTx": {
		Name:   "analyticQueryLargestAccBalTx",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    queryLargestAccBalRun,
	},
	"analyticQueryLargestTxValTx": {
		Name:   "analyticQueryLargestTxValTx",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    queryLargestTxValRun,
	},
	"analyticQueryTotalTxValTx": {
		Name:   "analyticQueryTotalTxValTx",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    queryTotalTxValRun,
	},
	"cpuHeavyTx": {
		Name:   "cpuHeavyTx",
		Weight: 10,
		Init:   cpuHeavyInit,
		Run:    cpuHeavyRun,
	},
	"doNothingTx": {
		Name:   "doNothingTx",
		Weight: 10,
		Init:   doNothingInit,
		Run:    doNothingRun,
	},
	"internalTxTC": {
		Name:   "internalTxTC",
		Weight: 10,
		Init:   internalTxInit,
		Run:    internalTxRun,
	},
	"internalTxMintNFTTC": {
		Name:   "internalTxMintNFTTC",
		Weight: 10,
		Init:   internalTxInit,
		Run:    internalTxMintNFTRun,
	},
	"ioHeavyTx": {
		Name:   "ioHeavyTx",
		Weight: 10,
		Init:   ioHeavyInit,
		Run:    ioHeavyRun,
	},
	"ioHeavyScanTx": {
		Name:   "ioHeavyScanTx",
		Weight: 10,
		Init:   ioHeavyInit,
		Run:    ioHeavyScanRun,
	},
	"ioHeavyWriteTx": {
		Name:   "ioHeavyWriteTx",
		Weight: 10,
		Init:   ioHeavyInit,
		Run:    ioHeavyWriteRun,
	},
	"largeMemoTC": {
		Name:   "largeMemoTC",
		Weight: 10,
		Init:   largeMemoInit,
		Run:    largeMemoRun,
	},
	"receiptCheckTC": {
		Name:   "receiptCheckTC",
		Weight: 10,
		Init:   receiptCheckInit,
		Run:    receiptCheckRun,
	},
	"smallBankTx": {
		Name:   "smallBankTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankRun,
	},
	"smallBankAlmagateTx": {
		Name:   "smallBankAlmagateTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankAlmagateRun,
	},
	"smallBankGetBalanceTx": {
		Name:   "smallBankGetBalanceTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankGetBalanceRun,
	},
	"smallBankSendPaymentTx": {
		Name:   "smallBankSendPaymentTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankUpdateBalanceRun,
	},
	"smallBankUpdateBalanceTx": {
		Name:   "smallBankUpdateBalanceTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankUpdateSavingRun,
	},
	"smallBankUpdateSavingTx": {
		Name:   "smallBankUpdateSavingTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankSendPaymentRun,
	},
	"smallBankWriteCheckTx": {
		Name:   "smallBankWriteCheckTx",
		Weight: 10,
		Init:   smallBankInit,
		Run:    smallBankWriteCheckRun,
	},
	"transferSignedTx": {
		Name:   "transferSignedTx",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    transferSignedRun,
	},
	"newValueTransferTC": {
		Name:   "newValueTransferTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    valueTransferRun,
	},
	"newValueTransferWithCancelTC": {
		Name:   "newValueTransferWithCancelTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newValueTransferWithCancelRun,
	},
	"newFeeDelegatedValueTransferTC": {
		Name:   "newFeeDelegatedValueTransferTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedValueTransferRun,
	},
	"newFeeDelegatedValueTransferWithRatioTC": {
		Name:   "newFeeDelegatedValueTransferWithRatioTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedValueTransferWithRatioRun,
	},
	"newValueTransferMemoTC": {
		Name:   "newValueTransferMemoTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newValueTransferMemoRun,
	},
	"newValueTransferLargeMemoTC": {
		Name:   "newValueTransferLargeMemoTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newValueTransferLargeMemoRun,
	},
	"newValueTransferSmallMemoTC": {
		Name:   "newValueTransferSmallMemoTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newValueTransferSmallMemoRun,
	},
	"newFeeDelegatedValueTransferMemoTC": {
		Name:   "newFeeDelegatedValueTransferMemoTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedValueTransferMemoRun,
	},
	"newFeeDelegatedValueTransferMemoWithRatioTC": {
		Name:   "newFeeDelegatedValueTransferMemoWithRatioTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedValueTransferMemoWithRatioRun,
	},
	"newAccountCreationTC": {
		Name:   "newAccountCreationTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newAccountCreationRun,
	},
	"newAccountUpdateTC": {
		Name:   "newAccountUpdateTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newAccountUpdateRun,
	},
	"newFeeDelegatedAccountUpdateTC": {
		Name:   "newFeeDelegatedAccountUpdateTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedAccountUpdateRun,
	},
	"newFeeDelegatedAccountUpdateWithRatioTC": {
		Name:   "newFeeDelegatedAccountUpdateWithRatioTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedAccountUpdateWithRatioRun,
	},
	"newSmartContractDeployTC": {
		Name:   "newSmartContractDeployTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newSmartContractDeployRun,
	},
	"newFeeDelegatedSmartContractDeployTC": {
		Name:   "newFeeDelegatedSmartContractDeployTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedSmartContractDeployRun,
	},
	"newFeeDelegatedSmartContractDeployWithRatioTC": {
		Name:   "newFeeDelegatedSmartContractDeployWithRatioTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedSmartContractDeployWithRatioRun,
	},
	"newSmartContractExecutionTC": {
		Name:   "newSmartContractExecutionTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newSmartContractExecutionRun,
	},
	"storageTrieWriteTC": {
		Name:   "storageTrieWriteTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    storageTrieWriteRun,
	},
	"newFeeDelegatedSmartContractExecutionTC": {
		Name:   "newFeeDelegatedSmartContractExecutionTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedSmartContractExecutionRun,
	},
	"newFeeDelegatedSmartContractExecutionWithRatioTC": {
		Name:   "newFeeDelegatedSmartContractExecutionWithRatioTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedSmartContractExecutionWithRatioRun,
	},
	"newCancelTC": {
		Name:   "newCancelTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newCancelRun,
	},
	"newFeeDelegatedCancelTC": {
		Name:   "newFeeDelegatedCancelTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedCancelRun,
	},
	"newFeeDelegatedCancelWithRatioTC": {
		Name:   "newFeeDelegatedCancelWithRatioTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    newFeeDelegatedCancelWithRatioRun,
	},
	"transferSignedWithCheckTx": {
		Name:   "transferSignedWithCheckTx",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    transferSignedWithCheckRun,
	},
	"transferUnsignedTx": {
		Name:   "transferUnsignedTx",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    transferUnsignedRun,
	},
	"userStorageSetTx": {
		Name:   "userStorageSetTx",
		Weight: 10,
		Init:   storageTrieWriteInit,
		Run:    storageTrieWriteSetRun,
	},
	"userStorageSetGetTx": {
		Name:   "userStorageSetGetTx",
		Weight: 10,
		Init:   storageTrieWriteInit,
		Run:    storageTrieWriteSetGetRun,
	},
	"ycsbTx": {
		Name:   "ycsbTx",
		Weight: 10,
		Init:   ycsbInit,
		Run:    ycsbRun,
	},
	"ycsbGetTx": {
		Name:   "ycsbGetTx",
		Weight: 10,
		Init:   ycsbInit,
		Run:    ycsbSetRun,
	},
	"ycsbSetTx": {
		Name:   "ycsbSetTx",
		Weight: 10,
		Init:   ycsbInit,
		Run:    ycsbGetRun,
	},
	"erc20TransferTC": {
		Name:   "erc20TransferTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    erc20TransferRun,
	},
	"erc721TransferTC": {
		Name:   "erc721TransferTC",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    erc721TransferRun,
	},
	"readGasPrice": {
		Name:   "readGasPrice",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    readApiCallGasPriceRun,
	},
	"readBlockNumber": {
		Name:   "readBlockNumber",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    readApiCallBlockNumberRun,
	},
	"readGetBlockByNumber": {
		Name:   "readGetBlockByNumber",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    readApiCallGetBlockByNumberRun,
	},
	"readGetAccount": {
		Name:   "readGetAccount",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    readApiCallGetAccountRun,
	},
	"readGetBlockWithConsensusInfoByNumber": {
		Name:   "readGetBlockWithConsensusInfoByNumber",
		Weight: 10,
		Init:   tcutil.InitTcConfig,
		Run:    readApiCallGetBlockWithConsensusInfoByNumberRun,
	},
	"readGetStorageAt": {
		Name:   "readGetStorageAt",
		Weight: 10,
		Init:   readApiCallContractInit,
		Run:    readApiCallContractGetStorageAtRun,
	},
	"readCall": {
		Name:   "readCall",
		Weight: 10,
		Init:   readApiCallContractInit,
		Run:    readApiCallContractCallRun,
	},
	"readEstimateGas": {
		Name:   "readEstimateGas",
		Weight: 10,
		Init:   readApiCallContractInit,
		Run:    readApiCallContractEstimateGasRun,
	},
	"ethereumTxLegacyTC": {
		Name:   "ethereumTxLegacyTC",
		Weight: 10,
		Init:   ethereumTxInit,
		Run:    ethereumTxLegacyRun,
	},
	"ethereumTxAccessListTC": {
		Name:   "ethereumTxAccessListTC",
		Weight: 10,
		Init:   ethereumTxInit,
		Run:    ethereumTxAccessListRun,
	},
	"ethereumTxDynamicFeeTC": {
		Name:   "ethereumTxDynamicFeeTC",
		Weight: 10,
		Init:   ethereumTxInit,
		Run:    ethereumTxDynamicFeeRun,
	},
	"newEthereumAccessListTC": {
		Name:   "newEthereumAccessListTC",
		Weight: 10,
		Init:   ethereumTxInit,
		Run:    newEthereumAccessListRun,
	},
	"newEthereumDynamicFeeTC": {
		Name:   "newEthereumDynamicFeeTC",
		Weight: 10,
		Init:   ethereumTxInit,
		Run:    newEthereumDynamicFeeRun,
	},
}
