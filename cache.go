package luoy

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	cache := &Cache{
		data: make(map[string]interface{}),
	}

	return cache
}

func (cache *Cache) Get(key string) (interface{}, bool) {
	value, ok := cache.data[key]
	return value, ok
}

func (cache *Cache) Insert(key string, value interface{}) bool {
	cache.data[key] = value
	return true
}

func (cache *Cache) Delete(key string) {
	delete(cache.data, key)
}
