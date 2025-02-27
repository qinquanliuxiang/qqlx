package controller

import (
	"errors"
	"qqlx/base/apierrs"
	"qqlx/base/handler"
	"qqlx/base/reason"
	"qqlx/pkg/jwt"
	"qqlx/schema"
	"qqlx/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc *service.UserService
}

func NewUserController(userContr *service.UserService) *UserController {
	return &UserController{
		userSvc: userContr,
	}
}

func (userContr *UserController) Register(c *gin.Context) {
	req := new(schema.UserRegistryRequest)
	if handler.BindAndCheck(c, req) {
		return
	}

	if err := userContr.userSvc.RegistryUser(c, req); err != nil {
		handler.ResponseServerError(c, err)
		return
	}
	handler.ResponseSuccess(c, nil)
}

func (userContr *UserController) Login(c *gin.Context) {
	req := new(schema.UserLoginRequest)
	if handler.BindAndCheck(c, req) {
		return
	}

	res, err := userContr.userSvc.Login(c, req)
	if err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, res)
}

func (userContr *UserController) UpdatePassword(c *gin.Context) {
	req := new(schema.UserUpdatePasswordRequest)
	if handler.BindAndCheck(c, req) {
		return
	}

	if err := checkUserPermission(c, req.ID); err != nil {
		handler.ResponseForbidden(c, err)
		return
	}

	if err := userContr.userSvc.UpdatePassword(c, req); err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, nil)
}

func (userContr *UserController) Update(c *gin.Context) {
	req := new(schema.UserUpdateRequest)
	if handler.BindAndCheck(c, req) {
		return
	}

	if err := checkUserPermission(c, req.ID); err != nil {
		handler.ResponseForbidden(c, err)
		return
	}

	if err := userContr.userSvc.UpdateUser(c, req); err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, nil)
}

func (userContr *UserController) UpdateRole(c *gin.Context) {
	req := new(schema.UserUpdateRoleRequest)
	if handler.BindAndCheck(c, req) {
		return
	}
	if err := userContr.userSvc.UpdateUserRole(c, req); err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, nil)
}

func (userContr *UserController) Delete(c *gin.Context) {
	req := new(schema.IDRequest)
	if handler.BindAndCheck(c, req) {
		return
	}
	if err := userContr.userSvc.DeleteUser(c, req); err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, nil)
}

func (userContr *UserController) List(c *gin.Context) {
	req := new(schema.UserListRequest)
	if handler.BindAndCheck(c, req) {
		return
	}
	res, err := userContr.userSvc.ListUser(c, req)
	if err != nil {
		handler.ResponseServerError(c, err)
		return
	}
	handler.ResponseSuccess(c, res)
}

func (userContr *UserController) GetById(c *gin.Context) {
	req := new(schema.IDRequest)
	if handler.BindAndCheck(c, req) {
		return
	}
	res, err := userContr.userSvc.GetUserBasicInfoByID(c, req)
	if err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, res)
}

func (userContr *UserController) GetMyInfo(c *gin.Context) {
	claims, err := jwt.GetMyCustomClaims(c)
	if err != nil {
		handler.ResponseUnauthorized(c, err)
		return
	}
	res, err := userContr.userSvc.GetUserBasicInfoByID(c, &schema.IDRequest{ID: claims.UserID})
	if err != nil {
		response(c, err)
		return
	}
	handler.ResponseSuccess(c, res)
}

// checkUserPermission
func checkUserPermission(c *gin.Context, reqID uint) error {
	claims, err := jwt.GetMyCustomClaims(c)
	if err != nil {
		return err
	}

	if claims.UserID != reqID {
		return apierrs.NewAuthError(reason.ErrPermission)
	}
	return nil
}

// response
// Return HTTP status code based on error
func response(c *gin.Context, err error) {
	if errors.Is(err, reason.ErrUserNotFound) {
		handler.ResponseNotFound(c, err)
		return
	} else if errors.Is(err, reason.ErrInvalidPassword) {
		handler.ResponseUnauthorized(c, err)
		return
	}
	handler.ResponseServerError(c, err)
}
