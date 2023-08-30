package infinite

import (
	"log"

	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

func (e Event) ToJsonLD() jsonld.Event {

	event := jsonld.NewEvent()
	event.Name = e.Title
	event.Identifier = e.ID
	event.BriefDescription = e.BriefDescription
	event.Description = e.Description
	event.URL = "https://infinite.industries/events/" + e.ID
	event.UpdatedAt = *e.UpdatedAt

	if e.Image != "" {
		event.Image = []string{e.Image}
	}

	// build SubjectOf and Offers other web pages
	if url := e.WebsiteLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, jsonld.NewWebPage(url, "Website"))
	}
	if url := e.FbEventLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, jsonld.NewWebPage(url, "Facebook Event"))
	}
	if url := e.EventbriteLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, jsonld.NewWebPage(url, "Eventbrite Link"))
		event.Offers = append(event.Offers, jsonld.NewOffer(url))
	}
	if url := e.TicketLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, jsonld.NewWebPage(url, "Ticket Link"))
		event.Offers = append(event.Offers, jsonld.NewOffer(url))
	}

	log.Printf("event.Offers: %+v", event.Offers)

	event.Location = e.Venue.ToJsonLD()

	event.StartDate = e.DateTimes[0].StartTime
	event.EndDate = e.DateTimes[0].EndTime
	return event
}

func (v Venue) ToJsonLD() jsonld.Place {
	place := jsonld.NewPlace(v.Name)
	place.Identifier = v.ID
	place.Address.StreetAddress = v.Street
	place.Address.AddressLocality = v.City
	place.Address.PostalCode = v.Zip
	place.Address.AddressRegion = v.State
	place.Address.AddressCountry = "US"
	place.HasMap = v.GMapLink
	return place
}
