package keeper

import (
	"github.com/abhinav/interchainDEX/x/ibcdex/types"
)

var _ types.QueryServer = Keeper{}
