package cache

import (
	"errors"
	"proxy-server/pkg/models"
	"sync"
)

var (
	ErrNoRecord     = errors.New("models: no matching record found")
	ErrRecordExists = errors.New("models: such record already exists")
)

type Cache struct {
	store sync.Map
}

// Set a value in the cache
func (c *Cache) Insert(key string, val models.Response) error {
	if _, err := c.Get(val.ID); err == nil {
		return ErrRecordExists
	}

	c.store.Store(val.ID, val)
	return nil
}

// Get a value by key from the cache
func (c *Cache) Get(key string) (models.Response, error) {
	val, ok := c.store.Load(key)
	if !ok {
		return models.Response{}, ErrNoRecord
	}

	return val.(models.Response), nil
}
