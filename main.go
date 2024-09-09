package main

import (
	"fmt"
	"github.com/mkut-developer/simple-cache/cache"
	"time"
)

func main() {
	const ttl = 5
	fmt.Println("-----------1-----------")
	c := cache.NewInMemoryCache[int](ttl)
	c.Set("key1", 1)
	c.Set("key2", 2)

	fmt.Println("-----------2-----------")
	fmt.Println(c.Size())

	fmt.Println("-----------3-----------")
	time.Sleep(time.Second * (ttl + 2))

	fmt.Println("-----------4-----------")
	fmt.Println("size", c.Size())
	fmt.Println(*c)

	fmt.Println("-----------5-----------")
}
