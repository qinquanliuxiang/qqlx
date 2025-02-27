package controller

import (
	"qqlx/base/handler"
	"qqlx/schema"
	"qqlx/service"

	"github.com/gin-gonic/gin"
)

type PolicyController struct {
	policySvc *service.PolicyService
}

func NewPolicyController(policySvc *service.PolicyService) *PolicyController {
	return &PolicyController{
		policySvc: policySvc,
	}
}

func (c *PolicyController) Create(ctx *gin.Context) {
	req := new(schema.PolicyCreateRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.policySvc.CreatePolicy(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *PolicyController) Delete(ctx *gin.Context) {
	req := new(schema.PolicyDeleteRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.policySvc.DeletePolicy(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *PolicyController) Update(ctx *gin.Context) {
	req := new(schema.PolicyUpdateRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	if err := c.policySvc.UpdatePolicy(ctx, req); err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, nil)
}

func (c *PolicyController) GetByID(ctx *gin.Context) {
	req := new(schema.PolicyGetByIDRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	res, err := c.policySvc.GetPolicyByID(ctx, req)
	if err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, res)
}

func (c *PolicyController) List(ctx *gin.Context) {
	req := new(schema.PolicyListRequest)
	if handler.BindAndCheck(ctx, req) {
		return
	}
	res, err := c.policySvc.List(ctx, req)
	if err != nil {
		handler.ResponseServerError(ctx, err)
		return
	}
	handler.ResponseSuccess(ctx, res)
}
