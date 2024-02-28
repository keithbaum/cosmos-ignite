package keeper

import (
	"foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VerifyTx verifies a transaction against a header
func (k Keeper) VerifyTx(
	ctx sdk.Context,
	txHash []byte,
	blockHeight int64,
	proof types.Proof,
	data types.TxData,
) (bool, error) {
	// TODO: implement
	return false, nil
}
