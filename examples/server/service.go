package main

import (
	"os"

	infinite "github.com/infinite-industries/infinite-go"
	"github.com/rs/zerolog"
)

type Service struct {
	prefix_path string
	log         zerolog.Logger
	client      *infinite.Client
}

func NewService(optionFuncs ...optionFunc) *Service {

	srv := Service{
		prefix_path: "/",
		log:         zerolog.New(os.Stderr).With().Str("component", "jsonld").Logger(),
		client:      infinite.New(),
	}

	for _, option := range optionFuncs {
		option(&srv)
	}
	srv.log.Debug().Msgf("service bound to %s", srv.prefix_path)
	return &srv
}

type optionFunc func(*Service)

func SetPrefix(prefix string) optionFunc {
	return func(s *Service) {
		if prefix != "" {
			s.prefix_path = prefix
		}
	}
}

func SetLogger(log zerolog.Logger) optionFunc {
	return func(s *Service) {
		s.log = log.With().Str("component", "jsonld").Logger()
	}
}

func SetClient(client *infinite.Client) optionFunc {
	return func(s *Service) {
		s.client = client
	}
}
