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
	Name    string
	Weight  int
	Fn      func()
	Init    func(accs []*account.Account, endpoint string, gp *big.Int)
	AccGrp  []*account.Account
	EndPint string
}

// InitTCList initializes TCs and returns a slice of TCs.
func InitTCList(accGrpForSignedTx []*account.Account, accGrpForUnsignedTx []*account.Account, gEndpoint string) (taskSet []*ExtendedTask) {
	taskSet = append(taskSet, &ExtendedTask{
		Name:    "analyticTx",
		Weight:  10,
		Fn:      analyticTC.Run,
		Init:    analyticTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "analyticQueryLargestAccBalTx",
		Weight:  10,
		Fn:      analyticTC.QueryLargestAccBal,
		Init:    analyticTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "analyticQueryLargestTxValTx",
		Weight:  10,
		Fn:      analyticTC.QueryLargestTxVal,
		Init:    analyticTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "analyticQueryTotalTxValTx",
		Weight:  10,
		Fn:      analyticTC.QueryTotalTxVal,
		Init:    analyticTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "cpuHeavyTx",
		Weight:  10,
		Fn:      cpuHeavyTC.Run,
		Init:    cpuHeavyTC.Init,
		AccGrp:  accGrpForSignedTx, //[nUserForSigned/2:],
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "doNothingTx",
		Weight:  10,
		Fn:      doNothingTC.Run,
		Init:    doNothingTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    internalTxTC.Name,
		Weight:  10,
		Fn:      internalTxTC.Run,
		Init:    internalTxTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    internalTxTC.NameMintNFT,
		Weight:  10,
		Fn:      internalTxTC.RunMintNFT,
		Init:    internalTxTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ioHeavyTx",
		Weight:  10,
		Fn:      ioHeavyTC.Run,
		Init:    ioHeavyTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ioHeavyScanTx",
		Weight:  10,
		Fn:      ioHeavyTC.Scan,
		Init:    ioHeavyTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ioHeavyWriteTx",
		Weight:  10,
		Fn:      ioHeavyTC.Write,
		Init:    ioHeavyTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "largeMemoTC",
		Weight:  10,
		Fn:      largeMemoTC.Run,
		Init:    largeMemoTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    receiptCheckTc.Name,
		Weight:  10,
		Fn:      receiptCheckTc.Run,
		Init:    receiptCheckTc.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankTx",
		Weight:  10,
		Fn:      smallBankTC.Run,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankAlmagateTx",
		Weight:  10,
		Fn:      smallBankTC.Almagate,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankGetBalanceTx",
		Weight:  10,
		Fn:      smallBankTC.GetBalance,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankSendPaymentTx",
		Weight:  10,
		Fn:      smallBankTC.SendPayment,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankUpdateBalanceTx",
		Weight:  10,
		Fn:      smallBankTC.UpdateBalance,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankUpdateSavingTx",
		Weight:  10,
		Fn:      smallBankTC.UpdateSaving,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "smallBankWriteCheckTx",
		Weight:  10,
		Fn:      smallBankTC.WriteCheck,
		Init:    smallBankTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "transferSignedTx",
		Weight:  10,
		Fn:      transferSignedTc.Run,
		Init:    transferSignedTc.Init,
		AccGrp:  accGrpForSignedTx, //[:nUserForSigned/2-1],
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newValueTransferTC",
		Weight:  10,
		Fn:      newValueTransferTC.Run,
		Init:    newValueTransferTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newValueTransferWithCancelTC",
		Weight:  10,
		Fn:      newValueTransferWithCancelTC.Run,
		Init:    newValueTransferWithCancelTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedValueTransferTC",
		Weight:  10,
		Fn:      newFeeDelegatedValueTransferTC.Run,
		Init:    newFeeDelegatedValueTransferTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedValueTransferWithRatioTC",
		Weight:  10,
		Fn:      newFeeDelegatedValueTransferWithRatioTC.Run,
		Init:    newFeeDelegatedValueTransferWithRatioTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newValueTransferMemoTC",
		Weight:  10,
		Fn:      newValueTransferMemoTC.Run,
		Init:    newValueTransferMemoTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newValueTransferLargeMemoTC",
		Weight:  10,
		Fn:      newValueTransferLargeMemoTC.Run,
		Init:    newValueTransferLargeMemoTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newValueTransferSmallMemoTC",
		Weight:  10,
		Fn:      newValueTransferSmallMemoTC.Run,
		Init:    newValueTransferSmallMemoTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedValueTransferMemoTC",
		Weight:  10,
		Fn:      newFeeDelegatedValueTransferMemoTC.Run,
		Init:    newFeeDelegatedValueTransferMemoTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedValueTransferMemoWithRatioTC",
		Weight:  10,
		Fn:      newFeeDelegatedValueTransferMemoWithRatioTC.Run,
		Init:    newFeeDelegatedValueTransferMemoWithRatioTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newAccountCreationTC",
		Weight:  10,
		Fn:      newAccountCreationTC.Run,
		Init:    newAccountCreationTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newAccountUpdateTC",
		Weight:  10,
		Fn:      newAccountUpdateTC.Run,
		Init:    newAccountUpdateTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedAccountUpdateTC",
		Weight:  10,
		Fn:      newFeeDelegatedAccountUpdateTC.Run,
		Init:    newFeeDelegatedAccountUpdateTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedAccountUpdateWithRatioTC",
		Weight:  10,
		Fn:      newFeeDelegatedAccountUpdateWithRatioTC.Run,
		Init:    newFeeDelegatedAccountUpdateWithRatioTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newSmartContractDeployTC",
		Weight:  10,
		Fn:      newSmartContractDeployTC.Run,
		Init:    newSmartContractDeployTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedSmartContractDeployTC",
		Weight:  10,
		Fn:      newFeeDelegatedSmartContractDeployTC.Run,
		Init:    newFeeDelegatedSmartContractDeployTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedSmartContractDeployWithRatioTC",
		Weight:  10,
		Fn:      newFeeDelegatedSmartContractDeployWithRatioTC.Run,
		Init:    newFeeDelegatedSmartContractDeployWithRatioTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newSmartContractExecutionTC",
		Weight:  10,
		Fn:      newSmartContractExecutionTC.Run,
		Init:    newSmartContractExecutionTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    storageTrieWriteTC.Name,
		Weight:  10,
		Fn:      storageTrieWriteTC.Run,
		Init:    storageTrieWriteTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedSmartContractExecutionTC",
		Weight:  10,
		Fn:      newFeeDelegatedSmartContractExecutionTC.Run,
		Init:    newFeeDelegatedSmartContractExecutionTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedSmartContractExecutionWithRatioTC",
		Weight:  10,
		Fn:      newFeeDelegatedSmartContractExecutionWithRatioTC.Run,
		Init:    newFeeDelegatedSmartContractExecutionWithRatioTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newCancelTC",
		Weight:  10,
		Fn:      newCancelTC.Run,
		Init:    newCancelTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedCancelTC",
		Weight:  10,
		Fn:      newFeeDelegatedCancelTC.Run,
		Init:    newFeeDelegatedCancelTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newFeeDelegatedCancelWithRatioTC",
		Weight:  10,
		Fn:      newFeeDelegatedCancelWithRatioTC.Run,
		Init:    newFeeDelegatedCancelWithRatioTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "transferSignedWithCheckTx",
		Weight:  10,
		Fn:      transferSignedWithCheckTc.Run,
		Init:    transferSignedWithCheckTc.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "transferUnsignedTx",
		Weight:  10,
		Fn:      transferUnsignedTc.Run,
		Init:    transferUnsignedTc.Init,
		AccGrp:  accGrpForUnsignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "userStorageSetTx",
		Weight:  10,
		Fn:      userStorageTC.RunSet,
		Init:    userStorageTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "userStorageSetGetTx",
		Weight:  10,
		Fn:      userStorageTC.RunSetGet,
		Init:    userStorageTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ycsbTx",
		Weight:  10,
		Fn:      ycsbTC.Run,
		Init:    ycsbTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ycsbGetTx",
		Weight:  10,
		Fn:      ycsbTC.Get,
		Init:    ycsbTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ycsbSetTx",
		Weight:  10,
		Fn:      ycsbTC.Set,
		Init:    ycsbTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    erc20TransferTC.Name,
		Weight:  10,
		Fn:      erc20TransferTC.Run,
		Init:    erc20TransferTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    erc721TransferTC.Name,
		Weight:  10,
		Fn:      erc721TransferTC.Run,
		Init:    erc721TransferTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readGasPrice",
		Weight:  10,
		Fn:      readApiCallTC.GasPrice,
		Init:    readApiCallTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readBlockNumber",
		Weight:  10,
		Fn:      readApiCallTC.BlockNumber,
		Init:    readApiCallTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readGetBlockByNumber",
		Weight:  10,
		Fn:      readApiCallTC.GetBlockByNumber,
		Init:    readApiCallTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readGetAccount",
		Weight:  10,
		Fn:      readApiCallTC.GetAccount,
		Init:    readApiCallTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readGetBlockWithConsensusInfoByNumber",
		Weight:  10,
		Fn:      readApiCallTC.GetBlockWithConsensusInfoByNumber,
		Init:    readApiCallTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readGetStorageAt",
		Weight:  10,
		Fn:      readApiCallContractTC.GetStorageAt,
		Init:    readApiCallContractTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readCall",
		Weight:  10,
		Fn:      readApiCallContractTC.Call,
		Init:    readApiCallContractTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "readEstimateGas",
		Weight:  10,
		Fn:      readApiCallContractTC.EstimateGas,
		Init:    readApiCallContractTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ethereumTxLegacyTC",
		Weight:  10,
		Fn:      ethereumTxLegacyTC.Run,
		Init:    ethereumTxLegacyTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ethereumTxAccessListTC",
		Weight:  10,
		Fn:      ethereumTxAccessListTC.Run,
		Init:    ethereumTxAccessListTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "ethereumTxDynamicFeeTC",
		Weight:  10,
		Fn:      ethereumTxDynamicFeeTC.Run,
		Init:    ethereumTxDynamicFeeTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newEthereumAccessListTC",
		Weight:  10,
		Fn:      newEthereumAccessListTC.Run,
		Init:    newEthereumAccessListTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	taskSet = append(taskSet, &ExtendedTask{
		Name:    "newEthereumDynamicFeeTC",
		Weight:  10,
		Fn:      newEthereumDynamicFeeTC.Run,
		Init:    newEthereumDynamicFeeTC.Init,
		AccGrp:  accGrpForSignedTx,
		EndPint: gEndpoint,
	})

	return taskSet
}
