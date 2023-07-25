package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

func (s *Service) Router(r chi.Router) {
	r.Get("/event/{id}", s.GetEvent)
	r.Get("/venue/{id}", s.GetVenue)
}

func (s *Service) GetEvent(w http.ResponseWriter, r *http.Request) {
	event_id := chi.URLParam(r, "id")

	event, err := s.client.Events.Get(event_id)
	if err != nil {
		s.log.Error().Err(err).Msgf("upstream error")
		w.WriteHeader(500)
		fmt.Fprintf(w, "Upstream error: %s", err)
		return
	}

	ldEvent := jsonld.NewEventFromApiV1(event)
	b, _ := json.MarshalIndent(ldEvent, "", "  ")

	w.Header().Set("Content-type", "application/json")
	//	w.Write(b)
	fmt.Fprintf(w, "%s\n", b)
	return
}

func (s *Service) GetVenue(w http.ResponseWriter, r *http.Request) {
	venue_id := chi.URLParam(r, "id")

	place, err := s.client.Venues.Get(venue_id)
	if err != nil {
		s.log.Error().Err(err).Msgf("upstream error")
		w.WriteHeader(500)
		fmt.Fprintf(w, "Upstream error: %s", err)
		return
	}

	ldPlace := jsonld.NewPlaceFromApiV1(place)
	b, _ := json.MarshalIndent(ldPlace, "", "  ")

	w.Header().Set("Content-type", "application/json")
	//	w.Write(b)
	fmt.Fprintf(w, "%s\n", b)
	return
}
