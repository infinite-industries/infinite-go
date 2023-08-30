package infinite

import (
	"encoding/json"
	"time"
)

type Version struct {
	Version           string   `json:"version"`
	SupportedVersions []string `json:"supportedVersions"`
}

type Event struct {
	ID               string     `json:"id,omitempty"`
	AdmissionFee     string     `json:"admission_fee"`
	BitlyLink        string     `json:"bitly_link,omitempty"`
	BriefDescription string     `json:"brief_description"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	Description      string     `json:"description"`
	EventbriteLink   NullString `json:"eventbrite_link,omitempty"`
	FbEventLink      NullString `json:"fb_event_link,omitempty"`
	Image            string     `json:"image"`
	Links            []string   `json:"links,omitempty"`
	Contact          string     `json:"organizer_contact,omitempty"`
	ReviewedByOrg    NullString `json:"reviewed_by_org"`
	Slug             string     `json:"slug"`
	SocialImage      NullString `json:"social_image"`
	Tags             []string   `json:"tags"`
	TicketLink       NullString `json:"ticket_link"`
	Title            string     `json:"title"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
	VenueID          string     `json:"venue_id,omitempty"`
	Verified         bool       `json:"verified"`
	WebsiteLink      NullString `json:"website_link"`
	Venue            *Venue     `json:"venue,omitempty"`
	DateTimes        []DateTime `json:"date_times"`
}

type DateTime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Timezone  string    `json:"timezone,omitempty"`
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

// The API can return null for a number of strings - NullString type helps deal
// with that.
type NullString struct {
	Str string
}

// meet the fmt.Stringer interface
func (n NullString) String() string {
	return n.Str
}

// custom json unmarshaller
func (n *NullString) UnmarshalJSON(b []byte) error {
	e := json.Unmarshal(b, &n.Str)
	return e
}

func (n NullString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.String() + `"`), nil
}

func (n NullString) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}
