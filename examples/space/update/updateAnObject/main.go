package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/space"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Space.Update(&space.UpdateRequest{
		Name:       "images/file.jpg",
		Visibility: "public",
		Object:     "<file bytes>",
	})
	fmt.Println(rsp, err)
}
