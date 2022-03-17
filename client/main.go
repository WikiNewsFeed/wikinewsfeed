package client

import (
	"strings"
	"time"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

type WikiClientOptions struct {
	MaxAge          time.Duration
	IncludeOriginal bool
}

type SubscribeFunc func(parser.Event)

func Get(page string, options WikiClientOptions) ([]parser.Event, error) {
	wikiPage, err := GetEventsPage(page, WikiRequestOptions{
		MaxAge: options.MaxAge,
	})
	if err != nil {
		return nil, err
	}

	parsedContent := strings.NewReader(wikiPage.Parse.Text["*"].(string))
	events, err := parser.Parse(parsedContent, parser.ParserOptions{
		IncludeOriginal: options.IncludeOriginal,
	})
	if err != nil {
		return nil, err
	}

	return events, nil
}

func Subscribe(call SubscribeFunc, frequency time.Duration) error {
	lastHashes := make(map[string]parser.Event)
	for {
		events, err := Get("", WikiClientOptions{
			MaxAge:          frequency,
			IncludeOriginal: false,
		})

		// Fill the hash map when just subscribed
		if len(lastHashes) == 0 {
			for _, event := range events {
				lastHashes[event.Checksum] = event
			}
		}

		for _, event := range events {
			_, present := lastHashes[event.Checksum]
			if !present {
				lastHashes[event.Checksum] = event
				call(event)
			}
		}

		if err != nil {
			return err
		}

		time.Sleep(frequency)
	}
}
