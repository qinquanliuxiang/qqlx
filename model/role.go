package model

type Role struct {
	*MetaData
	Name        string    `gorm:"comment:角色名称;uniqueIndex;size:50" json:"name"`
	Description string    `gorm:"comment:角色描述;size:1024" json:"description"`
	Policys     []*Policy `gorm:"many2many:role_policys;" json:"policys,omitempty"`
	Users       []*User   `json:"users,omitempty"`
}
