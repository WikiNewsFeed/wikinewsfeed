package web

import (
	"net/http"

	"github.com/gorilla/feeds"
	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

func Feed(w http.ResponseWriter, r *http.Request) {
	feedType := r.Context().Value("FeedType").(string)
	events := r.Context().Value("Events").([]parser.Event)
	feed := &feeds.Feed{
		Title:       "WikiNewsFeed - Feed",
		Description: "News aggregator powered by Wikipedia",
		// Link: &feeds.Link{
		// 	Href: fmt.Sprintf("http://localhost:8080/feed/%s", feedType),
		// },
		Copyright: "Creative Commons Attribution-ShareAlike License 3.0",
		Image: &feeds.Image{
			Url:   "https://i.imgur.com/5TMSNk0.png",
			Title: "WikiNewsFeed Logo",
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

	var generated string
	var feedError error

	switch feedType {
	case "atom":
		generated, feedError = feed.ToAtom()
	case "rss":
		generated, feedError = feed.ToRss()
	case "json":
		generated, feedError = feed.ToJSON()
	}

	if feedError != nil {
		http.Error(w, feedError.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(generated))
}
