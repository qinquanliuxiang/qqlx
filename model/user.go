package model

type User struct {
	*MetaData
	Name     string `gorm:"comment:用户名称;uniqueIndex;size:50"`
	Email    string `gorm:"comment:邮箱账号;uniqueIndex;size:100"`
	Role     *Role  // 用户角色
	RoleID   uint   `gorm:"comment:用户角色ID" json:"roleID"`
	Password string `gorm:"comment:用户密码;size:255"`
	Avatar   string `gorm:"comment:用户头像;size:1024"`
	Mobile   string `gorm:"comment:用户手机号;size:20"`
	Status   int    `gorm:"comment:用户状态,1:启用,0:禁用;default:1;size:1"`
}
