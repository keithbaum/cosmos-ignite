package keeper

import (
	"foochain/x/lightclient/types"
)

var _ types.QueryServer = Keeper{}
