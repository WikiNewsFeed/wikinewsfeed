package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wikinewsfeed/wikinewsfeed/web"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api", web.Api)
	mux.HandleFunc("/feed.rss", web.Rss)
	mux.HandleFunc("/feed.atom", web.Rss)
	mux.HandleFunc("/feed.json", web.Rss)
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs/.vuepress/dist")))
	http.ListenAndServe(":8080", mux)
}
