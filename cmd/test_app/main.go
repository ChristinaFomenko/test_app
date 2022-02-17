package test_app

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
	}

	log.Info().Msgf("Server started at " + "localhost:8000")
	log.Error().Msgf(srv.ListenAndServe().Error())
}
