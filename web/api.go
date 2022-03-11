package web

import (
	"encoding/json"
	"net/http"

	"github.com/wikinewsfeed/parser"
)

func Api(res http.ResponseWriter, req *http.Request) {
	wiki, _ := http.Get("https://en.wikipedia.org/wiki/Portal:Current_events")
	events, _ := parser.Parse(wiki.Body)
	parsed, _ := json.Marshal(events)

	res.Header().Add("Content-Type", "application/json")
	res.Header().Add("Cache-Control", "public, max-age=1800")
	res.Write(parsed)
}
