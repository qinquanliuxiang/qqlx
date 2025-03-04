package router

import (
	"qqlx/base/middleware"
	"qqlx/controller"

	"github.com/gin-gonic/gin"
)

type ApiRoute struct {
	userController   *controller.UserController
	roleController   *controller.RoleController
	policyController *controller.PolicyController
}

func NewApiRoute(
	userContr *controller.UserController,
	roleContr *controller.RoleController,
	policyController *controller.PolicyController,
) *ApiRoute {
	return &ApiRoute{
		userController:   userContr,
		roleController:   roleContr,
		policyController: policyController,
	}
}

func (a *ApiRoute) RegisterApiUserRoute(r *gin.RouterGroup, authorization *middleware.AuthorizationMiddleware) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", a.userController.Register)
		userGroup.POST("/login", a.userController.Login)
		userGroup.Use(middleware.Authentication())
		{
			userGroup.POST("/updatePassword", a.userController.UpdatePassword)
			userGroup.POST("/updateUser", a.userController.Update)
			userGroup.GET("/getUserInfo", a.userController.GetMyInfo)
			userGroup.POST("/deleteUser", authorization.Authorization(), a.userController.Delete)
			userGroup.GET("/getUserList", authorization.Authorization(), a.userController.List)
			userGroup.GET("/getUserById", authorization.Authorization(), a.userController.GetById)
			userGroup.GET("/getUserByEmail", authorization.Authorization(), a.userController.GetByEmail)
			userGroup.POST("/updateUserRole", authorization.Authorization(), a.userController.UpdateRole)
		}
	}
}

func (a *ApiRoute) RegisterApiRoleRoute(r *gin.RouterGroup, authorization *middleware.AuthorizationMiddleware) {
	roleGroup := r.Group("/role")
	roleGroup.Use(middleware.Authentication(), authorization.Authorization())
	roleGroup.GET("/getRoleById", a.roleController.GetByID)
	roleGroup.GET("/getRoleList", a.roleController.List)
	roleGroup.POST("/createRole", a.roleController.Create)
	roleGroup.POST("/deleteRole", a.roleController.Delete)
	roleGroup.POST("/updateRole", a.roleController.UpdateDesc)
	roleGroup.POST("/addRoleByPolicy", a.roleController.AddRoleByPolicy)
	roleGroup.POST("/deleteRoleByPolicy", a.roleController.DeleteRoleByPolicy)
}

func (a *ApiRoute) RegisterApiPolicyRoute(r *gin.RouterGroup, authorization *middleware.AuthorizationMiddleware) {
	roleGroup := r.Group("/policy")
	roleGroup.Use(middleware.Authentication(), authorization.Authorization())
	roleGroup.GET("/getPolicyById", a.policyController.GetByID)
	roleGroup.GET("/getPolicyList", a.policyController.List)
	roleGroup.POST("/createPolicy", a.policyController.Create)
	roleGroup.POST("/deletePolicy", a.policyController.Delete)
	roleGroup.POST("/updatePolicy", a.policyController.Update)
}
