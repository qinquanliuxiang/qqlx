package model

import "gorm.io/plugin/soft_delete"

const (
	StatusDisabled = iota
	StatusEnabled
)

type MetaData struct {
	ID        int                   `gorm:"primarykey" json:"id"`
	CreatedAt int                   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int                   `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:;index" json:"deletedAt"`
}
