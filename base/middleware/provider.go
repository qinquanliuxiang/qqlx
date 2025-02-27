package middleware

import (
	"qqlx/base"
	"qqlx/pkg/permissions"
	"qqlx/store"

	"github.com/google/wire"
)

var ProviderMiddleware = wire.NewSet(
	wire.Bind(new(base.GetUserStore), new(*store.GetUserStore)),
	wire.Bind(new(permissions.Authorizer), new(*permissions.Authoriz)),
	store.NewGetUserStore,
	permissions.NewAuthoriz,
	NewAuthorization,
)
