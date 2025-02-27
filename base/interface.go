package base

import (
	"context"
	"time"
	"qqlx/model"

	"gorm.io/gorm"
)

type UserQueryOption func(query *gorm.DB)

// WithUserRole 设置预加载 Roles
func WithUserRole() UserQueryOption {
	return func(query *gorm.DB) {
		query.Preload("Role")
	}
}

func WithUserPolicys() UserQueryOption {
	return func(query *gorm.DB) {
		query.Preload("Role.Policys")
	}
}

// GeneralUserStore 用户CRUD
type GeneralUserStore interface {
	Create(ctx context.Context, user *model.User) (err error)
	Save(ctx context.Context, user *model.User) (err error)
	Delete(ctx context.Context, user *model.User) (err error)
	List(ctx context.Context, page, pageSize int) (total int64, users []*model.User, err error)
	GetUserStore
}

// GetUserStore 获取用户
type GetUserStore interface {
	GetUserByName(ctx context.Context, name string, options ...UserQueryOption) (user *model.User, err error)
	GetUserByID(ctx context.Context, id uint, options ...UserQueryOption) (user *model.User, err error)
	GetUserByEmail(ctx context.Context, email string, options ...UserQueryOption) (user *model.User, err error)
}

type RoleQueryOptions struct {
	WithPolicys bool
	WithUsers   bool
}
type RoleQueryOption func(query *gorm.DB)

// WithRolePolicys 设置预加载 Policys
func WithRolePolicys() RoleQueryOption {
	return func(query *gorm.DB) {
		query.Preload("Policys")
	}
}

// WithRoleUsers 设置预加载 Users
func WithRoleUsers() RoleQueryOption {
	return func(query *gorm.DB) {
		query.Preload("Users")
	}
}

// GeneralRoleStore 角色CRUD
type GeneralRoleStore interface {
	Create(ctx context.Context, role *model.Role) (err error)
	Save(ctx context.Context, role *model.Role) (err error)
	Delete(ctx context.Context, role *model.Role) (err error)
	List(ctx context.Context, page, pageSize int) (total int64, roles []*model.Role, err error)
	GetRoleStore
}

// GetRoleStore 获取角色
type GetRoleStore interface {
	GetRoleByID(ctx context.Context, id uint, options ...RoleQueryOption) (role *model.Role, err error)
	GetRoleByName(ctx context.Context, name string, options ...RoleQueryOption) (role *model.Role, err error)
}

// AssociationPolicyer 角色关联策略
type AssociationPolicyer interface {
	AppendPolicy(ctx context.Context, role *model.Role, policy []*model.Policy) (err error)
	ReplacePolicy(ctx context.Context, role *model.Role, policy []*model.Policy) (err error)
	DeletePolicy(ctx context.Context, role *model.Role, policy []*model.Policy) (err error)
}

type PolicyQueryOption func(query *gorm.DB)

// WithPolicyRoles 设置预加载 Policys
func WithPolicyRoles() PolicyQueryOption {
	return func(query *gorm.DB) {
		query.Preload("Roles")
	}
}

// GeneralPolicyStore 策略CRUD
type GeneralPolicyStore interface {
	Create(ctx context.Context, policy *model.Policy) (err error)
	Save(ctx context.Context, policy *model.Policy) (err error)
	Delete(ctx context.Context, policy *model.Policy) (err error)
	List(ctx context.Context, page, pageSize int) (total int64, policys []*model.Policy, err error)
	GetPolicyStore
}

type GetPolicyStore interface {
	GetPolicyByID(ctx context.Context, id uint, options ...PolicyQueryOption) (policy *model.Policy, err error)
	GetPolicyByName(ctx context.Context, name string, options ...PolicyQueryOption) (policy *model.Policy, err error)
	GetPolicyByIDs(ctx context.Context, ids []uint) (policys []*model.Policy, err error)
}

// Cache 缓存
type Cache interface {
	GetString(ctx context.Context, key string) (data string, err error)
	// SetString 设置字符串
	//
	// expireTime 过期时间, nil 使用默认过期时间; &data.NeverExpires 表示永不过期
	SetString(ctx context.Context, key, value string, expireTime *time.Duration) (err error)
	GetInt64(ctx context.Context, key string) (data int64, err error)
	// SetInt64 设置整数
	//
	// expireTime 过期时间, nil 使用默认过期时间; &data.NeverExpires 表示永不过期
	SetInt64(ctx context.Context, key string, value int64, expireTime *time.Duration) (err error)
	Del(ctx context.Context, key string) (err error)
	Flush(ctx context.Context) (err error)
}
