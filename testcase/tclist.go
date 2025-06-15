package testcase

import (
	"math/big"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/testcase/blockbench/analyticTC"
	"github.com/kaiachain/kaia-load-tester/testcase/blockbench/doNothingTC"
	"github.com/kaiachain/kaia-load-tester/testcase/blockbench/ioHeavyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/blockbench/smallBankTC"
	"github.com/kaiachain/kaia-load-tester/testcase/blockbench/ycsbTC"
	"github.com/kaiachain/kaia-load-tester/testcase/cpuHeavyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/erc20TransferTC"
	"github.com/kaiachain/kaia-load-tester/testcase/erc721TransferTC"
	"github.com/kaiachain/kaia-load-tester/testcase/ethereumTxAccessListTC"
	"github.com/kaiachain/kaia-load-tester/testcase/ethereumTxDynamicFeeTC"
	"github.com/kaiachain/kaia-load-tester/testcase/ethereumTxLegacyTC"
	"github.com/kaiachain/kaia-load-tester/testcase/gaslessRevertTransactionTC"
	"github.com/kaiachain/kaia-load-tester/testcase/gaslessTransactionTC"
	"github.com/kaiachain/kaia-load-tester/testcase/internalTxTC"
	"github.com/kaiachain/kaia-load-tester/testcase/largeMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newAccountCreationTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newAccountUpdateTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newCancelTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newEthereumAccessListTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newEthereumDynamicFeeTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedAccountUpdateTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedAccountUpdateWithRatioTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedCancelTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedCancelWithRatioTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedSmartContractDeployTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedSmartContractDeployWithRatioTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedSmartContractExecutionTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedSmartContractExecutionWithRatioTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedValueTransferMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedValueTransferMemoWithRatioTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedValueTransferTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newFeeDelegatedValueTransferWithRatioTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newSmartContractDeployTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newSmartContractExecutionTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newValueTransferLargeMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newValueTransferMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newValueTransferSmallMemoTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newValueTransferTC"
	"github.com/kaiachain/kaia-load-tester/testcase/newValueTransferWithCancelTC"
	"github.com/kaiachain/kaia-load-tester/testcase/readApiCallContractTC"
	"github.com/kaiachain/kaia-load-tester/testcase/readApiCallTC"
	receiptCheckTc "github.com/kaiachain/kaia-load-tester/testcase/receiptCheckTC"
	"github.com/kaiachain/kaia-load-tester/testcase/storageTrieWriteTC"
	"github.com/kaiachain/kaia-load-tester/testcase/transferSignedTc"
	"github.com/kaiachain/kaia-load-tester/testcase/transferSignedWithCheckTc"
	"github.com/kaiachain/kaia-load-tester/testcase/transferUnsignedTc"
	"github.com/kaiachain/kaia-load-tester/testcase/userStorageTC"
)

type ExtendedTask struct {
	Name   string
	Weight int
	Fn     func()
	Init   func(accs []*account.Account, endpoint string, gp *big.Int)
}
type ExtendedTaskSet []*ExtendedTask

// TcList initializes TCs and returns a slice of TCs.
var TcList = map[string]*ExtendedTask{
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
	"gaslessTransactionTC": {
		Name:   "gaslessTransactionTC",
		Weight: 10,
		Fn:     gaslessTransactionTC.Run,
		Init:   gaslessTransactionTC.Init,
	},
	"gaslessRevertTransactionTC": {
		Name:   "gaslessRevertTransactionTC",
		Weight: 10,
		Fn:     gaslessRevertTransactionTC.Run,
		Init:   gaslessRevertTransactionTC.Init,
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
