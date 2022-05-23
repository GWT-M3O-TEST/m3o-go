package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cache"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cache.Decrement(&cache.DecrementRequest{
		Value: 2,
		Key:   "counter",
	})
	fmt.Println(rsp, err)
}
