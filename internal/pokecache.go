package internal

import (
	"time"
	"sync"
)

type Cache struct {
	mu sync.Mutex
	cacheItems map[string]cacheEntry
}

type cacheEntry struct {
	createdAt  time.Time
	val        []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheItems[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.cacheItems[key]
	if !ok {
		return []byte{}, false
	}
	return	item.val, true
}

func (c *Cache) reapLoop(interval time.Duration){
	//Uses the interval (time.Duration value passed to New Cache function) has passed, remove any entries in cache older than the interval
	//We can use time.Ticker to make this happen
	ticker := time.NewTicker(interval*time.Second)
	defer ticker.Stop()

	//Once we have a value
	for t:= range ticker.C {
		for key, entry := range c.cacheItems{
			delTime := entry.createdAt.Add(interval)
			if delTime.Compare(t) == -1 {
				c.mu.Lock()
				delete(c.cacheItems, key)
				c.mu.Unlock()
			}
		}

	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheItems: make(map[string]cacheEntry),
		mu:         sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}