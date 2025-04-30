package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mut      sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mut.Lock()
	c.entries[key] = entry
	c.mut.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	entry, ok := c.entries[key]
	c.mut.Unlock()

	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case _ = <-ticker.C:
			for k, v := range c.entries {
				if time.Now().Sub(v.createdAt) > c.interval {
					delete(c.entries, k)
				}
			}
		}
	}
}
