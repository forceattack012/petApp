package entities

import "time"

type Owner struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64 `json:"user_id"`
	PetID     int64 `json:"pet_id"`
}
