package cache

import "time"

type Cache struct {
	data map[string][]string
}

func NewCache() Cache {
	return Cache{make(map[string][]string)}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.data[key]
	if ok {
		return value[0], true
	}
	return "", false
}

func (c Cache) Put(key, value string) {
	c.data[key] = []string{value, ""}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	//	function empty
}
