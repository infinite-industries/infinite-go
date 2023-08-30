package curtaincall

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

var TicketURL = "https://www.theburltickets.com/events/buy-tickets"
var VenueURL = "https://theburlky.com/"
var InfiniteVenueID = "29fbb4f6-2f4c-47dc-8cfa-c9f0c468704d" // Infinite ID for The Burl

var TheBurl = jsonld.Place{
	Type:       "Place",
	Name:       "The Burl",
	HasMap:     "https://goo.gl/maps/MerUrvdgk9u",
	Identifier: "d9107b10-0479-11e8-808e-07b6d0fe568b",
	URL:        "https://theburlky.com",
	Address: jsonld.Address{
		StreetAddress:   "375 Thompson Road",
		AddressLocality: "Lexington",
		PostalCode:      "40508",
		AddressRegion:   "Kentucky",
		AddressCountry:  "US",
	},
}

type Event struct {
	ID                 string     `json:"id"`
	Date               UnixTimeMS `json:"date"`
	RSVPURL            string     `json:"rsvpURL"`
	EarliestSaleDate   UnixTimeMS `json:"earliestSaleDate"`
	DisabledPriceCodes []string   `json:"disabledPriceCodes"`
	BuyTicketsFileName string     `json:"buyTicketsFileName"`
	DoorsOpenTime      int        `json:"doorsOpenTime"` // this is a time-of-day represented as ms since midnight
	EventDesign        struct {
		BuyTicketsSummary     string `json:"buyTicketsSummary"`
		UpcomingEventsSummary string `json:"upcomingEventsSummary"`
	} `json:"eventDesign"`
	OrganizationID                 string     `json:"organizationId"`
	CreatedAt                      UnixTimeMS `json:"createdAt"`
	AccessibleSeating              bool       `json:"accessibleSeating"`
	PurchasePageURL                string     `json:"purchasePageURL"`
	LatestSaleDate                 UnixTimeMS `json:"latestSaleDate"`
	BuyTicketsSameAsUpcomingEvents bool       `json:"buyTicketsSameAsUpcomingEvents"`
	VenueID                        string     `json:"venueId"`
	UpcomingEventsImageHeight      float64    `json:"upcomingEventsImageHeight"`
	UpcomingEventsFileName         string     `json:"upcomingEventsFileName"`
	EventPublishDate               UnixTimeMS `json:"eventPublishDate"`
	BuyTicketsImageURL             string     `json:"buyTicketsImageURL"`
	PriceCodeAssignments           []struct {
		Quantity                 int        `json:"quantity"`
		MinTickets2              int        `json:"minTickets2"`
		MinTickets1              int        `json:"minTickets1"`
		DisplayName1             string     `json:"displayName1"`
		DisplayName2             string     `json:"displayName2"`
		Type                     string     `json:"type"`
		PriceCodeID              string     `json:"priceCodeId"`
		TotalFees2               int        `json:"totalFees2"`
		TotalFees1               int        `json:"totalFees1"`
		PublishSalesScheduleDate UnixTimeMS `json:"publishSalesScheduleDate"`
		SaleEndDate2             UnixTimeMS `json:"saleEndDate2"`
		SaleEndDate1             UnixTimeMS `json:"saleEndDate1"`
		MaxTickets1              int        `json:"maxTickets1"`
		SaleStartDate2           UnixTimeMS `json:"saleStartDate2"`
		MaxTickets2              int        `json:"maxTickets2"`
		SaleStartDate1           UnixTimeMS `json:"saleStartDate1"`
		FaceValue1               int        `json:"faceValue1"`
		FaceValue2               int        `json:"faceValue2"`
		PublishSaleType          string     `json:"publishSaleType"`
	} `json:"priceCodeAssignments"`
	UpdatedAt                UnixTimeMS `json:"updatedAt"`
	UpcomingEventsImageURL   string     `json:"upcomingEventsImageURL"`
	BuyTicketsImageWidth     int        `json:"buyTicketsImageWidth"`
	EventPublishType         string     `json:"eventPublishType"`
	Name                     string     `json:"name"`
	AgeRestriction           string     `json:"ageRestriction"`
	BuyTicketsImageHeight    float64    `json:"buyTicketsImageHeight"`
	UpcomingEventsImageWidth int        `json:"upcomingEventsImageWidth"`
	Status                   string     `json:"status"`
}

type UnixTimeMS struct {
	time.Time
}

func (u *UnixTimeMS) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.UnixMilli(timestamp)
	return nil
}

func (u UnixTimeMS) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", (u.Time.UnixMilli()))), nil
}
