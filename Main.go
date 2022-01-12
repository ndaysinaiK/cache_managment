
package main


import (
	"fmt"

)


func main() {

	var cache *Cache
	cache = NewCache()

	cache.SetValue("name", "eden meshack", 200000000)
	var name string
	name = cache.GetObject("name")
	fmt.Println(name)

}