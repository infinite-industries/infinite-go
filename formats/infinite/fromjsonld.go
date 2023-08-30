package infinite

import (
	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

func EventFromJsonLD(event jsonld.Event) Event {

	e := Event{
		Title:            event.Name,
		ID:               event.Identifier,
		BriefDescription: event.BriefDescription,
		Description:      event.Description,
		UpdatedAt:        &event.UpdatedAt,
	}

	if len(event.Image) > 1 {
		e.Image = event.Image[0]
	}

	// look at offers
	// - if the url of the offer is eventbrite, populate EventbriteLink
	// - if the name of the link is "Ticket Link" - populate TicketLink
	// - otherwise, the first link not matching is the ticket link

	// look at subjectOf
	// - if there is a fb event link populate FbEventLink
	// - if the name of the link is "Website" - populate WebsiteLink
	// - first url to fall through is considered the website for the event

	/*
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
	*/

	v := VenueFromJsonLD(event.Location)
	e.Venue = &v
	e.DateTimes = append(e.DateTimes, DateTime{
		StartTime: event.StartDate,
		EndTime:   event.EndDate,
	})

	return e
}

func VenueFromJsonLD(place jsonld.Place) Venue {
	v := Venue{
		ID:       place.Identifier,
		Street:   place.Address.StreetAddress,
		City:     place.Address.AddressLocality,
		Zip:      place.Address.PostalCode,
		State:    place.Address.AddressRegion,
		GMapLink: place.HasMap, // TODO: doublecheck URL to ensure it is a a GMap link
	}

	return v
}
