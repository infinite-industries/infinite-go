package infinite

import (
	"fmt"
)

// EventsService queries the events of the II API.
type EventsService service

// Get a specific event
func (v *EventsService) Get(id string) (Event, error) {

	eventResponse := new(EventResponse)
	errorResponse := new(ErrorResponse)

	path := fmt.Sprintf("%s", id)

	resp, err := v.sling.New().Get(path).Receive(eventResponse, errorResponse)

	return eventResponse.Event, newError(resp, err, errorResponse)

}

func (v *EventsService) CurrentVerified() ([]Event, error) {

	return v.list(true, true)

}

func (v *EventsService) list(current, verified bool) ([]Event, error) {

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

	resp, err := v.sling.New().Get(path).Receive(eventsResponse, errorResponse)

	// fmt.Printf("request: %v\n", resp.Request)
	// fmt.Printf("response: %v\n", resp)

	return eventsResponse.Events, newError(resp, err, errorResponse)

}
