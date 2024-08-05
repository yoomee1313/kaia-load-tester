package task

import (
	"math/big"

	"github.com/klaytn/klaytn-load-tester/klayslave/account"
	"github.com/klaytn/klaytn-load-tester/klayslave/blockbench/analyticTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/blockbench/doNothingTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/blockbench/ioHeavyTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/blockbench/smallBankTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/blockbench/ycsbTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/cpuHeavyTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/erc20TransferTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/erc721TransferTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/ethereumTxAccessListTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/ethereumTxDynamicFeeTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/ethereumTxLegacyTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/internalTxTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/largeMemoTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newAccountCreationTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newAccountUpdateTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newCancelTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newEthereumAccessListTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newEthereumDynamicFeeTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedAccountUpdateTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedAccountUpdateWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedCancelTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedCancelWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedSmartContractDeployTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedSmartContractDeployWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedSmartContractExecutionTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedSmartContractExecutionWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedValueTransferMemoTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedValueTransferMemoWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedValueTransferTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newFeeDelegatedValueTransferWithRatioTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newSmartContractDeployTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newSmartContractExecutionTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newValueTransferLargeMemoTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newValueTransferMemoTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newValueTransferSmallMemoTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newValueTransferTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/newValueTransferWithCancelTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/readApiCallContractTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/readApiCallTC"
	receiptCheckTc "github.com/klaytn/klaytn-load-tester/klayslave/receiptCheckTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/storageTrieWriteTC"
	"github.com/klaytn/klaytn-load-tester/klayslave/transferSignedTc"
	"github.com/klaytn/klaytn-load-tester/klayslave/transferSignedWithCheckTc"
	"github.com/klaytn/klaytn-load-tester/klayslave/transferUnsignedTc"
	"github.com/klaytn/klaytn-load-tester/klayslave/userStorageTC"
)

type ExtendedTask struct {
	Name   string
	Weight int
	Fn     func()
	Init   func(accs []*account.Account, endpoint string, gp *big.Int)
}
type ExtendedTaskSet []*ExtendedTask

