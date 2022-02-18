package handler

import (
	"github.com/ChristinaFomenko/test_app/pkg/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type CalculatorHandler struct {
	Service repository.Service
}

func NewCalculatorHandler(r *mux.Router, s repository.Service) {
	handler := CalculatorHandler{Service: s}
	r.HandleFunc("/api/add", handler.AddHandler).Methods("GET")
	r.HandleFunc("/api/div", handler.DivHandler).Methods("GET")
	r.HandleFunc("/api/mul", handler.MulHandler).Methods("GET")
	r.HandleFunc("/api/sub", handler.SubHandler).Methods("GET")
}

func (c *CalculatorHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "+"
	res := c.Service.Calculate(r.Context(), params)
	SendOKResponse(res, w)
}

func (c *CalculatorHandler) SubHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "-"
	res := c.Service.Calculate(r.Context(), params)
	SendOKResponse(res, w)
}

func (c *CalculatorHandler) MulHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "*"
	res := c.Service.Calculate(r.Context(), params)
	SendOKResponse(res, w)
}

func (c *CalculatorHandler) DivHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "/"
	res := c.Service.Calculate(r.Context(), params)
	SendOKResponse(res, w)

}
