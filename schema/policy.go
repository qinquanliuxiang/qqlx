package schema

import "qqlx/model"

type PolicyCreateRequest struct {
	Name   string `json:"name" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
type PolicyDeleteRequest struct {
	*IDRequest
}

type PolicyGetByIDRequest struct {
	*IDRequest
}

type PolicyGetByNameRequest struct {
	Name string `json:"name" binding:"required"`
}

type PolicyUpdateRequest struct {
	*IDRequest
	Desc string `json:"desc" binding:"required"`
}

type PolicyListRequest struct {
	*ListRequest
}

type PolicyListResponse struct {
	Total int64
	*ListRequest
	Items []*model.Policy
}
