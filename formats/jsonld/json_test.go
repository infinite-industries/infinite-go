package jsonld

import (
	"encoding/json"
	"testing"
)

const ii_event = `
{
  "@context": "https://schema.org",
  "@type": "Event",
  "name": "Lakecia Benjamin",
  "startDate": "2022-05-25T23:30:00.000Z",
  "endDate": "2022-05-26T01:30:00.000Z",
  "location": {
    "@type": "Place",
    "name": "The Kentucky Theatre",
    "address": {
      "@type": "PostalAddress",
      "streetAddress": "214 East Main Street",
      "addressLocality": "Lexington",
      "postalCode": "40507",
      "addressRegion": "KY",
      "addressCountry": "US",
      "identifier": "dbc17929-fa2e-4453-8f31-c3d4e21e14d0"
    },
    "hasMap": "https://goo.gl/maps/3kevicLsj6r"
  },
  "image": [
    "https://s3-us-west-2.amazonaws.com/infinite-industries-event-images/uploads/cef86545-616a-4bbc-8593-7119e9cf32e7.jpg"
   ],
  "brief_description": "Charismatic and dynamic saxophonist/bandleader Lakecia Benjamin at the historic Kentucky Theatre",
  "description": "<p>Charismatic and dynamic saxophonist/bandleader Lakecia Benjamin's electric stage presence and fiery sax work, has shared stages with Stevie Wonder, Alicia Keys, The Roots, Macy Gray and many others, brings the soul and funk up to a fever pitch in a show that's guaranteed to keep crowds dancing day or night.</p><p>Recorded after recovering from a traffic accident that left her jaw shattered, Lakecia Benjamin's 2023 album <em>Phoenix</em> has taken the jazz world by storm. It features collaborations with Terri Lyne Carrington, Wayne Shorter, Angela Davis, and Dianne Reeves. We're excited to present this dynamic jazz artist at the historic Kentucky Theatre.</p>",
  "offers": [
		{
			"@type": "Offer",
			"url": "https://www.tickettailor.com/events/originsjazzseries/891558",
			"price": "30",
			"priceCurrency": "USD",
			"validFrom": "2023-05-01T16:00:00-05:00"
		}
	],
  "performer": {
    "@type": "PerformingGroup",
    "name": "Lakecia Benjamin"
  },
  "url": "https://www.originsjazz.org/shows/lakecia-benjamin",
  "subjectOf": [
    {
      "@type": "WebPage",
      "url": "https://facebook-event/link",
      "name": "Facebook Event"
    }
  ]
}
`

func Test_Unmarshal(t *testing.T) {
	e := Event{}

	b := []byte(ii_event)

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
