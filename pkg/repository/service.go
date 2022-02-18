package repository

import (
	"context"
	"github.com/ChristinaFomenko/test_app/pkg/model"
)

type Service interface {
	Calculate(c context.Context, inputStruct model.Input) model.Result
}
