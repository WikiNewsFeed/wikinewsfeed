package web

import (
	"encoding/json"
	"net/http"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

type EventsResponse struct {
	Events  []parser.Event `json:"events"`
	License string         `json:"license"`
}

func Events(w http.ResponseWriter, r *http.Request) {
	events := r.Context().Value("Events").([]parser.Event)
	response, err := json.Marshal(EventsResponse{
		Events:  events,
		License: "https://creativecommons.org/licenses/by-sa/3.0",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
