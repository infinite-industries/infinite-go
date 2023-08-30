package jsonld

import "time"

type Event struct {
	Name             string           `json:"name"`
	StartDate        time.Time        `json:"startDate"`
	EndDate          time.Time        `json:"endDate"`
	Location         Place            `json:"location"`
	Identifier       string           `json:"identifier"`
	Image            []string         `json:"image,omitempty"`
	BriefDescription string           `json:"brief_description"`
	Description      string           `json:"description"`
	Offers           []Offer          `json:"offers,omitempty"`
	Performer        *PerformingGroup `json:"performer,omitempty"`
	URL              string           `json:"url"`
	SubjectOf        []WebPage        `json:"subjectOf,omitempty"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	Type             string           `json:"@type"`    // Event
	Context          string           `json:"@context"` // https://schema.org
}

func NewEvent() Event {
	return Event{
		Type:    "Event",
		Context: "https://schema.org",
	}
}

type Address struct {
	StreetAddress   string `json:"streetAddress"`
	AddressLocality string `json:"addressLocality"`
	PostalCode      string `json:"postalCode"`
	AddressRegion   string `json:"addressRegion"`
	AddressCountry  string `json:"addressCountry"`
	Type            string `json:"@type"` // Address
}

func NewAddress() Address {
	return Address{
		Type: "PostalAddress",
	}
}

type Place struct {
	Name       string  `json:"name"`
	Address    Address `json:"address,omitempty"`
	HasMap     string  `json:"hasMap,omitempty"`
	Identifier string  `json:"identifier,omitempty"`
	URL        string  `json:"url,omitempty"`
	Type       string  `json:"@type"` // Place
}

func NewPlace(name string) Place {
	return Place{
		Type:    "Place",
		Name:    name,
		Address: NewAddress(),
	}
}

type Offer struct {
	URL           string     `json:"url"`
	Price         string     `json:"price,omitempty"`
	PriceCurrency string     `json:"priceCurrency,omitempty"`
	ValidFrom     *time.Time `json:"validFrom, omitempty"`
	Type          string     `json:"@type"` // Offer
}

func NewOffer(url string) Offer {
	now := time.Now().Round(time.Minute)
	return Offer{
		Type:      "Offer",
		URL:       url,
		ValidFrom: &now,
	}
}

type PerformingGroup struct {
	Name string `json:"name"`
	Type string `json:"@type"` // PerformingGroup
}

func NewPerformingGroup() PerformingGroup {
	return PerformingGroup{
		Type: "PerformingGroup",
	}
}

type WebPage struct {
	URL  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"@type"` // WebPage
}

func NewWebPage(url, name string) WebPage {
	return WebPage{
		Type: "WebPage",
		URL:  url,
		Name: name,
	}
}
