package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/movie"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Movie.Search(&movie.SearchRequest{
		Page:               1,
		Region:             "US",
		Year:               2010,
		PrimaryReleaseYear: 2010,
		Language:           "en-US",
		Query:              "inception",
	})
	fmt.Println(rsp, err)
}
