package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mmcdole/gofeed"
)

func main() {
	r := chi.NewRouter()

	r.Get("/api/news", func(w http.ResponseWriter, r *http.Request) {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("https://news.google.com/rss/search?q=finance")

		articles := []map[string]string{}
		for _, item := range feed.Items[:5] {
			articles = append(articles, map[string]string{
				"title":     item.Title,
				"link":      item.Link,
				"published": item.Published,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(articles)
	})

	http.ListenAndServe(":8080", r)
}
