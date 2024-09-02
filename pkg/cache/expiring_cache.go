package cache

import (
	"sync"
	"time"
)

// item represent a cache item with a value and an expiration time
type item[V any] struct {
	value  V
	expire time.Time
}

// isExpired checks if the cache has expired
func (i item[V]) isExpired() bool {
	return time.Now().After(i.expire)
}

type TTLCache[K comparable, V any] struct {
	items map[K]item[V] // the map storing cache items
	mu    sync.Mutex    // Mutex for controlling concurrent access to the cache
	rw    sync.RWMutex  // RWMutex is thus preferable for data that is mostly read
}

func NewTTL[K comparable, V any]() *TTLCache[K, V] {
	c := &TTLCache[K, V]{
		items: make(map[K]item[V]),
	}

	go func() {
		for range time.Tick(5 * time.Second) {
			c.mu.Lock()
			for key, item := range c.items {
				if item.isExpired() {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}()

	return c
}

// Set adds a new item to the cache with the specified key, value, and
// time to live (TTL)
func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock() // Lock the mutex to ensure safe access to the cache
	defer c.mu.Unlock()

	// Set the cache item with the specified key, value, and expiry time
	c.items[key] = item[V]{
		value:  value,
		expire: time.Now().Add(ttl),
	}
}

// Get retrieves the value associated with the given key from the cache
func (c *TTLCache[K, V]) Get(key K) (V, bool) {
	c.rw.Lock()
	defer c.rw.Unlock()

	item, found := c.items[key]
	if !found {
		// if the key is not found, return the zero value for V and false
		return item.value, false
	}

	if item.isExpired() {
		// If the item has expired, remove it from the cache and return the
		// value and false
		delete(c.items, key)
		return item.value, false
	}

	// Otherwise return and value and true
	return item.value, true
}

// Remove removes the item with the specified key from the cache
func (c *TTLCache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Delete the item with the given key from the cache
	delete(c.items, key)
}

// Pop removes and return the item with the specified key from the cache
func (c *TTLCache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		// If the key is not found, return the zero value for V and false
		return item.value, false
	}

	// If the key is found, delete the item from the cache
	delete(c.items, key)

	if item.isExpired() {
		//If the item has expired, return the value and false
		return item.value, false
	}

	//Otherwise return the value and true
	return item.value, true
}

func (c *TTLCache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[K]item[V])
}
