package in_memory_cache_

type MyCache struct {
	cache map[string]string
}

func NewCache() *MyCache {
	return &MyCache{
		cache: make(map[string]string),
	}
}

func (m MyCache) Set(key string, value string) {
	m.cache[key] = value
}

func (m MyCache) Get(key string) (res string) {
	val, ok := m.cache[key]
	if ok {
		return val
	}

	return
}

func (m MyCache) Delete(key string) {
	_, ok := m.cache[key]
	if ok {
		delete(m.cache, key)
	}
}

func (m MyCache) Info() (res map[string]string) {
	return m.cache
}
