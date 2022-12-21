package store

import (
	"github.com/forceattack012/petAppApi/entities"
	"gorm.io/gorm"
)

type OwnerStore struct {
	*gorm.DB
}

func NewOwnerStore(c *gorm.DB) *OwnerStore {
	return &OwnerStore{c}
}

func (g *OwnerStore) Save(owner *[]entities.Owner) error {
	return g.Create(owner).Error
}
func (g *OwnerStore) Remove(id int) error {
	return g.Delete(&entities.Owner{}, id).Error
}
