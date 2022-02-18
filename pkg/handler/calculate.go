package handler

import (
	"context"
	"github.com/ChristinaFomenko/test_app/pkg/model"
	"time"
)

type Calculator struct {
	contextTimeout time.Duration
}

func NewCalculator(timeout time.Duration) *Calculator {
	return &Calculator{contextTimeout: timeout}
}

func (ca *Calculator) Calculate(c context.Context, params model.Input) model.Result {
	_, cancel := context.WithTimeout(c, ca.contextTimeout)
	defer cancel()
	res := model.Result{
		Success: true,
		ErrCode: "",
		Value:   getValue(params),
	}
	return res
}

func getValue(params model.Input) int {
	var sum int
	switch params.Operation {
	case "+":
		sum = params.A + params.B
	case "-":
		sum = params.A - params.B
	case "*":
		sum = params.A * params.B
	case "/":
		sum = params.A / params.B

	}
	return sum
}
