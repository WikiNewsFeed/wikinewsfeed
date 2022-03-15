package web

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/gorilla/mux"
	"github.com/wikinewsfeed/wikinewsfeed/client"
	"github.com/wikinewsfeed/wikinewsfeed/parser"
	"github.com/wikinewsfeed/wikinewsfeed/stats"
)

func EventContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/events" || strings.HasPrefix(r.URL.Path, "/feed/") {
			var page = ""
			if r.URL.Query().Has("page") {
				page = "/" + r.URL.Query().Get("page")
			}

			var includeOriginal = false
			if r.URL.Query().Has("includeOriginal") {
				includeOriginal = true
			}

			wikiPage, err := client.GetEventsPage(page, envy.Get("WNF_MAXAGE", "1800"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			parsedContent := strings.NewReader(wikiPage.Parse.Text["*"].(string))
			events, err := parser.Parse(parsedContent, includeOriginal)
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
			w.Header().Add("Cache-Control", fmt.Sprintf("public, max-age=%s", envy.Get("WNF_MAXAGE", "1800")))
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

func FeedAnalytics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/feed/") {
			subscriber := r.URL.Query().Get("subscribe")
			go stats.SubscribeIfNotAlready(subscriber)
			// go stats.IncrementHits(subscriber)
		}

		next.ServeHTTP(w, r)
	})
}
