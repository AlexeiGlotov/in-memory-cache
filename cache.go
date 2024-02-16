package in_memory_cache_

type stvalue struct {
	value string
}

type MyCache struct {
	cache map[string]stvalue
}

func NewCache() *MyCache {
	return &MyCache{
		cache: make(map[string]stvalue),
	}
}

func (m MyCache) Set(key string, value stvalue) {
	m.cache[key] = value
}

func (m MyCache) Get(key string) (res stvalue, value string) {
	val, ok := m.cache[key]
	if ok {
		return val, val.value
	}

	return
}

func (m MyCache) Delete(key string) {
	_, ok := m.cache[key]
	if ok {
		delete(m.cache, key)
	}
}

func (m MyCache) Info() (res map[string]stvalue) {
	return m.cache
}
