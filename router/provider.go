package router

import "github.com/google/wire"

var ProviderRouter = wire.NewSet(
	NewApiRoute,
)
