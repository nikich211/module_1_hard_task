package cache

import "sync"

type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

// Get получает значение по ключу
func (c *Cache) Get(k string) (v string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok = c.data[k]
	return
}

// Пример использования
func main() {
	cache := NewCache()
	cache.Set("key1", "value1")
	if val, ok := cache.Get("key1"); ok {
		println("Found value:", val)
	} else {
		println("Value not found")
	}
}
