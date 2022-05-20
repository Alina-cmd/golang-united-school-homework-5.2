package cache

import "time"

type valueDeadline struct {
	value     string
	deadline  time.Time
	unlimited bool
}

type Cache struct {
	m map[string]valueDeadline
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	cache, ok := c.m[key]
	if !ok {
		return "", ok
	}

	if cache.unlimited {
		return cache.value, true
	} else {
		if cache.deadline.Before(time.Now()) {
			return cache.value, true
		}
	}

	return "", false
}

func (c Cache) Put(key, value string) {
	c.m[key] = valueDeadline{value, time.Now(), true}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.m))
	for key, value := range c.m {
		if value.unlimited {
			keys = append(keys, key)
		} else if value.deadline.Before(time.Now()) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = valueDeadline{value, deadline, false}
}
