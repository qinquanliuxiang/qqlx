package model

import "gorm.io/plugin/soft_delete"

const (
	StatusDisabled = iota
	StatusEnabled
)

type MetaData struct {
	ID        uint                  `gorm:"primarykey" json:"id"`
	CreatedAt uint                  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt uint                  `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:;index" json:"deletedAt"`
}
