package main

import (
	"encoding/json"
	"testing"

	infinite "github.com/infinite-industries/infinite-go"
)

const new_event = `
{
  "venue_id": "f467e7a0-a066-11ea-aa51-cdc3fe7afefa",
  "title": "Infinite Gallery Opening",
  "slug": "infinite-gallery-opening",
  "multi_day": false,
  "image": "https://i.picsum.photos/id/258/536/354.jpg?hmac=FJZvafgClrsfFxn1Ce6YeBIo2958pGQCb4jCbEc3SRA",
  "social_image": "https://picsum.photos/536/354",
  "admission_fee": "5",
  "organizer_contact": "bob.vance@refridgeration.com",
  "brief_description": "The gallery is open",
  "description": "<h2>The Gallery Is Open</h2><p>Some details</p>",
  "website_link": "https://www.wegotrats.com",
  "ticket_link": "https://www.wegotussometickets.com",
  "fb_event_link": "https://thebookwhoshallnotbenames.com/foobar",
  "eventbrite_link": "https://eventbrite.com/foobar",
  "bitly_link": "https://bitly.com/foobar",
  "reviewed_by_org": "radio-mc-radio-station",
  "date_times": [
    {
      "start_time": "2023-08-14T13:56:01.817Z",
      "end_time": "2023-08-15T13:56:01.817Z"
    }
  ],
  "tags": [
    "mode:online"
  ],
  "links": []
}
`

func TestCreateEvent(t *testing.T) {
	client := infinite.New()

	e := infinite.Event{}

	b := []byte(new_event)

	err := json.Unmarshal(b, &e)
	if err != nil {
		t.Error(err)
		return
	}

	new_evt, err := client.Events.Create(e)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Created new event: %+v", new_evt)
	}

}
