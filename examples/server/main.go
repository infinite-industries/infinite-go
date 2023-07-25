package main

import (
	"net/http"
	"os"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

var log = zerolog.New(os.Stderr).With().Timestamp().Logger()

func init() {
	initConfig()

	log = httplog.NewLogger("ii", httplog.Options{
		LogLevel:        loglevel, // from config
		JSON:            true,
		TimeFieldFormat: time.RFC3339,
		TimeFieldName:   "time",
	})
}

func main() {

	log.Info().Msgf("Log level set to %s", loglevel)
	log.Debug().Msgf("config initialized: address=%v, prefix_path=%v, loglevel=%v", address, prefix_path, loglevel)

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(log))

	log.Debug().Msgf("binding middleware.Hearbeat to %sping", prefix_path)
	router.Use(middleware.Heartbeat(prefix_path + "ping"))

	svc := NewService()

	router.Route(prefix_path, svc.Router)

	log.Info().Msgf("starting server on %s", address)

	http.ListenAndServe(address, router)

}
