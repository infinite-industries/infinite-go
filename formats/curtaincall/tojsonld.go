package curtaincall

import (
	"fmt"
	"net/url"
	"time"

	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

func (e Event) ToJsonLD() jsonld.Event {
	event := jsonld.NewEvent()
	event.Name = e.Name
	event.Identifier = e.ID
	event.BriefDescription = e.Name
	event.Description = e.EventDesign.UpcomingEventsSummary
	event.UpdatedAt = e.UpdatedAt.Time

	if e.UpcomingEventsImageURL != "" {
		event.Image = []string{e.UpcomingEventsImageURL}
	}

	ticket_link, _ := url.JoinPath(TicketURL, e.ID)

	event.SubjectOf = append(event.SubjectOf, jsonld.NewWebPage(ticket_link, "Ticket Link"))
	event.Offers = append(event.Offers, jsonld.Offer{
		Type:          "Offer",
		URL:           ticket_link,
		Price:         fmt.Sprintf("%d", e.PriceCodeAssignments[0].FaceValue1/100),
		PriceCurrency: "USD",
		ValidFrom:     e.PriceCodeAssignments[0].SaleStartDate1.Time,
	})

	event.Location = TheBurl

	event.StartDate = e.Date.Time
	event.EndDate = e.Date.Time.Add(time.Hour * time.Duration(3))
	// .Add(time.Millisecond * time.Duration(e.DoorsOpenTime))
	//	event.EndDate = e.DateTimes[0].EndTime
	return event

}
