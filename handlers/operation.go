package handlers

import (
	"goculator/components"
	"goculator/services"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type OperationHandler struct {}

func NewOperationHandler() OperationHandler {
	return OperationHandler{}
}

func doEquals(w http.ResponseWriter, r *http.Request) error {
		err := services.DoEquals()
		if err != nil {
			components.Bsod().Render(r.Context(), w)
			return err
		}
		
		return nil
}

func doOperator(w http.ResponseWriter, r *http.Request, op services.OperationType) error {
	state := services.State()
	var err error
	switch state.RenderState {
	case services.StateInput:
		err = doEquals(w, r)
	case services.StateResult:
		state.Input = ""
	}

	state.Operation = op
	return err
}

func (h OperationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation := vars["operation"]
	state := services.State()

	state.Error = false
	switch operation {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		if state.RenderState == services.StateResult {
			state.Input = ""
		}
		state.RenderState = services.StateInput
		state.Input += operation
	case "period":
		if strings.Contains(state.Input, ".") {
			components.Bsod().Render(r.Context(), w)
			return
		}

		state.Input += "."
	case "eq":
		if err := doEquals(w, r); err != nil {
			return
		}
	case "add":
		if err := doOperator(w, r, services.OperationAdd); err != nil {
			return
		}
	case "sub":
		if err := doOperator(w, r, services.OperationSubtract); err != nil {
			return
		}
	case "mul":
		if err := doOperator(w, r, services.OperationMultiply); err != nil {
			return
		}
	case "div":
		if err := doOperator(w, r, services.OperationDivide); err != nil {
			return
		}
	case "clear":
		state.Operation = services.OperationAdd
		state.Result = 0
		state.RenderState = services.StateResult
		state.Input = ""
		state.Error = false
	}

	components.Input(state.Rendered()).Render(r.Context(), w)
}
