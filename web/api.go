package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

type WikiResponse struct {
	Parse struct {
		Text map[string]interface{}
	}
}

func Api(res http.ResponseWriter, req *http.Request) {
	page := ""
	if req.URL.Query().Has("page") {
		page = "/" + req.URL.Query().Get("page")
	}

	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&format=json&smaxage=1800&page=Portal:Current_events%s&prop=text", page)
	wiki, _ := http.Get(url)

	var content WikiResponse
	json.NewDecoder(wiki.Body).Decode(&content)

	r := strings.NewReader(content.Parse.Text["*"].(string))
	events, _ := parser.Parse(r, false)
	parsed, _ := json.Marshal(events)

	res.Header().Add("Content-Type", "application/json")
	res.Header().Add("Cache-Control", "public, max-age=1800")
	res.Write(parsed)
}
