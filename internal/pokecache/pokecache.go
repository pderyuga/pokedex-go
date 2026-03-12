package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	items map[string]cacheEntry
	mu    sync.RWMutex // The RWMutex to protect the map
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		items: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)

	return c
}

// Set adds or updates an item in the map.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()         // Lock for exclusive write access
	defer c.mu.Unlock() // Ensure the lock is released
	c.items[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves an item from the map.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()         // RLock for concurrent read access
	defer c.mu.RUnlock() // Ensure the read lock is released
	item, ok := c.items[key]
	return item.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.items {
			if item.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
