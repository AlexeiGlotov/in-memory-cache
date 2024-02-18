package in_memory_cache_

import (
	"sync"
	"time"
)

type ElemValue struct {
	value      string
	ttl        time.Duration
	timeCreate time.Time
}

type MyCache struct {
	cache map[string]ElemValue
	mu    sync.Mutex
}

func NewCache() *MyCache {
	return &MyCache{
		cache: make(map[string]ElemValue),
	}
}

func (m MyCache) Set(key string, value string, ttl time.Duration) {
	m.cache[key] = ElemValue{value: value, ttl: ttl, timeCreate: time.Now()}
}

func (m MyCache) Get(key string) (res ElemValue, value string) {

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
	_, ok := m.cache[key]
	if ok {
		delete(m.cache, key)
	}
}

func (m MyCache) Info() (res map[string]ElemValue) {
	return m.cache
}
