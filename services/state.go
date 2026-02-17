package services

import "strconv"

type OperationType int8
const (
	OperationAdd        OperationType = 0
	OperationSubtract   OperationType = 1
	OperationDivide     OperationType = 2
	OperationMultiply   OperationType = 3
)

type RenderState int8
const (
	StateResult    RenderState = 0
	StateInput     RenderState = 1
)

type GlobalState struct {
	Result        float64
  Operation     OperationType
	Input         string
	RenderState   RenderState
	Error         bool
}

var state GlobalState

func (s GlobalState) Rendered() string {
	if state.RenderState == StateInput {
		return s.Input
	}

	return strconv.FormatFloat(s.Result, 'f', -1, 64)
}

func (s GlobalState) InputFloat() (float64, error) {
	res, err := strconv.ParseFloat(s.Input, 64)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func State() *GlobalState {
  return &state
}

