package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func GetEventsPage(title string, smaxage string) (*WikiResponse, error) {
	apiUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&format=json&smaxage=%s&page=Portal:Current_events%s&prop=text", smaxage, title)
	apiResponse, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	defer apiResponse.Body.Close()

	var page WikiResponse
	decodeError := json.NewDecoder(apiResponse.Body).Decode(&page)
	if decodeError != nil {
		return nil, decodeError
	}

	// Check if Wikipedia API's response error isn't empty
	var emptyError = WikiResponseError{}
	if page.Error != emptyError {
		return nil, errors.New(page.Error.Info)
	}

	return &page, nil
}
