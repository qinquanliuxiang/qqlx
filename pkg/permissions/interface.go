package permissions

import (
	"context"
	"fmt"
	"qqlx/base/apierrs"

	"github.com/casbin/casbin/v2"
)

type GeneralAuthorizer interface {
	// 创建role拥有的权限
	GetRolePolicyByName(ctx context.Context, role string) (policys [][]string, err error)
	// CreateRolePolicy 创建role拥有的权限
	//
	// policys [][]string{role, path, method}
	CreateRolePolicys(ctx context.Context, policys [][]string) (err error)
	// DeleteRolePolicy 删除role拥有的权限
	//
	// policys [][]string{role, path, method}
	DeleteRolePolicys(ctx context.Context, policys [][]string) (err error)
	// UpdateRolePolicy 更新role拥有的权限
	//
	// policys [][]string{role, path, method}
	UpdateRolePolicys(ctx context.Context, roleName string, policys [][]string) (err error)
}

type GeneralAuthorizStore struct {
	enforcer *casbin.Enforcer
}

func NewGeneralAuthorizStore(enforcer *casbin.Enforcer) *GeneralAuthorizStore {
	return &GeneralAuthorizStore{
		enforcer: enforcer,
	}
}
func (a *GeneralAuthorizStore) GetRolePolicyByName(ctx context.Context, role string) (policys [][]string, err error) {
	policys, err = a.enforcer.GetFilteredPolicy(0, role)
	if err != nil {
		return nil, apierrs.NewGetError(fmt.Errorf("failed to get policy, %w", err))
	}
	return policys, nil
}

func (a *GeneralAuthorizStore) CreateRolePolicys(ctx context.Context, policys [][]string) (err error) {
	ok, err := a.enforcer.AddPolicies(policys)
	if err != nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to create policy, %w", err))
	}
	if !ok {
		return apierrs.NewCreateError(fmt.Errorf("failed to create policy, policy already exists"))
	}
	return nil
}

func (a *GeneralAuthorizStore) DeleteRolePolicys(ctx context.Context, policys [][]string) (err error) {
	ok, err := a.enforcer.RemovePolicies(policys)
	if err != nil {
		return apierrs.NewDeleteError(fmt.Errorf("failed to delete policy, %w", err))
	}
	if !ok {
		return apierrs.NewDeleteError(fmt.Errorf("failed to delete policy, policy not exists"))
	}
	return nil
}

func (a *GeneralAuthorizStore) UpdateRolePolicys(ctx context.Context, roleName string, policys [][]string) (err error) {
	oldPolicys, err := a.GetRolePolicyByName(ctx, roleName)
	if err != nil {
		return err
	}

	ok, err := a.enforcer.UpdatePolicies(oldPolicys, policys)
	if err != nil {
		return apierrs.NewUpdateError(fmt.Errorf("failed to update policy, %w", err))
	}
	if !ok {
		return apierrs.NewUpdateError(fmt.Errorf("failed to update policy, policy not exists"))
	}
	return nil
}

type Authorizer interface {
	EnforceWithCtx(ctx context.Context, sub, obj, act string) (ok bool, err error)
}
type Authoriz struct {
	enforcer *casbin.Enforcer
}

func NewAuthoriz(enforcer *casbin.Enforcer) *Authoriz {
	return &Authoriz{
		enforcer: enforcer,
	}
}

func (a *Authoriz) EnforceWithCtx(_ context.Context, sub, obj, act string) (ok bool, err error) {
	ok, err = a.enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, apierrs.NewAuthError(fmt.Errorf("failed to enforce, %w", err))
	}
	return ok, nil
}
