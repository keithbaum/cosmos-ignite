package keeper

import (
	"context"
	"foochain/x/foo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
	LightclientKeeper types.LightclientKeeper
}

func NewQueryServerImpl(keeper Keeper, lightclientKeeper types.LightclientKeeper) types.QueryServer {
	return &queryServer{Keeper: keeper, LightclientKeeper: lightclientKeeper}
}

func (k queryServer) VerifyTx(goCtx context.Context, msg *types.QueryVerifyTxRequest) (*types.QueryVerifyTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var txHash = []byte(msg.TxHash)
	var blockHeight = msg.BlockHeight
	var proof = msg.Proof
	var data = msg.TxData

	verified, err := k.LightclientKeeper.VerifyTx(ctx, txHash, blockHeight, *proof, *data)
	if err != nil {
		return nil, err
	}
	return &types.QueryVerifyTxResponse{Verified: verified}, nil
}
