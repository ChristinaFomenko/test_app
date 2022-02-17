package handler

import (
	"encoding/json"
	"errors"
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
	w.Write(mes)
}
