package keeper

import (
	"context"
	"foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const AuthorizedAdmin = "admin_signer_address"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) StoreExternalChain1BlockMerleRoot(goCtx context.Context, msg *types.StoreExternalChain1BlockMerleRootRequest) (*types.StoreExternalChain1BlockMerleRootResponse, error) {
	if msg.Sender != AuthorizedAdmin {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid or unauthorized sender address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var block_height = msg.BlockHeight
	var merkle_root = msg.MerkleRoot

	err := k.storeExternalChain1MerkleRoot(ctx, block_height, merkle_root)
	if err != nil {
		return nil, err
	}
	return &types.StoreExternalChain1BlockMerleRootResponse{}, nil
}
