package controller

import (
	"qqlx/base/handler"
	"qqlx/schema"
	"qqlx/service"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleSvc *service.RoleService
}

func NewRoleController(roleSvc *service.RoleService) *RoleController {
	return &RoleController{
		roleSvc: roleSvc,
	}
}

func (c *RoleController) Create(ctx *gin.Context) {
	req := new(schema.RoleCreateRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.roleSvc.CreateRole(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *RoleController) Delete(ctx *gin.Context) {
	req := new(schema.IDRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.roleSvc.DeleteRole(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *RoleController) UpdateDesc(ctx *gin.Context) {
	req := new(schema.RoleUpdateRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.roleSvc.UpdateRoleDesc(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *RoleController) AddRoleByPolicy(ctx *gin.Context) {
	req := new(schema.RoleUpdatePolicyRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.roleSvc.AddByPolicy(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *RoleController) DeleteRoleByPolicy(ctx *gin.Context) {
	req := new(schema.RoleDeltePolicyRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.roleSvc.DeleteByPolicy(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *RoleController) GetByID(ctx *gin.Context) {
	req := new(schema.IDRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	res, err := c.roleSvc.GetRoleByID(ctx, req)
	if err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, res)
}

func (c *RoleController) List(ctx *gin.Context) {
	req := new(schema.RoleListRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	res, err := c.roleSvc.ListRole(ctx, req)
	if err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, res)
}
