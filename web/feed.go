package web

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gorilla/feeds"
	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

func Feed(w http.ResponseWriter, r *http.Request) {
	feedType := r.Context().Value("FeedType").(string)
	events := r.Context().Value("Events").([]parser.Event)
	feed := &feeds.Feed{
		Title:       "WikiNewsFeed",
		Description: "News aggregator powered by Wikipedia",
		Link: &feeds.Link{
			Href: envy.Get("WNF_URL", "http://localhost:8080"),
		},
		Author:    &feeds.Author{Name: "Wikipedia contributors"},
		Copyright: "Creative Commons Attribution-ShareAlike License 3.0",
		Image: &feeds.Image{
			Url:   fmt.Sprintf("%s/favicon.ico", envy.Get("WNF_URL", "http://localhost:8080")),
			Title: "WikiNewsFeed Logo",
		},
	}

	for _, event := range events {
		feed.Add(&feeds.Item{
			Id:          event.Checksum,
			Title:       event.PrimaryTopic.Title,
			Link:        &feeds.Link{Href: event.PrimaryTopic.ExternalUrl},
			Source:      &feeds.Link{Href: event.PrimarySource.Url},
			Description: event.Text,
			Content:     event.Html,
			Created:     event.Date,
			Author:      feed.Author,
		})
	}

	var generated string
	var feedError error

	switch feedType {
	case "atom":
		generated, feedError = feed.ToAtom()
		w.Header().Add("Content-Type", "application/atom+xml")
	case "rss":
		generated, feedError = feed.ToRss()
		w.Header().Add("Content-Type", "application/rss+xml")
	case "json":
		generated, feedError = feed.ToJSON()
		w.Header().Add("Content-Type", "application/feed+json")
	}

	if feedError != nil {
		http.Error(w, feedError.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(generated))
}
