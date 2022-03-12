package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/wikinewsfeed/wikinewsfeed/client"
	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

func EventContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/feed/") {
			var page = ""
			if r.URL.Query().Has("page") {
				page = "/" + r.URL.Query().Get("page")
			}

			wikiPage, err := client.GetEventsPage(page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			parsedContent := strings.NewReader(wikiPage.Parse.Text["*"].(string))
			events, err := parser.Parse(parsedContent, false)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			eventsContext := context.WithValue(r.Context(), "Events", events)
			next.ServeHTTP(w, r.WithContext(eventsContext))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CacheHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/feed/") {
			maxAge, exists := os.LookupEnv("MAXAGE")
			if !exists {
				maxAge = "1800"
			}

			w.Header().Add("Cache-Control", fmt.Sprintf("public, max-age=%s", maxAge))
		}

		next.ServeHTTP(w, r)
	})
}

func FeedType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/feed/") {
			feedType := mux.Vars(r)["type"]
			if feedType != "atom" && feedType != "rss" && feedType != "json" {
				http.Error(w, "No such feed type", http.StatusNotFound)
				return
			}

			feedContext := context.WithValue(r.Context(), "FeedType", feedType)
			next.ServeHTTP(w, r.WithContext(feedContext))
			return
		}

		next.ServeHTTP(w, r)
	})
}
