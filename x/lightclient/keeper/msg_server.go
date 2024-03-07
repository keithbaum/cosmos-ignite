package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const AuthorizedAdmin = "cosmos1q0pha0zu457p5pz68y95tzaljryy3ln59sum5t"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) StoreExternalChain1TxHeaders(goCtx context.Context, msg *types.MsgStoreExternalChain1TxHeadersRequest) (*types.MsgStoreExternalChain1TxHeadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signerAddress, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	authorizedAdmin, _ := sdk.AccAddressFromHexUnsafe(AuthorizedAdmin)
	if signerAddress.Equals(authorizedAdmin) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid or unauthorized sender address")
	}

	var block_height = msg.BlockHeight
	var merkle_root = msg.MerkleRoot

	err = k.storeExternalChain1MerkleRoot(ctx, block_height, []byte(merkle_root))
	if err != nil {
		return nil, err
	}
	return &types.MsgStoreExternalChain1TxHeadersResponse{}, nil
}
