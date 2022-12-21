package entities

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Owners    []Owner `json:"owners" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
