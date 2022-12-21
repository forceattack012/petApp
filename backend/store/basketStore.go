package store

import (
	"github.com/ReneKroon/ttlcache"
	"github.com/forceattack012/petAppApi/entities"
)

type BasketStore struct {
	Cache *ttlcache.Cache
}

func NewBasketStore(g *ttlcache.Cache) *BasketStore {
	return &BasketStore{g}
}

func (s *BasketStore) Add(key string, e *entities.Basket) {
	s.Cache.Set(key, e)
}

func (s *BasketStore) Get(key string) (*entities.Basket, bool) {
	if basket, found := s.Cache.Get(key); found {
		return basket.(*entities.Basket), found
	}

	return nil, false
}
