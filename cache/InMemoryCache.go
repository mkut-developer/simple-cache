package cache

import (
	"fmt"
	"sync"
	"time"
)

type InMemoryCache[T any] struct {
	cache map[string]cachedItem[T]
	ttl   int64
	mutex sync.RWMutex
}

func NewInMemoryCache[T any](ttl int64) *InMemoryCache[T] {
	if ttl <= 0 {
		ttl = 60
	}

	cache := InMemoryCache[T]{
		cache: make(map[string]cachedItem[T]),
		ttl:   ttl,
		mutex: sync.RWMutex{},
	}

	go cache.startEviction()

	return &cache
}

func (c *InMemoryCache[T]) Set(key string, item T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(key) == 0 {
		return
	}

	newItem := *newCachedItem(item, c.ttl)
	c.cache[key] = newItem
}

func (c *InMemoryCache[T]) Get(key string) (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var zero T
	if len(key) == 0 {
		return zero, false
	}

	item, ok := c.cache[key]
	if !ok || item.isExpired() {
		return zero, false
	}

	return item.getItem(), true
}

func (c *InMemoryCache[T]) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(key) == 0 {
		return
	}

	delete(c.cache, key)
}

func (c *InMemoryCache[T]) startEviction() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, item := range c.cache {
			if item.isExpired() {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}

func (c *InMemoryCache[T]) Print() {
	fmt.Printf("ttl: %d", c.ttl)
	fmt.Println()
	fmt.Println("cache:")
	for key, value := range c.cache {
		fmt.Println(key, value.getItem())
	}
	fmt.Println()
}

func (c *InMemoryCache[T]) Size() int {
	return len(c.cache)
}
