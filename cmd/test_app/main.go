package main

import (
	"github.com/ChristinaFomenko/test_app/pkg/config"
	"github.com/ChristinaFomenko/test_app/pkg/handler"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	service := handler.NewCalculator(config.Timeout.ContextTimeout)
	handler.NewCalculatorHandler(r, service)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
	}

	log.Info().Msgf("Server started at " + "localhost:8000")
	log.Error().Msgf(srv.ListenAndServe().Error())
}
