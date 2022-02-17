package model

type Input struct {
	A         int
	B         int
	Operation string
}

type Result struct {
	Success bool
	ErrCode string
	Value   float64
}
