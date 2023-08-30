package curtaincall

import (
	"encoding/json"
	"testing"
)

const event = `
{
  "id": "SykXTJ8qxsWr52VmvJtI",
  "date": 1692921600000,
  "rsvpURL": "",
  "earliestSaleDate": 1687194000265,
  "disabledPriceCodes": [],
  "buyTicketsFileName": "c25c4967-036c-45bf-ab26-2000980942c3.png",
  "doorsOpenTime": 68400000,
  "eventDesign": {
    "buyTicketsSummary": "<h1><strong>Drew Charron + Urameshi + Kinni Moon &amp; Ground Control</strong></h1><h2>8/24/2023</h2><p><em>INDOOR SHOW</em></p><p>All ages</p><p>Doors 7:00 PM | Show 8:00 PM</p>",
    "upcomingEventsSummary": "<h1><strong>Drew Charron + Urameshi + Kinni Moon &amp; Ground Control</strong></h1><h2>8/24/2023</h2><p><em>INDOOR SHOW</em></p><p>All ages</p><p>Doors 7:00 PM | Show 8:00 PM</p>"
  },
  "organizationId": "Zt3mxndQxagKv9H7hpXr",
  "createdAt": 1687194433377,
  "accessibleSeating": false,
  "purchasePageURL": "",
  "latestSaleDate": 1692928800000,
  "buyTicketsSameAsUpcomingEvents": true,
  "venueId": "HQq1zuKtoiIwwkE9yyTT",
  "upcomingEventsImageHeight": 133.02083333333334,
  "upcomingEventsFileName": "335ba900-30e8-4a6e-b7ba-7fdffe58ac5f.png",
  "eventPublishDate": 1687194433377,
  "buyTicketsImageURL": "https://storage.googleapis.com/gale-648cf.appspot.com/c25c4967-036c-45bf-ab26-2000980942c3.png",
  "priceCodeAssignments": [
    {
      "quantity": 400,
      "minTickets2": 1,
      "minTickets1": 1,
      "displayName1": "Advanced",
      "displayName2": "Day of",
      "type": "STANDARD",
      "priceCodeId": "MtMMeGMauMWvcYfKkXhj",
      "totalFees2": 360,
      "totalFees1": 300,
      "publishSalesScheduleDate": 1687194424595,
      "saleEndDate2": 1692928800000,
      "saleEndDate1": 1692849600000,
      "maxTickets1": 4,
      "saleStartDate2": 1692849600000,
      "maxTickets2": 4,
      "saleStartDate1": 1687194000265,
      "faceValue1": 1500,
      "faceValue2": 1800,
      "publishSaleType": "IMMEDIATELY"
    }
  ],
  "updatedAt": 1687194433377,
  "upcomingEventsImageURL": "https://storage.googleapis.com/gale-648cf.appspot.com/335ba900-30e8-4a6e-b7ba-7fdffe58ac5f.png",
  "buyTicketsImageWidth": 200,
  "eventPublishType": "IMMEDIATELY",
  "name": "Drew Charron + Urameshi + Kinni Moon & Ground Control",
  "ageRestriction": "NONE",
  "buyTicketsImageHeight": 133.02083333333334,
  "upcomingEventsImageWidth": 200,
  "status": "COMPLETE"
}
`

func Test_Unmarshal(t *testing.T) {
	e := Event{}

	b := []byte(event)

	err := json.Unmarshal(b, &e)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", e)

	n, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Logf(string(n))
}

func Test_ToJsonLD(t *testing.T) {
	e := Event{}
	b := []byte(event)
	err := json.Unmarshal(b, &e)
	if err != nil {
		t.Error(err)
	}

	ld_event := e.ToJsonLD()
	t.Logf("%+v", ld_event)

}
