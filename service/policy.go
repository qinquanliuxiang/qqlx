package service

import (
	"context"
	"fmt"
	"qqlx/base"
	"qqlx/base/apierrs"
	"qqlx/model"
	"qqlx/schema"
)

type PolicyService struct {
	policyStore base.GeneralPolicyStore
}

func NewPolicyService(policyStore base.GeneralPolicyStore) *PolicyService {
	return &PolicyService{
		policyStore: policyStore,
	}
}

func (s *PolicyService) CreatePolicy(ctx context.Context, req *schema.PolicyCreateRequest) (err error) {
	return s.policyStore.Create(ctx, &model.Policy{
		Name:        req.Name,
		Path:        req.Path,
		Method:      req.Method,
		Description: req.Desc,
	})
}

// DeletePolicy 删除策略
func (s *PolicyService) DeletePolicy(ctx context.Context, req *schema.PolicyDeleteRequest) (err error) {
	policy, err := s.policyStore.GetPolicyByID(ctx, req.ID, base.WithPolicyRoles())
	if err != nil {
		return err
	}
	if policy != nil && len(policy.Roles) > 0 {
		var roleNames []string
		for _, role := range policy.Roles {
			roleNames = append(roleNames, role.Name)
		}
		return apierrs.NewDeleteError(fmt.Errorf("failed to delete policy, policy has roles: %s", roleNames))
	}
	return s.policyStore.Delete(ctx, policy)
}

// UpdatePolicy 更新策略描述信息
func (s *PolicyService) UpdatePolicy(ctx context.Context, req *schema.PolicyUpdateRequest) (err error) {
	policy, err := s.policyStore.GetPolicyByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if policy.Description == req.Desc {
		return nil
	}

	policy.Description = req.Desc
	return s.policyStore.Save(ctx, policy)
}

func (s *PolicyService) GetPolicyByID(ctx context.Context, req *schema.PolicyGetByIDRequest) (policy *model.Policy, err error) {
	return s.policyStore.GetPolicyByID(ctx, req.ID)
}

func (s *PolicyService) GetPolicyByName(ctx context.Context, req *schema.PolicyGetByNameRequest) (policy *model.Policy, err error) {
	return s.policyStore.GetPolicyByName(ctx, req.Name)
}

func (s *PolicyService) List(ctx context.Context, req *schema.PolicyListRequest) (res *schema.PolicyListResponse, err error) {
	total, policys, err := s.policyStore.List(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &schema.PolicyListResponse{
		Total: total,
		ListRequest: &schema.ListRequest{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Items: policys,
	}, nil
}
