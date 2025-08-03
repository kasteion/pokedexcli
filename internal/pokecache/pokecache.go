package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu *sync.Mutex
}

// var cache *Cache

func NewCache(interval time.Duration) Cache {
	// if cache == nil {
	// 	cache = &Cache{ 
	// 		entries: map[string]cacheEntry{}, 
	// 		mu: &sync.Mutex{}, 
	// 	}
	// 	go cache.reapLoop(interval)
	// }
	
	// return cache

	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{ 
		createdAt: time.Now().UTC(), 
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	// if ok {
	// 	return entry.val, ok
	// }

	// return nil, false
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
	

	// for {
	// 	t := <-ticker.C;
	// 	fmt.Println("Tick at", t)
	// }

	// go func() {
	// 	for {
	// 		t := <-ticker.C;
	// 		for k, v := range(c.entries) {
	// 			if t.Sub(v.createdAt) >= c.interval {
	// 				c.mu.Lock()
	// 				delete(c.entries, k)
	// 				c.mu.Unlock()
	// 			}
	// 		}
	// 	}
	// }()
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}