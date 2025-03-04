package schema

import "qqlx/model"

type UserRegistryRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"email"`
	Mobile   string `json:"mobile"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginResponse struct {
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

type UserUpdatePasswordRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required,min=8"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=8"`
}

type UserUpdateRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Mobile string `json:"mobile"`
}

type UserResponse struct {
	*model.MetaData
	Name   string      `json:"name"`
	Avatar string      `json:"avatar"`
	Email  string      `json:"email" binding:"email"`
	Mobile string      `json:"mobile"`
	RoleID int         `json:"roleID"`
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
	UserID int `json:"userID" binding:"required,gte=1"`
	RoleID int `json:"roleID" binding:"required,gte=1"`
}

type UserGetByEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
}
