package cache

import (
	"errors"
	"l0_wb/internal/pkg/domain"
)

type Cache struct {
	data map[string]domain.Order
}

func New() *Cache {
	cache := &Cache{data: make(map[string]domain.Order)}
	return cache
}

func (c *Cache) Get(index string) (domain.Order, error) {
	val, ok := c.data[index]
	if ok {
		return val, nil

	} else {
		return val, errors.New("Failed to get from cache")
	}
}

func (c *Cache) Update(index string, value domain.Order) {
	c.data[index] = value
}
