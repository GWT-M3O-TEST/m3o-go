package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/minecraft"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Minecraft.Ping(&minecraft.PingRequest{
		Address: "funcraft.net",
	})
	fmt.Println(rsp, err)
}
