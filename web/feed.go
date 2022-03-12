package web

import (
	"net/http"

	"github.com/gorilla/feeds"
	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

func Feed(w http.ResponseWriter, r *http.Request) {
	events := r.Context().Value("Events").([]parser.Event)

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(atom))
}
