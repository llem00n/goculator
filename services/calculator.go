package services

import (
	"errors"
	"strconv"
)

func DoEquals() error {
  val, err := strconv.ParseFloat(state.Input, 64)
	state.RenderState = StateResult
	if err != nil {
    return err
	}

	switch state.Operation {
	case OperationAdd:
		state.Result += val
	case OperationSubtract:
		state.Result -= val
	case OperationDivide:
		if val == 0 {
			return errors.New("UH!")
		}
		state.Result /= val
	case OperationMultiply:
		state.Result *= val
	default:
		return errors.New("Bad operation")
	}

	return nil
}
