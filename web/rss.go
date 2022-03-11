package web

import (
	"log"
	"net/http"

	"github.com/gorilla/feeds"
	"github.com/wikinewsfeed/parser"
)

func Rss(res http.ResponseWriter, req *http.Request) {
	wiki, _ := http.Get("https://en.wikipedia.org/wiki/Portal:Current_events")
	events, _ := parser.Parse(wiki.Body)

	feed := &feeds.Feed{
		Title: "WikiNewsFeed",
		Link:  &feeds.Link{Href: "http://localhost:8080/feed.json"},
		Image: &feeds.Image{
			Url:   "https://upload.wikimedia.org/wikipedia/commons/7/77/Wikipedia_svg_logo.svg",
			Title: "Wikipedia Logo",
			Link:  "https://upload.wikimedia.org/wikipedia/commons/7/77/Wikipedia_svg_logo.svg",
		},
	}

	for _, event := range events {
		feed.Add(&feeds.Item{
			Title:       event.PrimaryTopic.Title,
			Link:        &feeds.Link{Href: event.PrimaryTopic.ExternalUrl},
			Source:      &feeds.Link{Href: event.PrimarySource.Url},
			Description: event.Text,
			Content:     event.Html,
			Created:     event.Date,
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	res.Header().Add("Cache-Control", "public, max-age=1800")
	res.Write([]byte(atom))
}
