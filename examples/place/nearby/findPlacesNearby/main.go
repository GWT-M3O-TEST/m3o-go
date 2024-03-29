package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/place"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Place.Nearby(&place.NearbyRequest{
		Location: "51.5074577,-0.1297515",
		Keyword:  "tesco",
		Type:     "store",
	})
	fmt.Println(rsp, err)
}
