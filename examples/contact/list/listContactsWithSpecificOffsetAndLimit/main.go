package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/contact"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Contact.List(&contact.ListRequest{
		Offset: 1,
		Limit:  1,
	})
	fmt.Println(rsp, err)
}
