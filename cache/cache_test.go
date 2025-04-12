package cache

import (
	"strconv"
	"sync"
	"testing"
)

// TestSetAndGet проверяет корректность установки и получения значений.
func TestSetAndGet(t *testing.T) {
	cache := NewCache()

	// Устанавливаем значения
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Проверяем, что по ключам возвращаются ожидаемые значения.
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected 'value1' for key1, got '%s'", val)
	}
	if val, ok := cache.Get("key2"); !ok || val != "value2" {
		t.Errorf("Expected 'value2' for key2, got '%s'", val)
	}
	if _, ok := cache.Get("nonexistent"); ok {
		t.Error("Expected false for nonexistent key, got true")
	}
}

// TestConcurrentAccess проверяет потокобезопасность кеша при конкурентном доступе.
func TestConcurrentAccess(t *testing.T) {
	cache := NewCache()
	var wg sync.WaitGroup

	// Параллельно записываем и читаем данные
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := "key" + strconv.Itoa(n)
			cache.Set(key, "value"+strconv.Itoa(n))
			if val, ok := cache.Get(key); !ok || val != "value"+strconv.Itoa(n) {
				t.Errorf("Expected value for %s, got %s", key, val)
			}
		}(i)
	}

	wg.Wait()
}
