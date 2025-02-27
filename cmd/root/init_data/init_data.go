package init_data

import (
	"context"
	"fmt"
	"qqlx/base/conf"
	"qqlx/base/constant"
	"qqlx/base/data"
	"qqlx/base/helpers"
	"qqlx/base/logger"
	"qqlx/model"
	"qqlx/pkg/permissions"
	"qqlx/schema"
	"qqlx/service"
	"qqlx/store"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init data",
	Long:  "init data",
	PreRun: func(cmd *cobra.Command, args []string) {
		helpers.PreRun(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cf, err := cmd.Flags().GetString(constant.FlagConfigPath)
		if err != nil {
			panic(err)
		}
		casbin, err := cmd.Flags().GetString(constant.FlagCasbinModePath)
		if err != nil {
			panic(err)
		}
		initData(cf, casbin)
	},
}

func initData(cf, casbinFilePath string) {
	conf.LoadConfig(cf)
	logger.InitLogger()
	defer zap.S().Sync()
	db, close, err := data.NewDB()
	defer close()
	if err != nil {
		panic(err)
	}
	enforcer, err := permissions.InitCasbin(casbinFilePath)
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.Role{}, &model.Policy{}); err != nil {
		panic(err)
	}
	generalAuthorizRepo := permissions.NewGeneralAuthorizStore(enforcer)
	var generalAuthorizer permissions.GeneralAuthorizer
	generalAuthorizer = generalAuthorizRepo
	userRepo := store.NewUserStore(db)
	roleRepo := store.NewRoleStore(db)
	userSvc := service.NewUserService(userRepo, roleRepo, nil, generalAuthorizer)

	adminPolicys := &model.Policy{
		Name:   "admin",
		Path:   "*",
		Method: "*",
	}

	viewPolicys := &model.Policy{
		Name:   "view",
		Path:   "*",
		Method: "GET",
	}

	userPolicys := []*model.Policy{
		{
			Name:        "deleteUser",
			Path:        "/api/v1/user/deleteUser",
			Method:      "POST",
			Description: "删除用户",
		},
		{
			Name:        "getUserLis",
			Path:        "/api/v1/user/getUserList",
			Method:      "GET",
			Description: "获取用户列表",
		},
		{
			Name:        "updateUserRole",
			Path:        "/api/v1/user/updateUserRole",
			Method:      "POST",
			Description: "更新用户角色",
		},
		{
			Name:        "getUserList",
			Path:        "/api/v1/user/getUserList",
			Method:      "GET",
			Description: "获取用户列表",
		},
		{
			Name:        "getUserById",
			Path:        "/api/v1/user/getUserById",
			Method:      "GET",
			Description: "获取用户信息",
		},
	}

	rolePolicys := []*model.Policy{
		{
			Name:        "getRoleById",
			Path:        "/api/v1/role/getRoleById",
			Method:      "GET",
			Description: "获取角色信息",
		},
		{
			Name:        "getRoleList",
			Path:        "/api/v1/role/getRoleList",
			Method:      "GET",
			Description: "获取角色列表",
		},
		{
			Name:        "createRole",
			Path:        "/api/v1/role/createRole",
			Method:      "POST",
			Description: "创建角色",
		},
		{
			Name:        "deleteRole",
			Path:        "/api/v1/role/deleteRole",
			Method:      "POST",
			Description: "删除角色",
		},
		{
			Name:        "updateRole",
			Path:        "/api/v1/role/updateRole",
			Method:      "POST",
			Description: "更新角色描述",
		},
		{
			Name:        "addRoleByPolicy",
			Path:        "/api/v1/role/addRoleByPolicy",
			Method:      "POST",
			Description: "增加角色权限",
		},
		{
			Name:        "deleteRoleByPolicy",
			Path:        "/api/v1/role/deleteRoleByPolicy",
			Method:      "POST",
			Description: "删除角色权限",
		},
	}

	policyPolisys := []*model.Policy{
		{
			Name:        "getPolicyById",
			Path:        "/api/v1/policy/getPolicyById",
			Method:      "GET",
			Description: "获取策略信息",
		},
		{
			Name:        "getPolicyList",
			Path:        "/api/v1/policy/getPolicyList",
			Method:      "GET",
			Description: "获取策略列表",
		},
		{
			Name:        "createPolicy",
			Path:        "/api/v1/policy/createPolicy",
			Method:      "POST",
			Description: "创建策略",
		},
		{
			Name:        "deletePolicy",
			Path:        "/api/v1/policy/deletePolicy",
			Method:      "POST",
			Description: "删除策略",
		},
		{
			Name:        "updatePolicy",
			Path:        "/api/v1/policy/updatePolicy",
			Method:      "POST",
			Description: "更新策略",
		},
	}

	savePolicys := make([]*model.Policy, 0, 2+len(userPolicys)+len(rolePolicys)+len(policyPolisys))
	savePolicys = append(savePolicys, adminPolicys)
	savePolicys = append(savePolicys, viewPolicys)
	savePolicys = append(savePolicys, userPolicys...)
	savePolicys = append(savePolicys, rolePolicys...)
	savePolicys = append(savePolicys, policyPolisys...)
	if err := db.Create(savePolicys).Error; err != nil {
		panic(fmt.Errorf("failed to create policys, %w", err))
	}

	role := []*model.Role{
		{
			Name:        "admin",
			Description: "超级管理员",
		},
		{
			Name:        "view",
			Description: "查看",
		},
	}
	if err := db.Create(role).Error; err != nil {
		panic(fmt.Errorf("failed to create roles, %w", err))
	}

	if err := userSvc.RegistryUser(context.Background(), &schema.UserRegistryRequest{
		Name:     "admin",
		Password: "12345678",
		Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Mobile:   "13800000000",
	}); err != nil {
		panic(err)
	}

	user, err := userRepo.GetUserByName(context.TODO(), "admin")
	if err != nil {
		panic(err)
	}
	var adminRole *model.Role
	if err := db.Model(adminRole).Take(&adminRole, "name = ?", "admin").Error; err != nil {
		panic(fmt.Errorf("failed to get admin role, %w", err))
	}
	if err := db.Model(&user).Association("Role").Append(adminRole); err != nil {
		panic(fmt.Errorf("failed to append role, %w", err))
	}

	var p *model.Policy
	if err := db.Model(&p).First(&p, "name = ?", "admin").Error; err != nil {
		panic(fmt.Errorf("failed to get admin policy, %w", err))
	}
	if err := db.Model(adminRole).Association("Policys").Append(p); err != nil {
		panic(fmt.Errorf("failed to append policy, %w", err))
	}

	save := make([][]string, 0, 1)
	save = append(save, []string{p.Name, p.Path, p.Method})
	if err := generalAuthorizer.CreateRolePolicys(context.TODO(), save); err != nil {
		panic(fmt.Errorf("failed to save role policys, %w", err))
	}

	var viewRole *model.Role
	if err := db.Model(viewRole).Take(&viewRole, "name = ?", "view").Error; err != nil {
		panic(fmt.Errorf("failed to get view role, %w", err))
	}
	var viewPolicy *model.Policy
	if err := db.Model(viewPolicy).First(&viewPolicy, "name = ?", "view").Error; err != nil {
		panic(fmt.Errorf("failed to get view policy, %w", err))
	}
	if err := db.Model(viewRole).Association("Policys").Append(viewPolicy); err != nil {
		panic(fmt.Errorf("viewPolicy failed to append viewRole, %w", err))
	}

	save = make([][]string, 0, 1)
	save = append(save, []string{viewPolicy.Name, viewPolicy.Path, viewPolicy.Method})
	if err := generalAuthorizer.CreateRolePolicys(context.TODO(), save); err != nil {
		panic(err)
	}
}
