package schema

import "qqlx/model"

type UserRegistryRequest struct {
	Name         string `json:"name" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Avatar       string `json:"avatar"`
	Email        string `json:"email" binding:"email"`
	LarkUsername string `json:"larkUsername"`
	Mobile       string `json:"mobile"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required_without=Email"`
	Email    string `json:"email" binding:"required_without=Username,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginResponse struct {
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

type UserUpdatePasswordRequest struct {
	*IDRequest
	OldPassword string `json:"oldPassword" binding:"required,min=8"`
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}

type UserUpdateRequest struct {
	*IDRequest
	Avatar       string `json:"avatar"`
	Email        string `json:"email" binding:"email"`
	LarkUsername string `json:"larkUsername"`
	Mobile       string `json:"mobile"`
}

type UserResponse struct {
	*model.MetaData
	Name   string      `json:"name"`
	Avatar string      `json:"avatar"`
	Email  string      `json:"email" binding:"email"`
	Mobile string      `json:"mobile"`
	RoleID uint        `json:"roleID"`
	Role   *model.Role `json:"role,omitempty"`
	Status int         `json:"status"`
}

type UserListRequest struct {
	*ListRequest
}

type UserListResponse struct {
	Total int64 `json:"total"`
	*ListRequest
	Items []*UserResponse `json:"items"`
}

type UserUpdateRoleRequest struct {
	UserID uint `json:"userID" binding:"required,gte=1"`
	RoleID uint `json:"roleID" binding:"required,gte=1"`
}
