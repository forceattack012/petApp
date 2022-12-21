package domain

import "github.com/forceattack012/petAppApi/entities"

type BasketDomain interface {
	Add(key string, e *entities.Basket)
	Get(key string) (*entities.Basket, bool)
}
