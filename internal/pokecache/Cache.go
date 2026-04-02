package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	var newCache = Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	_, exists := c.cacheMap[key]
	if exists == false {
		c.cacheMap[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.cacheMap[key]
	c.mu.Unlock()
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	newTicker := time.NewTicker(time.Second)
	for range newTicker.C {
		c.mu.Lock()
		for key, value := range c.cacheMap {
			timeDifference := time.Since(value.createdAt)
			if timeDifference > interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
