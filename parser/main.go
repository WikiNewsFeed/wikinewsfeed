package parser

import (
	"io"
	"net/url"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type EventSource struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Domain string `json:"domain"`
}

type EventPage struct {
	Title       string `json:"title"`
	Uri         string `json:"uri"`
	ExternalUrl string `json:"external_url"`
}

type Event struct {
	Html          string        `json:"html"`
	HtmlOriginal  string        `json:"html_original,omitempty"`
	Text          string        `json:"text"`
	TextOriginal  string        `json:"text_original,omitempty"`
	Category      string        `json:"category"`
	Topics        []EventPage   `json:"topics"`
	PrimaryTopic  EventPage     `json:"primary_topic"`
	Sources       []EventSource `json:"sources"`
	PrimarySource EventSource   `json:"primary_source"`
	References    []EventPage   `json:"references"`
	Date          time.Time     `json:"date"`
	DateOriginal  string        `json:"date_original,omitempty"`
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

func Parse(content io.Reader, includeOriginal bool) ([]Event, error) {
	doc, _ := goquery.NewDocumentFromReader(content)

	const selector = ".vevent ul:not(.current-events-navbar) li:not(:has(ul))"
	var output = []Event{}

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		text := s.Text()

		var sources []EventSource
		s.Find("a.external").Each(func(i int, source *goquery.Selection) {
			sourceName := source.Text()

			// Escape source in parentheses
			sourceNameRegex := regexp.MustCompile(`\((.*?)\)`).FindStringSubmatch(sourceName)
			if len(sourceNameRegex) > 0 {
				sourceName = sourceNameRegex[1]
			}

			sourceUrl, _ := source.Attr("href")
			parsedUrl, _ := url.Parse(sourceUrl)

			sources = append(sources, EventSource{
				Name:   sourceName,
				Url:    sourceUrl,
				Domain: parsedUrl.Hostname(),
			})
		})

		var references []EventPage
		s.Find("a:not(.external)").Each(func(i int, reference *goquery.Selection) {
			referenceTitle := reference.Text()
			referenceUri, _ := reference.Attr("href")

			references = append(references, EventPage{
				Title:       referenceTitle,
				Uri:         referenceUri,
				ExternalUrl: "https://en.wikipedia.org" + referenceUri,
			})
		})

		// Strip sources and replace internal links with external
		stripped := s.Find("a.external").Remove().End()
		htmlStripped, _ := stripped.Html()
		htmlStripped = regexp.MustCompile(`href="(.*?)"`).ReplaceAllString(htmlStripped, `href="https://en.wikipedia.org$1"`)
		textStripped := stripped.Text()

		var topics []EventPage
		primaryTopic := s.Parent().Parent().Find("a")
		primaryTopicTitle, _ := primaryTopic.Html()
		primaryTopicUri, _ := primaryTopic.Attr("href")

		topics = append(topics, EventPage{
			Title:       primaryTopicTitle,
			Uri:         primaryTopicUri,
			ExternalUrl: "https://en.wikipedia.org" + primaryTopicUri,
		})

		date, _ := s.Parents().Find(".bday.dtstart").Html()
		dateFormatted, _ := time.Parse("2006-01-02", date)

		var event = Event{
			Html:          htmlStripped,
			Text:          textStripped,
			Topics:        topics,
			PrimaryTopic:  getPrimaryTopic(topics),
			Sources:       sources,
			PrimarySource: getPrimarySource(sources),
			References:    references,
			Date:          dateFormatted,
		}

		// Include original content?
		if includeOriginal {
			event.HtmlOriginal = html
			event.TextOriginal = text
			event.DateOriginal = date
		}

		output = append(output, event)
	})

	return output, nil
}
