package handler

import (
	"context"
	"github.com/ChristinaFomenko/test_app/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculation(t *testing.T) {
	type testTable struct {
		input  model.Input
		output model.Result
	}
	service := NewCalculator(10)
	ctx := context.Background()
	testCases := []testTable{
		{
			input: model.Input{
				A:         2,
				B:         2,
				Operation: "+",
			},
			output: model.Result{
				Success: true,
				ErrCode: "",
				Value:   4,
			},
		},
		{
			input: model.Input{
				A:         5,
				B:         5,
				Operation: "*",
			},
			output: model.Result{
				Success: true,
				ErrCode: "",
				Value:   25,
			},
		},
		{
			input: model.Input{
				A:         10,
				B:         5,
				Operation: "-",
			},
			output: model.Result{
				Success: true,
				ErrCode: "",
				Value:   5,
			},
		},
		{
			input: model.Input{
				A:         30,
				B:         5,
				Operation: "/",
			},
			output: model.Result{
				Success: true,
				ErrCode: "",
				Value:   6,
			},
		},
	}

	for _, testCase := range testCases {
		output := service.Calculate(ctx, testCase.input)
		assert.Equal(t, output, testCase.output)
	}
}
