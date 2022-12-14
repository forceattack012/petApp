package entities

import (
	"time"

	"github.com/forceattack012/petAppApi/file"
)

type Pet struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string      `json:"name" binding:"required"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Age         string      `json:"age"`
	Files       []file.File `json:"files" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Owners      []Owner     `json:"owners" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
