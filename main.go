package main

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wikinewsfeed/wikinewsfeed/web"
)

func main() {
	mux := mux.NewRouter()
	mux.Use(web.CacheHeaders)
	mux.Use(web.CorsHeaders)
	mux.Use(web.FeedType)
	mux.Use(web.FeedAnalytics)
	mux.Use(web.EventContext)

	mux.HandleFunc("/api/events", web.Events)
	mux.HandleFunc("/feed/{type}", web.Feed)
	mux.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/feed.html", http.StatusMovedPermanently)
	})
	mux.Handle("/metrics", promhttp.Handler())
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs/.vuepress/dist")))
	http.ListenAndServe(fmt.Sprintf(":%s", envy.Get("PORT", "8080")), mux)
}
