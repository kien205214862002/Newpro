package common

import (
	"time"

	"gorm.io/gorm"
)

type SQLModel struct {
	Id        int       `json:"-" gorm:"column:id"`
	FakeId    string    `json:"id" gorm:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	// gorm.DeletedAt: Dùng để soft delete
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
