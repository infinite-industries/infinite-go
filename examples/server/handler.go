package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/infinite-industries/infinite-go/formats/infinite"
	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

func (s *Service) Router(r chi.Router) {
	r.Get("/event/{id}", s.GetEvent)
	r.Get("/venue/{id}", s.GetVenue)
	r.Get("/compare-event/{id}", s.CompareEvent)
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

	ldEvent := event.ToJsonLD()
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

	ldPlace := place.ToJsonLD()
	b, _ := json.MarshalIndent(ldPlace, "", "  ")

	w.Header().Set("Content-type", "application/json")
	//	w.Write(b)
	fmt.Fprintf(w, "%s\n", b)
	return
}

func (s *Service) CompareEvent(w http.ResponseWriter, r *http.Request) {
	event_id := chi.URLParam(r, "id")

	event, err := s.client.Events.Get(event_id)
	if err != nil {
		s.log.Error().Err(err).Msgf("upstream error")
		w.WriteHeader(500)
		fmt.Fprintf(w, "Upstream error: %s", err)
		return
	}

	jsonldEvent := event.ToJsonLD()

	rebuildEvent := infinite.EventFromJsonLD(jsonldEvent)

	type comparison_type struct {
		Original infinite.Event `json:"original"`
		Jsonld   jsonld.Event   `json:"jsonld"`
		Rebuild  infinite.Event `json:"rebuild"`
	}

	comparison := comparison_type{
		Original: event,
		Jsonld:   jsonldEvent,
		Rebuild:  rebuildEvent,
	}

	b, _ := json.MarshalIndent(comparison, "", "  ")

	w.Header().Set("Content-type", "application/json")
	//	w.Write(b)
	fmt.Fprintf(w, "%s\n", b)
	return
}
