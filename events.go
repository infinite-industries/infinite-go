package infinite

import (
	"fmt"
)

// EventsService queries the events of the II API.
type EventsService service

type EventsParams struct {
	Tags []string `url:"tags,omitempty"`
}

// Get a specific event
func (v *EventsService) Get(id string) (Event, error) {

	eventResponse := new(EventResponse)
	errorResponse := new(ErrorResponse)

	path := fmt.Sprintf("%s", id)

	resp, err := v.sling.New().Get(path).Receive(eventResponse, errorResponse)

	return eventResponse.Event, newError(resp, err, errorResponse)

}

func (v *EventsService) CurrentVerified(tags ...string) ([]Event, error) {

	return v.list(true, true, tags...)

}

func (v *EventsService) list(current, verified bool, tags ...string) ([]Event, error) {

	eventsResponse := new(EventsResponse)
	errorResponse := new(ErrorResponse)

	path := ""

	if current == true {
		path = "current-verified"
	} else if verified == true {
		path = "verified"
	} else if verified == false {
		path = "non-verified"
	}

	params := &EventsParams{Tags: tags}
	resp, err := v.sling.New().Get(path).QueryStruct(params).Receive(eventsResponse, errorResponse)

	// fmt.Printf("request: %v\n", resp.Request)
	// fmt.Printf("response: %v\n", resp)

	return eventsResponse.Events, newError(resp, err, errorResponse)

}
