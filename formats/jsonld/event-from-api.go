package jsonld

import (
	infinite "github.com/infinite-industries/infinite-go"
)

func NewEventFromApiV1(e infinite.Event) Event {
	event := NewEvent()
	event.Name = e.Title
	event.Identifier = e.ID
	event.BriefDescription = e.BriefDescription
	event.Description = e.Description
	event.URL = "https://infinite.industries/events/" + e.ID
	event.UpdatedAt = e.UpdatedAt

	if e.Image != "" {
		event.Image = []string{e.Image}
	}

	// build SubjectOf and Offers other web pages
	if url := e.WebsiteLink; url != "" {
		event.SubjectOf = append(event.SubjectOf, NewWebPage(url, "Website"))
	}
	if url := e.FbEventLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, NewWebPage(url, "Facebook Event"))
	}
	if url := e.EventbriteLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, NewWebPage(url, "Eventbrite Link"))
		event.Offers = append(event.Offers, NewOffer(url))
	}
	if url := e.TicketLink.String(); url != "" {
		event.SubjectOf = append(event.SubjectOf, NewWebPage(url, "Ticket Link"))
		event.Offers = append(event.Offers, NewOffer(url))
	}

	event.Location = NewPlaceFromApiV1(e.Venue)

	event.StartDate = e.DateTimes[0].StartTime
	event.EndDate = e.DateTimes[0].EndTime
	return event
}

func NewPlaceFromApiV1(v infinite.Venue) Place {
	place := NewPlace(v.Name)
	place.Identifier = v.ID
	place.Address.StreetAddress = v.Street
	place.Address.AddressLocality = v.City
	place.Address.PostalCode = v.Zip
	place.Address.AddressRegion = v.State
	place.Address.AddressCountry = "US"
	place.HasMap = v.GMapLink
	return place
}
