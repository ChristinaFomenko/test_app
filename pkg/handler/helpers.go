package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChristinaFomenko/test_app/pkg/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func parse(r *http.Request) (model.Input, error) {
	st := model.Input{}
	err := r.ParseForm()
	if err != nil {
		return model.Input{}, errors.New("err when parsing form")
	}
	return st, nil
}

func sendError(errorMessage string, code int, w http.ResponseWriter) {
	log.Error().Msgf(errorMessage)
	w.WriteHeader(code)
	mes, _ := json.Marshal(model.Result{ErrCode: errorMessage})
	_, err := w.Write(mes)
	if err != nil {
		return
	}
}

func SendOKResponse(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	serializedData, err := json.Marshal(data)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	_, err = w.Write(serializedData)
	if err != nil {
		message := fmt.Sprintf("HttpResponse while writing is socket: %s", err.Error())
		log.Error().Msgf(message)
		return
	}
	log.Info().Msgf("OK message sent")
}
