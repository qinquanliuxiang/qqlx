package model

type Policy struct {
	*MetaData
	Name        string  `gorm:"comment:名称;size:50;uniqueIndex:idx_policy_name_path_method" json:"name"`
	Path        string  `gorm:"comment:路径;size:128;uniqueIndex:idx_policy_name_path_method" json:"path"`
	Method      string  `gorm:"comment:方法;size:10;uniqueIndex:idx_policy_name_path_method" json:"method"`
	Description string  `gorm:"comment:描述;size:1024" json:"description"`
	Roles       []*Role `gorm:"many2many:role_policys;" json:"roles,omitempty"`
}

func (p *Policy) TableName() string {
	return "policys"
}
