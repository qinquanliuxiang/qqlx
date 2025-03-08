package schema

import "qqlx/model"

type RoleCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

type RoleUpdateRequest struct {
	ID   int    `json:"id" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

type RoleListRequest struct {
	*ListRequest
}

type RoleListResponse struct {
	Total int64 `json:"total"`
	*ListRequest
	Items []*model.Role `json:"items"`
}

type RoleUpdatePolicyRequest struct {
	RoleID   int    `json:"roleID" binding:"required"`
	PolicyID []uint `json:"policyID" binding:"required"`
}

type RoleDeltePolicyRequest struct {
	RoleID   int    `json:"roleID" binding:"required"`
	PolicyID []uint `json:"policyID" binding:"required"`
}
