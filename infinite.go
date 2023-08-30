package client

import (
	"fmt"

	"github.com/infinite-industries/infinite-go/formats/infinite"
)

const (
	libraryVersion   = "0.1"
	defaultBaseURL   = "https://api.infinite.industries/"
	defaultUserAgent = "infinite-go/" + libraryVersion
)

type EventResponse struct {
	Event  infinite.Event `json:"event"`
	Status string         `json:"status"`
}

type EventsResponse struct {
	Events []infinite.Event `json:"events"`
	Status string           `json:"status"`
}

type VenueResponse struct {
	Venue  infinite.Venue `json:"venue"`
	Status string         `json:"status"`
}

type VenuesResponse struct {
	Venues []infinite.Venue `json:"venues"`
	Status string           `json:"status"`
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorStr   string `json:"error"`
	//	Message    string `json:"message"` // the api sometimes returns an array for this?
}

func (e *ErrorResponse) Error() string {
	if e.StatusCode != 0 || e.ErrorStr != "" {
		return fmt.Sprintf("service error: %d; %s;\n", e.StatusCode, e.ErrorStr)
	}
	return ""
}
