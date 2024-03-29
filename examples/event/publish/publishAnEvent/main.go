package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/event"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Event.Publish(&event.PublishRequest{
		Topic:   "user",
		Message: &event.Json{},
	})
	fmt.Println(rsp, err)
}
