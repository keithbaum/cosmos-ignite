package foo_test

import (
	"testing"

	keepertest "foochain/testutil/keeper"
	"foochain/testutil/nullify"
	"foochain/x/foo"
	"foochain/x/foo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FooKeeper(t)
	foo.InitGenesis(ctx, *k, genesisState)
	got := foo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
