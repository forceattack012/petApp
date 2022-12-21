package store

import (
	"github.com/forceattack012/petAppApi/entities"
	"gorm.io/gorm"
)

type UserStore struct {
	*gorm.DB
}

func NewUserStore(g *gorm.DB) *UserStore {
	return &UserStore{g}
}

func (s *UserStore) Create(user *entities.User) error {
	return s.DB.Create(user).Error
}

func (s *UserStore) GetUser(username string, password string) (entities.User, error) {
	var user entities.User
	err := s.Where("username = ? AND password = ?", username, password).Find(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
