package cache

import "time"

type valueItem struct {
	value string
	date  time.Time
}

type Cache struct {
	data map[string]valueItem
}

func NewCache() Cache {
	return Cache{make(map[string]valueItem)}
}

func (c Cache) Get(key string) (string, bool) {
	valueItem, ok := c.data[key]
	if ok {
		return valueItem.value, true
	}
	c.cleanExpired()
	return "", false
}

func (c Cache) Put(key, value string) {
	c.PutTill(key, value, time.Time{})
}

func (c Cache) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range c.data {
		keys = append(keys, k)
	}
	c.cleanExpired()
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = valueItem{value: value, date: deadline}
	c.cleanExpired()
}

func (c Cache) cleanExpired() {
	now := time.Now()
	for key, valueItem := range c.data {
		if !valueItem.date.IsZero() && valueItem.date.Before(now) {
			delete(c.data, key)
		}
	}
}
