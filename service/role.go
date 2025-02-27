package service

import (
	"context"
	"errors"
	"fmt"
	"qqlx/base"
	"qqlx/base/apierrs"
	"qqlx/base/helpers"
	"qqlx/model"
	"qqlx/pkg/permissions"
	"qqlx/schema"
)

type RoleService struct {
	generalRoleStore  base.GeneralRoleStore
	getPolicyStore    base.GetPolicyStore
	appendPolicyStore base.AssociationPolicyer
	casbinStore       permissions.GeneralAuthorizer
}

func NewRoleService(
	generalRoleStore base.GeneralRoleStore,
	policyStore base.GetPolicyStore,
	appendStore base.AssociationPolicyer,
	casbinStore permissions.GeneralAuthorizer,
) *RoleService {
	return &RoleService{
		generalRoleStore:  generalRoleStore,
		getPolicyStore:    policyStore,
		casbinStore:       casbinStore,
		appendPolicyStore: appendStore,
	}
}

func (r *RoleService) CreateRole(ctx context.Context, req *schema.RoleCreateRequest) (err error) {
	role := &model.Role{
		Name:        req.Name,
		Description: req.Desc,
	}
	return r.generalRoleStore.Create(ctx, role)
}

// DeleteRole 删除角色
func (r *RoleService) DeleteRole(ctx context.Context, req *schema.IDRequest) (err error) {
	role, err := r.generalRoleStore.GetRoleByID(ctx, req.ID, base.WithRoleUsers())
	if err != nil {
		return err
	}
	if role.Users != nil && len(role.Users) > 0 {
		var userNames []string
		for _, user := range role.Users {
			userNames = append(userNames, user.Name)
		}
		return apierrs.NewDeleteError(fmt.Errorf("failed to delete role, role has users: %s", userNames))
	}
	return r.generalRoleStore.Delete(ctx, role)
}

// UpdateRoleDesc 更新角色描述信息
func (r *RoleService) UpdateRoleDesc(ctx context.Context, req *schema.RoleUpdateRequest) (err error) {
	role, err := r.generalRoleStore.GetRoleByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if role.Description == req.Desc {
		return nil
	}

	role.Description = req.Desc
	return r.generalRoleStore.Save(ctx, role)
}

// AddByPolicy 增加 casbin 角色权限
func (r *RoleService) AddByPolicy(ctx context.Context, req *schema.RoleUpdatePolicyRequest) (err error) {
	role, err := r.generalRoleStore.GetRoleByID(ctx, req.RoleID, base.WithRolePolicys())
	if err != nil {
		return err
	}
	reqPolicys, err := r.getPolicyStore.GetPolicyByIDs(ctx, req.PolicyID)
	if err != nil {
		return err
	}
	if len(reqPolicys) != len(req.PolicyID) {
		return apierrs.NewUpdateError(errors.New("failed to update policy, policy not exists"))
	}

	err = r.appendPolicyStore.AppendPolicy(ctx, role, reqPolicys)
	if err != nil {
		return err
	}

	// 更新 casbin 策略
	save := helpers.GetCasbinRole(role.Name, reqPolicys)
	return r.casbinStore.CreateRolePolicys(ctx, save)
}

// DeleteByPolicy 删除 casbin 角色权限
func (r *RoleService) DeleteByPolicy(ctx context.Context, req *schema.RoleDeltePolicyRequest) (err error) {
	role, err := r.generalRoleStore.GetRoleByID(ctx, req.RoleID, base.WithRolePolicys())
	if err != nil {
		return err
	}
	reqPolicys, err := r.getPolicyStore.GetPolicyByIDs(ctx, req.PolicyID)
	if err != nil {
		return err
	}

	if len(reqPolicys) != len(req.PolicyID) {
		return apierrs.NewUpdateError(errors.New("failed to update policy, policy not exists"))
	}

	if err := r.appendPolicyStore.DeletePolicy(ctx, role, reqPolicys); err != nil {
		return err
	}

	deleteRole := helpers.GetCasbinRole(role.Name, reqPolicys)
	return r.casbinStore.DeleteRolePolicys(ctx, deleteRole)
}

func (r *RoleService) GetRoleByID(ctx context.Context, req *schema.IDRequest) (role *model.Role, err error) {
	return r.generalRoleStore.GetRoleByID(ctx, req.ID)
}

func (r *RoleService) ListRole(ctx context.Context, req *schema.RoleListRequest) (data *schema.RoleListResponse, err error) {
	total, roles, err := r.generalRoleStore.List(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &schema.RoleListResponse{
		Total: total,
		ListRequest: &schema.ListRequest{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Items: roles,
	}, nil
}
