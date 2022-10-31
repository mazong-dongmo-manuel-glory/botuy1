package newsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	SourceId    string   `json:"source_id"`
	Keywords    string   `json:"keywords"`
	Creator     string   `json:"creator"`
	ImageUrl    string   `json:"image_url"`
	VideoUrl    string   `json:"video_url"`
	Description string   `json:"description"`
	PubDate     string   `json:"pubdate"`
	County      []string `json:"country"`
	Category    []string `json:"category"`
	Language    string   `json:"language"`
}
type Response struct {
	Status       string   `json:"status"`
	TotalResults int      `json:"totalResults"`
	Results      []Result `json:"results"`
}

func FetchNews(lang string) *Response {
	url := fmt.Sprintf("https://newsdata.io/api/1/news?apikey=pub_1284394204d1d7a6770f1ba935341455b7c10&language=%v&q=tech", lang)
	content, err := http.Get(url)
	if err != nil {
		return nil
	}
	decoder := json.NewDecoder(content.Body)
	response := &Response{}
	decoder.Decode(response)

	return response
}
