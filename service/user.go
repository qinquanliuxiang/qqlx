package service

import (
	"context"
	"errors"
	"fmt"
	"qqlx/base"
	"qqlx/base/apierrs"
	"qqlx/base/constant"
	"qqlx/base/data"
	"qqlx/base/logger"
	"qqlx/base/reason"
	"qqlx/model"
	"qqlx/pkg/jwt"
	"qqlx/pkg/permissions"
	"qqlx/schema"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userStore    base.GeneralUserStore
	getRoleStore base.GetRoleStore
	cache        base.Cache
	casbin       permissions.GeneralAuthorizer
}

func NewUserService(userStore base.GeneralUserStore, getRoleStore base.GetRoleStore, cache base.Cache, casbin permissions.GeneralAuthorizer) *UserService {
	return &UserService{
		userStore:    userStore,
		getRoleStore: getRoleStore,
		cache:        cache,
		casbin:       casbin,
	}
}

func (u *UserService) RegistryUser(ctx context.Context, req *schema.UserRegistryRequest) (err error) {
	if _, err = u.userStore.GetUserByName(ctx, req.Name); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		encryptPassword, err := u.encryptPassword(ctx, req.Password)
		if err != nil {
			return err
		}
		err = u.userStore.Create(ctx, &model.User{
			Name:     req.Name,
			Password: encryptPassword,
			Avatar:   req.Avatar,
			Email:    req.Email,
			Mobile:   req.Mobile,
		})
		if err != nil {
			return apierrs.NewCreateError(err)
		}
		return nil
	}
	return apierrs.NewCreateError(fmt.Errorf("user %s already exists", req.Name))
}

func (u *UserService) Login(ctx context.Context, req *schema.UserLoginRequest) (res *schema.UserLoginResponse, err error) {
	logger.WithContext(ctx, true).Debugf("user login, request: %+v", req)
	var user *model.User
	if req.Email != "" {
		user, err = u.userStore.GetUserByEmail(ctx, req.Email, base.WithUserRole())
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, apierrs.NewAuthError(reason.ErrUserNotFound)
			}
			return nil, err
		}
	}

	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return nil, reason.ErrUserNotFound
	}
	if !u.verifyPassword(ctx, req.Password, user.Password) {
		return nil, reason.ErrInvalidPassword
	}

	err = u.cache.SetString(ctx, constant.RoleCacheKeyPrefix+strconv.Itoa(user.ID), user.Role.Name, &data.NeverExpires)
	if err != nil {
		return nil, err
	}

	token, err := jwt.NewClaims(user.ID, user.Name).GenerateToken()
	if err != nil {
		return nil, err
	}
	res = &schema.UserLoginResponse{
		User: &schema.UserResponse{
			MetaData: &model.MetaData{ID: user.ID, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt},
			Name:     user.Name,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Mobile:   user.Mobile,
			Role:     user.Role,
			RoleID:   user.RoleID,
			Status:   user.Status,
		},
		Token: token,
	}
	return res, err
}

func (u *UserService) DeleteUser(ctx context.Context, req *schema.IDRequest) (err error) {
	var user *model.User
	user, err = u.userStore.GetUserByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return reason.ErrUserNotFound
	}

	user.Status = model.StatusDisabled
	err = u.userStore.Save(ctx, user)
	if err != nil {
		return err
	}
	return u.userStore.Delete(ctx, user)
}

func (u *UserService) UpdatePassword(ctx context.Context, userid int, req *schema.UserUpdatePasswordRequest) (err error) {
	var user *model.User

	if req.NewPassword != req.ConfirmPassword {
		return reason.ErrPasswordNotMatch
	}

	user, err = u.userStore.GetUserByID(ctx, userid)
	if err != nil {
		return err
	}
	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return reason.ErrUserNotFound
	}

	if !u.verifyPassword(ctx, req.OldPassword, user.Password) {
		return reason.ErrInvalidPassword
	}
	encryptPassword, err := u.encryptPassword(ctx, req.NewPassword)
	if err != nil {
		return err
	}
	user.Password = encryptPassword
	return u.userStore.Save(ctx, user)
}

