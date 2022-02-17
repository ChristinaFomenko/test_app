package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type CalculatorHandler struct {
}

func NewCalculatorHandler(r *mux.Router) {
	handler := CalculatorHandler{}
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
}

func (c *CalculatorHandler) SubHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "-"
}

func (c *CalculatorHandler) MulHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "*"
}

func (c *CalculatorHandler) DivHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parse(r)
	if err != nil {
		sendError(err.Error(), 400, w)
		return
	}
	params.Operation = "/"

}
