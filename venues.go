package infinite

import (
	"fmt"
)

// VenuesService queries the venues of the II API.
type VenuesService service

// verify license for an organization
func (v *VenuesService) Get(id string) (Venue, error) {

	venueResponse := new(VenueResponse)
	errorResponse := new(ErrorResponse)

	path := fmt.Sprintf("%s", id)

	resp, err := v.sling.New().Get(path).Receive(venueResponse, errorResponse)

	return venueResponse.Venue, newError(resp, err, errorResponse)

}
