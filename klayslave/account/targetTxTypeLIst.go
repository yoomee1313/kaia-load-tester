package account

import (
	"math/big"

	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/common"
)

type TargetTxType struct {
	Description string
	GenerateTx  func(account *Account, nonce uint64, suggestedGasPrice *big.Int) *types.Transaction
}

// TargetTxTypeList defines the list of auction target tx types.
var TargetTxTypeList = map[string]*TargetTxType{
	"VT": {
		Description: "VT creates a ValueTransfer that sends a 1kei to itself.",
		GenerateTx: func(account *Account, nonce uint64, suggestedGasPrice *big.Int) *types.Transaction {
			signer := types.NewEIP155Signer(chainID)
			tx, err := types.NewTransactionWithMap(types.TxTypeValueTransfer, map[types.TxValueKeyType]interface{}{
				types.TxValueKeyNonce:    nonce,
				types.TxValueKeyTo:       account.address,
				types.TxValueKeyAmount:   common.Big1,
				types.TxValueKeyGasLimit: uint64(100000),
				types.TxValueKeyGasPrice: suggestedGasPrice,
				types.TxValueKeyFrom:     account.address,
			})
			if err != nil {
				return nil
			}

			err = tx.SignWithKeys(signer, account.privateKey)
			if err != nil {
				return nil
			}
			return tx
		},
	},
}
