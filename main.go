package main

import (
	"net/http"

	"github.com/wikinewsfeed/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", web.Api)
	mux.HandleFunc("/feed.rss", web.Rss)
	mux.HandleFunc("/feed.atom", web.Rss)
	mux.HandleFunc("/feed.json", web.Rss)
	http.ListenAndServe(":3000", mux)
}
