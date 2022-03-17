package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type WikiResponse struct {
	Error WikiResponseError
	Parse struct {
		Text map[string]interface{}
	}
}

type WikiResponseError struct {
	Code string
	Info string
}

type WikiRequestOptions struct {
	MaxAge          time.Duration
	IncludeOriginal bool
}

var requestUrl = "https://en.wikipedia.org/w/api.php?action=parse&format=json&smaxage=%v&page=Portal:Current_events%s&prop=text"

func GetEventsPage(page string, options WikiRequestOptions) (*WikiResponse, error) {
	apiUrl := fmt.Sprintf(requestUrl, options.MaxAge.Seconds(), page)
	apiResponse, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	defer apiResponse.Body.Close()

	var wikiPage WikiResponse
	decodeError := json.NewDecoder(apiResponse.Body).Decode(&wikiPage)
	if decodeError != nil {
		return nil, decodeError
	}

	// Check if Wikipedia API's response error isn't empty
	var emptyError = WikiResponseError{}
	if wikiPage.Error != emptyError {
		return nil, errors.New(wikiPage.Error.Info)
	}

	return &wikiPage, nil
}
