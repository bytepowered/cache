package main

import (
	"fmt"
	"github.com/bytepowered/cache"
)

func main() {
	gc := cache.New(10).
		LFU().
		Build()
	gc.Set("key", "ok")

	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", v)
}
