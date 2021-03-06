package parser

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Event struct {
	Html          string        `json:"html"`
	OriginalHtml  string        `json:"html_original,omitempty"`
	Text          string        `json:"text"`
	OriginalText  string        `json:"text_original,omitempty"`
	Category      string        `json:"category"`
	Topics        []EventPage   `json:"topics"`
	PrimaryTopic  EventPage     `json:"primary_topic"`
	Sources       []EventSource `json:"sources"`
	PrimarySource EventSource   `json:"primary_source"`
	References    []EventPage   `json:"references"`
	Date          time.Time     `json:"date"`
	OriginalDate  string        `json:"date_original,omitempty"`
	Checksum      string        `json:"checksum"`
	Page          string        `json:"page"`
	Contributors  string        `json:"contributors"`
}

type EventPage struct {
	Title       string `json:"title"`
	Uri         string `json:"uri"`
	ExternalUrl string `json:"external_url"`
}

type EventSource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ParserOptions struct {
	IncludeOriginal bool
}

func getPrimaryTopic(topics []EventPage) EventPage {
	if len(topics) > 0 {
		return topics[0]
	} else {
		return EventPage{}
	}
}

func getPrimarySource(sources []EventSource) EventSource {
	if len(sources) > 0 {
		return sources[0]
	} else {
		return EventSource{}
	}
}

func calculateChecksum(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func Parse(content io.Reader, options ParserOptions) ([]Event, error) {
	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		return nil, err
	}

	eventNodes := doc.Find(".vevent ul:not(.current-events-navbar) li:not(:has(ul))")
	var events = []Event{}
	for _, eventNode := range eventNodes.Nodes {
		event := goquery.NewDocumentFromNode(eventNode)
		html, err := event.Html()
		if err != nil {
			return nil, err
		}
		text := event.Text()
		page, _ := event.Parents().Find(".vevent").Attr("id")
		wikiPage := "https://en.wikipedia.org/wiki/Portal:Current_events/" + page
		contributors := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&titles=Portal:Current_events/%s&prop=contributors&format=json", page)

		var sources []EventSource
		sourcesNodes := event.Find("a.external")
		for _, sourceNode := range sourcesNodes.Nodes {
			source := goquery.NewDocumentFromNode(sourceNode)
			sourceName := source.Text()

			// Escape source in parentheses (if matches regex)
			sourceNameRegex := regexp.MustCompile(`\((.*?)\)`).FindStringSubmatch(sourceName)
			if len(sourceNameRegex) > 0 {
				sourceName = sourceNameRegex[1]
			}

			sourceUrl, _ := source.Attr("href")

			sources = append(sources, EventSource{
				Name: sourceName,
				Url:  sourceUrl,
			})
		}

		var references []EventPage
		referencesNodes := event.Find("a:not(.external)")
		for _, referenceNode := range referencesNodes.Nodes {
			reference := goquery.NewDocumentFromNode(referenceNode)
			referenceTitle := reference.Text()
			referenceUri, _ := reference.Attr("href")

			references = append(references, EventPage{
				Title:       referenceTitle,
				Uri:         referenceUri,
				ExternalUrl: "https://en.wikipedia.org" + referenceUri,
			})
		}

		// Strip sources and replace internal links with external
		stripped := event.Find("a.external").Remove().End()
		htmlStripped, err := stripped.Html()
		if err != nil {
			return nil, err
		}
		htmlStripped = regexp.MustCompile(`href="(.*?)"`).ReplaceAllString(htmlStripped, `href="https://en.wikipedia.org$1"`)
		textStripped := stripped.Text()

		var topics []EventPage
		primaryTopic := event.Parent().Parent().Find("a")
		primaryTopicTitle, err := primaryTopic.Html()
		if err != nil {
			return nil, err
		}
		primaryTopicUri, _ := primaryTopic.Attr("href")

		topics = append(topics, EventPage{
			Title:       primaryTopicTitle,
			Uri:         primaryTopicUri,
			ExternalUrl: "https://en.wikipedia.org" + primaryTopicUri,
		})

		date, _ := event.Parents().Find(".bday.dtstart").Html()
		dateFormatted, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}

		var parsedEvent = Event{
			Html:          htmlStripped,
			Text:          textStripped,
			Topics:        topics,
			PrimaryTopic:  getPrimaryTopic(topics),
			Sources:       sources,
			PrimarySource: getPrimarySource(sources),
			References:    references,
			Date:          dateFormatted,
			Checksum:      calculateChecksum(textStripped),
			Page:          wikiPage,
			Contributors:  contributors,
		}

		// Include original content?
		if options.IncludeOriginal {
			parsedEvent.OriginalHtml = html
			parsedEvent.OriginalText = text
			parsedEvent.OriginalDate = date
		}

		events = append(events, parsedEvent)
	}

	return events, nil
}
