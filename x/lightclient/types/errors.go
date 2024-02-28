package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lightclient module sentinel errors
var (
	ErrSample             = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidBlockHeight = sdkerrors.Register(ModuleName, 1101, "invalid block height")
	ErrInvalidMerkleRoot  = sdkerrors.Register(ModuleName, 1102, "invalid merkle root")
)
