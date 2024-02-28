package keeper

import (
	"foochain/x/foo/types"
)

var _ types.QueryServer = Keeper{}
