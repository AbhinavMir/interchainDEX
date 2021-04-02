package keeper

import (
	"github.com/abhinav/interchainDEX/x/interchainDEX/types"
)

var _ types.QueryServer = Keeper{}
