package main
import (
	
	"sync"

)

type Cache struct {
	objects map[string]CacheObject
	mutex   *sync.RWMutex
}  


type CacheObject struct {
	Value      string
	TimeToLive int64
}


