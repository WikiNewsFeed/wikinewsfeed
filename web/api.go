package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

func Api(w http.ResponseWriter, r *http.Request) {
	events := r.Context().Value("Events").([]parser.Event)
	parsed, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(events[0])

	w.Header().Add("Content-Type", "application/json")
	w.Write(parsed)
}
