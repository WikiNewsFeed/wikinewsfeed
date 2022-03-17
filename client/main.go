package client

import (
	"strings"
	"time"

	"github.com/wikinewsfeed/wikinewsfeed/parser"
)

// type WikiSubscribeOptions struct {
// 	Period string
// 	Call   SubscribeFunc
// }

type WikiClientOptions struct {
	MaxAge          time.Duration
	IncludeOriginal bool
}

type SubscribeFunc func([]parser.Event)
type SubscribeEachFunc func(parser.Event)

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

func Subscribe(page string, call SubscribeFunc, frequency time.Duration) error {
	var lastDelta = 0
	for {
		freshEvents, err := Get(page, WikiClientOptions{
			MaxAge:          frequency,
			IncludeOriginal: false,
		})
		if err != nil {
			return err
		}

		// Reset delta if after initial subscription
		if lastDelta == 0 {
			lastDelta = len(freshEvents)
		}

		if len(freshEvents) > lastDelta {
			delta := len(freshEvents) - lastDelta
			call(freshEvents[:delta])
		}

		lastDelta = len(freshEvents)
		time.Sleep(frequency)
	}
}

func SubscribeEach(page string, call SubscribeEachFunc, frequency time.Duration) error {
	return Subscribe(page, func(events []parser.Event) {
		for _, event := range events {
			call(event)
		}
	}, frequency)
}
