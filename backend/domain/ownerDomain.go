package domain

import "github.com/forceattack012/petAppApi/entities"

type OwnerDomain interface {
	Save(owner *[]entities.Owner) error
	Remove(id int) error
}
