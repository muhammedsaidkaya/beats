package add_docker_metadata

import (
 "github.com/patrickmn/go-cache"
 "time"
)

var customCache = NewCache(time.Duration(10)*time.Minute, time.Duration(10)*time.Minute)

type Cache struct {
 inMemory *cache.Cache
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
 return &Cache{inMemory: cache.New(defaultExpiration, cleanupInterval)}
}

func (_cache Cache) SetCache(key string, value interface{}) {
 _cache.inMemory.Set(key, value, cache.DefaultExpiration)
}

func (_cache Cache) GetCache(key string) (interface{}, bool) {
 var data interface{}
 var found bool
 data, found = _cache.inMemory.Get(key)
 if found {
  return data, found
 }
 return nil, found
}