func (u *UserService) UpdateUser(ctx context.Context, userid int, req *schema.UserUpdateRequest) (err error) {
	var user *model.User
	user, err = u.userStore.GetUserByID(ctx, userid)
	if err != nil {
		return err
	}
	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return reason.ErrUserNotFound
	}

	isUpdated := false
	if req.Name != "" && req.Name != user.Name {
		user.Name = req.Name
		isUpdated = true
	}
	if req.Mobile != "" && req.Mobile != user.Mobile {
		user.Mobile = req.Mobile
		isUpdated = true
	}
	if req.Avatar != "" && req.Avatar != user.Avatar {
		user.Avatar = req.Avatar
		isUpdated = true
	}
	if !isUpdated {
		return nil
	}
	return u.userStore.Save(ctx, user)
}

// UpdateUserRole 更新用户角色
func (u *UserService) UpdateUserRole(ctx context.Context, req *schema.UserUpdateRoleRequest) (err error) {
	var user *model.User
	user, err = u.userStore.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return reason.ErrUserNotFound
	}

	if user.RoleID == req.RoleID {
		return apierrs.NewUpdateError(fmt.Errorf("user %s already in role %s", user.Name, user.Role.Name))
	}
	user.RoleID = req.RoleID
	role, err := u.getRoleStore.GetRoleByID(ctx, req.RoleID)
	if err != nil {
		return err
	}
	if err := u.cache.Del(ctx, constant.RoleCacheKeyPrefix+user.Name); err != nil {
		return err
	}
	if err := u.userStore.Save(ctx, user); err != nil {
		return err
	}
	if err := u.cache.Del(ctx, constant.RoleCacheKeyPrefix+user.Name); err != nil {
		return err
	}

	return u.cache.SetString(ctx, constant.RoleCacheKeyPrefix+user.Name, role.Name, &data.NeverExpires)
}

func (u *UserService) GetUserBasicInfoByID(ctx context.Context, req *schema.IDRequest) (res *schema.UserResponse, err error) {
	user, err := u.userStore.GetUserByID(ctx, req.ID, base.WithUserRole(), base.WithUserPolicys())
	if err != nil {
		return nil, err
	}
	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return nil, reason.ErrUserNotFound
	}

	return &schema.UserResponse{
		MetaData: &model.MetaData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Name:   user.Name,
		Avatar: user.Avatar,
		Email:  user.Email,
		Status: user.Status,
		Mobile: user.Mobile,
		Role:   user.Role,
		RoleID: user.RoleID,
	}, nil
}

func (u *UserService) GetUserBasicInfoByEmail(ctx context.Context, req *schema.UserGetByEmailRequest) (res *schema.UserResponse, err error) {
	user, err := u.userStore.GetUserByEmail(ctx, req.Email, base.WithUserRole(), base.WithUserPolicys())
	if err != nil {
		return nil, err
	}
	if user.Status == model.StatusDisabled {
		logger.WithContext(ctx, true).Errorf("user has been disabled, userID = %v", user.ID)
		return nil, reason.ErrUserNotFound
	}

	return &schema.UserResponse{
		MetaData: &model.MetaData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Name:   user.Name,
		Avatar: user.Avatar,
		Email:  user.Email,
		Status: user.Status,
		Mobile: user.Mobile,
		Role:   user.Role,
		RoleID: user.RoleID,
	}, nil
}

func (u *UserService) ListUser(ctx context.Context, req *schema.UserListRequest) (data *schema.UserListResponse, err error) {
	if req.ListRequest == nil {
		return nil, reason.ErrInvalidRequest
	}
	total, users, err := u.userStore.List(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return u.formatUserList(req, total, users), nil
}

func (u *UserService) formatUserList(req *schema.UserListRequest, total int64, users []*model.User) *schema.UserListResponse {
	res := &schema.UserListResponse{
		ListRequest: &schema.ListRequest{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Total: total,
		Items: make([]*schema.UserResponse, 0, len(users)),
	}

	for _, user := range users {
		meta := &model.MetaData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		res.Items = append(res.Items, &schema.UserResponse{
			MetaData: meta,
			Name:     user.Name,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Mobile:   user.Mobile,
			Status:   user.Status,
			RoleID:   user.RoleID,
		})
	}

	return res
}

// encryptPassword 加密密码
func (us *UserService) encryptPassword(_ context.Context, Pass string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(Pass), bcrypt.DefaultCost)
	if err != nil {
		return "", apierrs.NewAuthError(fmt.Errorf("failed to encrypt password, %w", err))
	}
	return string(hashPwd), nil
}

// verifyPassword 验证密码
func (us *UserService) verifyPassword(_ context.Context, loginPass, userPass string) bool {
	if len(loginPass) == 0 && len(userPass) == 0 {
		return true
	}
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(loginPass))
	return err == nil
}
