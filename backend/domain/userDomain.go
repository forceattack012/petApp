package domain

import "github.com/forceattack012/petAppApi/entities"

type UserDomain interface {
	Create(user *entities.User) error
	GetUser(username string, password string) (*entities.User, error)
}
