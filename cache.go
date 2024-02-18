package in_memory_cache

import (
	"sync"
	"time"
)

type ElemValue struct {
	value      interface{}
	ttl        time.Duration
	timeCreate time.Time
}

type MyCache struct {
	cache map[string]ElemValue
	mu    *sync.Mutex
}

func NewCache() *MyCache {
	return &MyCache{
		cache: make(map[string]ElemValue),
		mu:    new(sync.Mutex),
	}
}

func (m *MyCache) Set(key string, value interface{}, ttl time.Duration) {
	m.mu.Lock()
	m.cache[key] = ElemValue{value: value, ttl: ttl, timeCreate: time.Now()}
	m.mu.Unlock()
}

func (m MyCache) Get(key string) (res ElemValue, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.cache[key]
	if ok {
		if time.Since(m.cache[key].timeCreate) > m.cache[key].ttl {
			m.Delete(key)
			return
		} else {
			return val, val.value
		}

	}

	return
}

func (m MyCache) Delete(key string) {
	m.mu.Lock()
	_, ok := m.cache[key]
	if ok {
		delete(m.cache, key)
	}
	m.mu.Unlock()
}

func (m MyCache) Info() map[string]ElemValue {
	return m.cache
}
