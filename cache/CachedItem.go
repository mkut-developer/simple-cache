package cache

import "time"

type cachedItem[T any] struct {
	time int64
	item T
}

func newCachedItem[T any](item T, ttl int64) *cachedItem[T] {
	return &cachedItem[T]{
		time: time.Now().Unix() + ttl,
		item: item,
	}
}

func (c *cachedItem[T]) isExpired() bool {
	return time.Now().Unix() > c.time
}

func (c *cachedItem[T]) getItem() T {
	return c.item
}
