package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/neutron-org/neutron/testutil/keeper"
	"github.com/neutron-org/neutron/testutil/nullify"
	"github.com/neutron-org/neutron/x/fiattokenfactory/keeper"
	"github.com/neutron-org/neutron/x/fiattokenfactory/types"
)

func createTestPauser(keeper *keeper.Keeper, ctx sdk.Context) types.Pauser {
	item := types.Pauser{}
	keeper.SetPauser(ctx, item)
	return item
}

func TestPauserGet(t *testing.T) {
	keeper, ctx := keepertest.FiatTokenfactoryKeeper(t)
	item := createTestPauser(keeper, ctx)
	rst, found := keeper.GetPauser(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
