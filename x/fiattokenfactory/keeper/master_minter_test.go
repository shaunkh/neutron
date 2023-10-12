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

func createTestMasterMinter(keeper *keeper.Keeper, ctx sdk.Context) types.MasterMinter {
	item := types.MasterMinter{}
	keeper.SetMasterMinter(ctx, item)
	return item
}

func TestMasterMinterGet(t *testing.T) {
	keeper, ctx := keepertest.FiatTokenfactoryKeeper(t)
	item := createTestMasterMinter(keeper, ctx)
	rst, found := keeper.GetMasterMinter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
