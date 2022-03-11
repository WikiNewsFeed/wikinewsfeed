package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wikinewsfeed/wikinewsfeed/web"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/events", web.Api)
	mux.HandleFunc("/feed/{feed}", web.Rss)
	mux.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/feed.html", http.StatusMovedPermanently)
	})
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs/.vuepress/dist")))
	http.ListenAndServe(":8080", mux)
}
