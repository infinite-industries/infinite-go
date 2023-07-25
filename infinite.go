package infinite

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	libraryVersion   = "0.1"
	defaultBaseURL   = "https://api.infinite.industries/"
	defaultUserAgent = "infinite-go/" + libraryVersion
)

type Version struct {
	Version           string   `json:"version"`
	SupportedVersions []string `json:"supportedVersions"`
}

type Event struct {
	ID               string     `json:"id"`
	AdmissionFee     string     `json:"admission_fee"`
	BitlyLink        NullString `json:"bitly_link"`
	BriefDescription string     `json:"brief_description"`
	CreatedAt        time.Time  `json:"createdAt"`
	Description      string     `json:"description"`
	EventbriteLink   NullString `json:"eventbrite_link"`
	FbEventLink      NullString `json:"fb_event_link"`
	Image            string     `json:"image"`
	Links            []string   `json:"links"`
	ReviewedByOrg    NullString `json:"reviewed_by_org"`
	Slug             string     `json:"slug"`
	SocialImage      string     `json:"social_image"`
	Tags             []string   `json:"tags"`
	TicketLink       NullString `json:"ticket_link"`
	Title            string     `json:"title"`
	UpdatedAt        time.Time  `json:"updatedAt"`
	VenueID          NullString `json:"venue_id"`
	Verified         bool       `json:"verified"`
	WebsiteLink      string     `json:"website_link"`
	Venue            Venue      `json:"venue"`
	DateTimes        []struct {
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
		Timezone  string    `json:"timezone"`
	} `json:"date_times"`
}

type Venue struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Address       string    `json:"address"`
	GMapLink      string    `json:"g_map_link"`
	GPSLat        float64   `json:"gps_lat"`
	GPSLong       float64   `json:"gps_long"`
	Street        string    `json:"street"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	Neighborhood  string    `json:"neighborhood"`
	IsSoftDeleted bool      `json:"is_soft_deleted"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type EventResponse struct {
	Event  Event  `json:"event"`
	Status string `json:"status"`
}

type EventsResponse struct {
	Events []Event `json:"events"`
	Status string  `json:"status"`
}

type VenueResponse struct {
	Venue  Venue  `json:"venue"`
	Status string `json:"status"`
}

type VenuesResponse struct {
	Venues []Venue `json:"venues"`
	Status string  `json:"status"`
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

type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`
	// For paginated result sets, the number of results to include per page.
	PageSize int `url:"pagesize,omitempty"`
}

// The API can return null for a number of strings - this type helps deal with that.
type NullString struct {
	str string
}

// meet the fmt.Stringer interface
func (n NullString) String() string {
	return n.str
}

// custom json unmarshaller
func (n *NullString) UnmarshalJSON(b []byte) error {
	e := json.Unmarshal(b, &n.str)
	return e
}

func (n NullString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.String() + `"`), nil
}

func (n NullString) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}
