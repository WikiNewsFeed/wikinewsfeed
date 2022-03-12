package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

type WikiResponseError struct {
	Code string
	Info string
}

type WikiResponse struct {
	Error WikiResponseError
	Parse struct {
		Text map[string]interface{}
	}
}

func EventContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/feed/") {
			var page = ""
			if r.URL.Query().Has("page") {
				page = "/" + r.URL.Query().Get("page")
			}

			apiUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&format=json&smaxage=1800&page=Portal:Current_events%s&prop=text", page)
			apiResponse, err := http.Get(apiUrl)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var content WikiResponse
			decodeError := json.NewDecoder(apiResponse.Body).Decode(&content)
			if decodeError != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Check if Wikipedia API's response error isn't empty
			var emptyError = WikiResponseError{}
			if content.Error != emptyError {
				if content.Error.Code == "missingtitle" {
					http.Error(w, content.Error.Info, http.StatusNotFound)
				} else {
					http.Error(w, content.Error.Info, http.StatusBadRequest)
				}

				return
			}

			parsedContent := strings.NewReader(content.Parse.Text["*"].(string))
			events, err := parser.Parse(parsedContent, false)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			eventsContext := context.WithValue(r.Context(), "Events", events)
			next.ServeHTTP(w, r.WithContext(eventsContext))
		} else {
			next.ServeHTTP(w, r)
		}
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
