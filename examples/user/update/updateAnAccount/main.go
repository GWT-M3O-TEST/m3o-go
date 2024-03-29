package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.Update(&user.UpdateRequest{
		Id:       "user-1",
		Email:    "joe+2@example.com",
		Username: "joe",
	})
	fmt.Println(rsp, err)
}