// InitTCList initializes TCs and returns a slice of TCs.
var tcList = map[string]*ExtendedTask{
	"analyticTx": {
		Name:   "analyticTx",
		Weight: 10,
		Fn:     analyticTC.Run,
		Init:   analyticTC.Init,
	},
	"analyticQueryLargestAccBalTx": {
		Name:   "analyticQueryLargestAccBalTx",
		Weight: 10,
		Fn:     analyticTC.QueryLargestAccBal,
		Init:   analyticTC.Init,
	},
	"analyticQueryLargestTxValTx": {
		Name:   "analyticQueryLargestTxValTx",
		Weight: 10,
		Fn:     analyticTC.QueryLargestTxVal,
		Init:   analyticTC.Init,
	},
	"analyticQueryTotalTxValTx": {
		Name:   "analyticQueryTotalTxValTx",
		Weight: 10,
		Fn:     analyticTC.QueryTotalTxVal,
		Init:   analyticTC.Init,
	},
	"cpuHeavyTx": {
		Name:   "cpuHeavyTx",
		Weight: 10,
		Fn:     cpuHeavyTC.Run,
		Init:   cpuHeavyTC.Init,
		//AccGrp:  accGrpForSignedTx, //[nUserForSigned/2:],
		//EndPint: gEndpoint,
	},
	"doNothingTx": {
		Name:   "doNothingTx",
		Weight: 10,
		Fn:     doNothingTC.Run,
		Init:   doNothingTC.Init,
	},
	internalTxTC.Name: {
		Name:   internalTxTC.Name,
		Weight: 10,
		Fn:     internalTxTC.Run,
		Init:   internalTxTC.Init,
	},
	internalTxTC.NameMintNFT: &ExtendedTask{
		Name:   internalTxTC.NameMintNFT,
		Weight: 10,
		Fn:     internalTxTC.RunMintNFT,
		Init:   internalTxTC.Init,
	},
	"ioHeavyTx": {
		Name:   "ioHeavyTx",
		Weight: 10,
		Fn:     ioHeavyTC.Run,
		Init:   ioHeavyTC.Init,
	},
	"ioHeavyScanTx": {
		Name:   "ioHeavyScanTx",
		Weight: 10,
		Fn:     ioHeavyTC.Scan,
		Init:   ioHeavyTC.Init,
	},
	"ioHeavyWriteTx": {
		Name:   "ioHeavyWriteTx",
		Weight: 10,
		Fn:     ioHeavyTC.Write,
		Init:   ioHeavyTC.Init,
	},
	"largeMemoTC": {
		Name:   "largeMemoTC",
		Weight: 10,
		Fn:     largeMemoTC.Run,
		Init:   largeMemoTC.Init,
	},
	receiptCheckTc.Name: {
		Name:   receiptCheckTc.Name,
		Weight: 10,
		Fn:     receiptCheckTc.Run,
		Init:   receiptCheckTc.Init,
	},
	"smallBankTx": {
		Name:   "smallBankTx",
		Weight: 10,
		Fn:     smallBankTC.Run,
		Init:   smallBankTC.Init,
	},
	"smallBankAlmagateTx": {
		Name:   "smallBankAlmagateTx",
		Weight: 10,
		Fn:     smallBankTC.Almagate,
		Init:   smallBankTC.Init,
	},
	"smallBankGetBalanceTx": {
		Name:   "smallBankGetBalanceTx",
		Weight: 10,
		Fn:     smallBankTC.GetBalance,
		Init:   smallBankTC.Init,
	},
	"smallBankSendPaymentTx": {
		Name:   "smallBankSendPaymentTx",
		Weight: 10,
		Fn:     smallBankTC.SendPayment,
		Init:   smallBankTC.Init,
	},
	"smallBankUpdateBalanceTx": {
		Name:   "smallBankUpdateBalanceTx",
		Weight: 10,
		Fn:     smallBankTC.UpdateBalance,
		Init:   smallBankTC.Init,
	},
	"smallBankUpdateSavingTx": {
		Name:   "smallBankUpdateSavingTx",
		Weight: 10,
		Fn:     smallBankTC.UpdateSaving,
		Init:   smallBankTC.Init,
	},
	"smallBankWriteCheckTx": {
		Name:   "smallBankWriteCheckTx",
		Weight: 10,
		Fn:     smallBankTC.WriteCheck,
		Init:   smallBankTC.Init,
	},
	"transferSignedTx": {
		Name:   "transferSignedTx",
		Weight: 10,
		Fn:     transferSignedTc.Run,
		Init:   transferSignedTc.Init,
		//AccGrp:  accGrpForSignedTx, //[:nUserForSigned/2-1],
		//EndPint: gEndpoint,
	},
	"newValueTransferTC": {
		Name:   "newValueTransferTC",
		Weight: 10,
		Fn:     newValueTransferTC.Run,
		Init:   newValueTransferTC.Init,
	},
	"newValueTransferWithCancelTC": {
		Name:   "newValueTransferWithCancelTC",
		Weight: 10,
		Fn:     newValueTransferWithCancelTC.Run,
		Init:   newValueTransferWithCancelTC.Init,
	},
	"newFeeDelegatedValueTransferTC": {
		Name:   "newFeeDelegatedValueTransferTC",
		Weight: 10,
		Fn:     newFeeDelegatedValueTransferTC.Run,
		Init:   newFeeDelegatedValueTransferTC.Init,
	},
	"newFeeDelegatedValueTransferWithRatioTC": {
		Name:   "newFeeDelegatedValueTransferWithRatioTC",
		Weight: 10,
		Fn:     newFeeDelegatedValueTransferWithRatioTC.Run,
		Init:   newFeeDelegatedValueTransferWithRatioTC.Init,
	},
	"newValueTransferMemoTC": {
		Name:   "newValueTransferMemoTC",
		Weight: 10,
		Fn:     newValueTransferMemoTC.Run,
		Init:   newValueTransferMemoTC.Init,
	},
	"newValueTransferLargeMemoTC": {
		Name:   "newValueTransferLargeMemoTC",
		Weight: 10,
		Fn:     newValueTransferLargeMemoTC.Run,
		Init:   newValueTransferLargeMemoTC.Init,
	},
	"newValueTransferSmallMemoTC": {
		Name:   "newValueTransferSmallMemoTC",
		Weight: 10,
		Fn:     newValueTransferSmallMemoTC.Run,
		Init:   newValueTransferSmallMemoTC.Init,
	},
	"newFeeDelegatedValueTransferMemoTC": {
		Name:   "newFeeDelegatedValueTransferMemoTC",
		Weight: 10,
		Fn:     newFeeDelegatedValueTransferMemoTC.Run,
		Init:   newFeeDelegatedValueTransferMemoTC.Init,
	},
	"newFeeDelegatedValueTransferMemoWithRatioTC": {
		Name:   "newFeeDelegatedValueTransferMemoWithRatioTC",
		Weight: 10,
		Fn:     newFeeDelegatedValueTransferMemoWithRatioTC.Run,
		Init:   newFeeDelegatedValueTransferMemoWithRatioTC.Init,
	},
	"newAccountCreationTC": {
		Name:   "newAccountCreationTC",
		Weight: 10,
		Fn:     newAccountCreationTC.Run,
		Init:   newAccountCreationTC.Init,
	},
	"newAccountUpdateTC": {
		Name:   "newAccountUpdateTC",
		Weight: 10,
		Fn:     newAccountUpdateTC.Run,
		Init:   newAccountUpdateTC.Init,
	},
	"newFeeDelegatedAccountUpdateTC": {
		Name:   "newFeeDelegatedAccountUpdateTC",
		Weight: 10,
		Fn:     newFeeDelegatedAccountUpdateTC.Run,
		Init:   newFeeDelegatedAccountUpdateTC.Init,
	},
	"newFeeDelegatedAccountUpdateWithRatioTC": {
		Name:   "newFeeDelegatedAccountUpdateWithRatioTC",
		Weight: 10,
		Fn:     newFeeDelegatedAccountUpdateWithRatioTC.Run,
		Init:   newFeeDelegatedAccountUpdateWithRatioTC.Init,
	},
	"newSmartContractDeployTC": {
		Name:   "newSmartContractDeployTC",
		Weight: 10,
		Fn:     newSmartContractDeployTC.Run,
		Init:   newSmartContractDeployTC.Init,
	},
	"newFeeDelegatedSmartContractDeployTC": {
		Name:   "newFeeDelegatedSmartContractDeployTC",
		Weight: 10,
		Fn:     newFeeDelegatedSmartContractDeployTC.Run,
		Init:   newFeeDelegatedSmartContractDeployTC.Init,
	},
	"newFeeDelegatedSmartContractDeployWithRatioTC": {
		Name:   "newFeeDelegatedSmartContractDeployWithRatioTC",
		Weight: 10,
		Fn:     newFeeDelegatedSmartContractDeployWithRatioTC.Run,
		Init:   newFeeDelegatedSmartContractDeployWithRatioTC.Init,
	},
	"newSmartContractExecutionTC": {
		Name:   "newSmartContractExecutionTC",
		Weight: 10,
		Fn:     newSmartContractExecutionTC.Run,
		Init:   newSmartContractExecutionTC.Init,
	},
	storageTrieWriteTC.Name: {
		Name:   storageTrieWriteTC.Name,
		Weight: 10,
		Fn:     storageTrieWriteTC.Run,
		Init:   storageTrieWriteTC.Init,
	},
	"newFeeDelegatedSmartContractExecutionTC": {
		Name:   "newFeeDelegatedSmartContractExecutionTC",
		Weight: 10,
		Fn:     newFeeDelegatedSmartContractExecutionTC.Run,
		Init:   newFeeDelegatedSmartContractExecutionTC.Init,
	},
	"newFeeDelegatedSmartContractExecutionWithRatioTC": {
		Name:   "newFeeDelegatedSmartContractExecutionWithRatioTC",
		Weight: 10,
		Fn:     newFeeDelegatedSmartContractExecutionWithRatioTC.Run,
		Init:   newFeeDelegatedSmartContractExecutionWithRatioTC.Init,
	},
	"newCancelTC": {
		Name:   "newCancelTC",
		Weight: 10,
		Fn:     newCancelTC.Run,
		Init:   newCancelTC.Init,
	},
	"newFeeDelegatedCancelTC": {
		Name:   "newFeeDelegatedCancelTC",
		Weight: 10,
		Fn:     newFeeDelegatedCancelTC.Run,
		Init:   newFeeDelegatedCancelTC.Init,
	},
	"newFeeDelegatedCancelWithRatioTC": {
		Name:   "newFeeDelegatedCancelWithRatioTC",
		Weight: 10,
		Fn:     newFeeDelegatedCancelWithRatioTC.Run,
		Init:   newFeeDelegatedCancelWithRatioTC.Init,
	},
	"transferSignedWithCheckTx": {
		Name:   "transferSignedWithCheckTx",
		Weight: 10,
		Fn:     transferSignedWithCheckTc.Run,
		Init:   transferSignedWithCheckTc.Init,
	},
	"transferUnsignedTx": {
		Name:   "transferUnsignedTx",
		Weight: 10,
		Fn:     transferUnsignedTc.Run,
		Init:   transferUnsignedTc.Init,
	},
	"userStorageSetTx": {
		Name:   "userStorageSetTx",
		Weight: 10,
		Fn:     userStorageTC.RunSet,
		Init:   userStorageTC.Init,
	},
	"userStorageSetGetTx": {
		Name:   "userStorageSetGetTx",
		Weight: 10,
		Fn:     userStorageTC.RunSetGet,
		Init:   userStorageTC.Init,
	},
	"ycsbTx": {
		Name:   "ycsbTx",
		Weight: 10,
		Fn:     ycsbTC.Run,
		Init:   ycsbTC.Init,
	},
	"ycsbGetTx": {
		Name:   "ycsbGetTx",
		Weight: 10,
		Fn:     ycsbTC.Get,
		Init:   ycsbTC.Init,
	},
	"ycsbSetTx": {
		Name:   "ycsbSetTx",
		Weight: 10,
		Fn:     ycsbTC.Set,
		Init:   ycsbTC.Init,
	},
	erc20TransferTC.Name: {
		Name:   erc20TransferTC.Name,
		Weight: 10,
		Fn:     erc20TransferTC.Run,
		Init:   erc20TransferTC.Init,
	},
	erc721TransferTC.Name: {
		Name:   erc721TransferTC.Name,
		Weight: 10,
		Fn:     erc721TransferTC.Run,
		Init:   erc721TransferTC.Init,
	},
	"readGasPrice": {
		Name:   "readGasPrice",
		Weight: 10,
		Fn:     readApiCallTC.GasPrice,
		Init:   readApiCallTC.Init,
	},
	"readBlockNumber": {
		Name:   "readBlockNumber",
		Weight: 10,
		Fn:     readApiCallTC.BlockNumber,
		Init:   readApiCallTC.Init,
	},
	"readGetBlockByNumber": {
		Name:   "readGetBlockByNumber",
		Weight: 10,
		Fn:     readApiCallTC.GetBlockByNumber,
		Init:   readApiCallTC.Init,
	},
	"readGetAccount": {
		Name:   "readGetAccount",
		Weight: 10,
		Fn:     readApiCallTC.GetAccount,
		Init:   readApiCallTC.Init,
	},
	"readGetBlockWithConsensusInfoByNumber": {
		Name:   "readGetBlockWithConsensusInfoByNumber",
		Weight: 10,
		Fn:     readApiCallTC.GetBlockWithConsensusInfoByNumber,
		Init:   readApiCallTC.Init,
	},
	"readGetStorageAt": {
		Name:   "readGetStorageAt",
		Weight: 10,
		Fn:     readApiCallContractTC.GetStorageAt,
		Init:   readApiCallContractTC.Init,
	},
	"readCall": {
		Name:   "readCall",
		Weight: 10,
		Fn:     readApiCallContractTC.Call,
		Init:   readApiCallContractTC.Init,
	},
	"readEstimateGas": {
		Name:   "readEstimateGas",
		Weight: 10,
		Fn:     readApiCallContractTC.EstimateGas,
		Init:   readApiCallContractTC.Init,
	},
	"ethereumTxLegacyTC": {
		Name:   "ethereumTxLegacyTC",
		Weight: 10,
		Fn:     ethereumTxLegacyTC.Run,
		Init:   ethereumTxLegacyTC.Init,
	},
	"ethereumTxAccessListTC": {
		Name:   "ethereumTxAccessListTC",
		Weight: 10,
		Fn:     ethereumTxAccessListTC.Run,
		Init:   ethereumTxAccessListTC.Init,
	},
	"ethereumTxDynamicFeeTC": {
		Name:   "ethereumTxDynamicFeeTC",
		Weight: 10,
		Fn:     ethereumTxDynamicFeeTC.Run,
		Init:   ethereumTxDynamicFeeTC.Init,
	},
	"newEthereumAccessListTC": {
		Name:   "newEthereumAccessListTC",
		Weight: 10,
		Fn:     newEthereumAccessListTC.Run,
		Init:   newEthereumAccessListTC.Init,
	},
	"newEthereumDynamicFeeTC": {
		Name:   "newEthereumDynamicFeeTC",
		Weight: 10,
		Fn:     newEthereumDynamicFeeTC.Run,
		Init:   newEthereumDynamicFeeTC.Init,
	},
}
