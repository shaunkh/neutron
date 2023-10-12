package keeper

import (
	"github.com/neutron-org/neutron/x/fiattokenfactory/types"
)

var _ types.QueryServer = Keeper{}
