package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu    *sync.RWMutex
	store map[string]interface{}
}

func New() *Cache {
	return &Cache{
		mu:    &sync.RWMutex{},
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, exists := c.store[key]; exists {
		return value
	} else {
		fmt.Println("Not found!")
		return nil
	}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}
