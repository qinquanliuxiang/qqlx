package store

import (
	"qqlx/base"
	"qqlx/base/data"
	"qqlx/pkg/permissions"

	"github.com/google/wire"
)

var ProviderStore = wire.NewSet(
	wire.Bind(new(base.Cache), new(*data.Redis)),
	wire.Bind(new(base.GeneralUserStore), new(*UserStore)),
	wire.Bind(new(base.GeneralRoleStore), new(*GeneralRoleStore)),
	wire.Bind(new(base.GeneralPolicyStore), new(*PolicyStore)),
	wire.Bind(new(base.AssociationPolicyer), new(*RoleAssociationStore)),
	wire.Bind(new(permissions.GeneralAuthorizer), new(*permissions.GeneralAuthorizStore)),
	wire.Bind(new(base.GetRoleStore), new(*GeneralRoleStore)),
	wire.Bind(new(base.GetPolicyStore), new(*PolicyStore)),
	data.CreateRDB,
	data.NewDB,
	data.NewRedis,
	NewUserStore,
	NewRoleStore,
	NewPolicyStore,
	NewRoleAssociationStore,
	permissions.NewGeneralAuthorizStore,
	permissions.InitCasbin,
)
