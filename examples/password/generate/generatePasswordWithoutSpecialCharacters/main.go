package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/password"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Password.Generate(&password.GenerateRequest{
		Special:   false,
		Uppercase: true,
		Length:    16,
		Lowercase: true,
		Numbers:   true,
	})
	fmt.Println(rsp, err)
}
