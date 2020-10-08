package main

import (
	"fmt"
	"github.com/bytepowered/cache"
	"time"
)

func main() {
	gc := cache.New(10).
		LFU().
		LoaderFunc(func(key interface{}) (interface{}, error) {
			return fmt.Sprintf("%v-value", key), nil
		}).
		Build()

	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	v2, err := gc.GetOrLoad("key2", func(key interface{}) (interface{}, *time.Duration, error) {
		return fmt.Sprintf("%v-ol-value", key), cache.NoExpiration, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(v2)
}
