package keeper

import (
	"context"
	"foochain/x/foo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
	LightclientKeeper types.LightclientKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, lightclientKeeper types.LightclientKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper, LightclientKeeper: lightclientKeeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) VerifyTx(goCtx context.Context, msg *types.VerifyTxRequest) (*types.VerifyTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var tx_hash = msg.TxHash
	var block_height = msg.BlockHeight
	var proof = msg.Proof
	var data = msg.Data

	verified, err := k.LightclientKeeper.VerifyTx(ctx, tx_hash, block_height, proof, data)
	if err != nil {
		return nil, err
	}
	return &types.VerifyTxResponse{Verified: verified}, nil
}
