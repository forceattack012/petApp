package domain

import "github.com/forceattack012/petAppApi/entities"

type Owner interface {
	Save(owner *[]entities.Owner) error
	Remove(id int) error
}
