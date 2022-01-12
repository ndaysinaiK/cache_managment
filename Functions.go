
package main 

import (
	
	"time"
	"sync"
	
)  

func (cacheObject CacheObject) IfExpired() bool {
	if cacheObject.TimeToLive == 0 {
		return false
	}
	return time.Now().UnixNano() > cacheObject.TimeToLive
}



func NewCache() *Cache {
	return &Cache{
		objects: make(map[string]CacheObject),
		mutex:   &sync.RWMutex{},
	}
}


func (cache Cache) GetObject(cacheKey string) string {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	var object CacheObject
	object = cache.objects[cacheKey]

	if object.IfExpired() {
		delete(cache.objects, cacheKey)
		return ""
	}
	return object.Value
}

func (cache Cache) SetValue(cacheKey string, cacheValue string, timeToLive time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.objects[cacheKey] = CacheObject{
		Value:      cacheValue,
		TimeToLive: time.Now().Add(timeToLive).UnixNano(),
	}
}