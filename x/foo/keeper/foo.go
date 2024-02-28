package keeper

import (
	lightclienttypes "foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Foo uses lightclient module to verify transactions and does something with it
func (k Keeper) Foo(ctx sdk.Context) error {
	// TODO: get txHash, blockHeight, proof, data from somewhere
	txHash := []byte("0x123")
	blockHeight := int64(42)
	proof := lightclienttypes.Proof{}
	data := lightclienttypes.TxData{}

	verified, err := k.lightclientKeeper.VerifyTx(
		ctx,
		txHash,
		blockHeight,
		proof,
		data,
	)
	if err != nil {
		return err
	}
	if verified {
		ctx.Logger().Info("Transaction verified", "txHash", txHash)
	} else {
		ctx.Logger().Info("Transaction not verified", "txHash", txHash)
	}
	return nil
}
