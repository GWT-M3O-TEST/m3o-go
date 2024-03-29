package news

import (
	"go.m3o.com/client"
)

type News interface {
	Headlines(*HeadlinesRequest) (*HeadlinesResponse, error)
}

func NewNewsService(token string) *NewsService {
	return &NewsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type NewsService struct {
	client *client.Client
}

// Get the latest news headlines
func (t *NewsService) Headlines(request *HeadlinesRequest) (*HeadlinesResponse, error) {

	rsp := &HeadlinesResponse{}
	return rsp, t.client.Call("news", "Headlines", request, rsp)

}

type Article struct {
	// article title
	Title string `json:"title,omitempty"`
	// article id
	Id string `json:"id,omitempty"`
	// image url
	ImageUrl string `json:"image_url,omitempty"`
	// the article language
	Language string `json:"language,omitempty"`
	// the locale
	Locale string `json:"locale,omitempty"`
	// time it was published
	PublishedAt string `json:"published_at,omitempty"`
	// first 60 characters of article body
	Snippet string `json:"snippet,omitempty"`
	// source of news
	Source string `json:"source,omitempty"`
	// url of the article
	Url string `json:"url,omitempty"`
	// categories
	Categories []string `json:"categories,omitempty"`
	// article description
	Description string `json:"description,omitempty"`
	// related keywords
	Keywords string `json:"keywords,omitempty"`
}

type HeadlinesRequest struct {
	// comma separated list of languages to retrieve in e.g en,es
	Language string `json:"language,omitempty"`
	// comma separated list of countries to include e.g us,ca
	Locale string `json:"locale,omitempty"`
	// date published on in YYYY-MM-DD format
	Date string `json:"date,omitempty"`
}

type HeadlinesResponse struct {
	Articles []Article `json:"articles,omitempty"`
}
