package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Logs(&function.LogsRequest{
		Name:     "helloworld",
		LogsType: "build",
	})
	fmt.Println(rsp, err)
}
