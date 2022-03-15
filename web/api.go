package web

import (
	"encoding/json"
	"net/http"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
	"github.com/wikinewsfeed/wikinewsfeed/stats"
)

func Events(w http.ResponseWriter, r *http.Request) {
	events := r.Context().Value("Events").([]parser.Event)
	parsed, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(parsed)
}

func Stats(w http.ResponseWriter, r *http.Request) {
	feedStats, err := stats.GetStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parsed, err := json.Marshal(feedStats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(parsed)
}
