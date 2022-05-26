package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/rss"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Rss.Add(&rss.AddRequest{
		Name:     "bbc",
		Url:      "http://feeds.bbci.co.uk/news/rss.xml",
		Category: "news",
	})
	fmt.Println(rsp, err)
}
