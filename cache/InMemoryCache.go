package cache

import (
	"fmt"
)

type InMemoryCache[T any] struct {
	cache map[string]cachedItem[T]
	ttl   int64
}

func NewInMemoryCache[T any](ttl int64) *InMemoryCache[T] {
	if ttl <= 0 {
		ttl = 60
	}

	return &InMemoryCache[T]{
		cache: make(map[string]cachedItem[T]),
		ttl:   ttl,
	}
}

func (c *InMemoryCache[T]) Set(key string, item T) {
	if len(key) == 0 {
		return
	}

	newItem := *newCachedItem(item, c.ttl)
	c.cache[key] = newItem
}

func (c *InMemoryCache[T]) Get(key string) (T, bool) {
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
	if len(key) == 0 {
		return
	}

	delete(c.cache, key)
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
