package handlers

import (
	"goculator/components"
	"goculator/services"
	"net/http"
)

type PageHandler struct {}

func NewPageHandler() PageHandler {
	return PageHandler{}
}

func (h PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	components.Page(services.State().Rendered()).Render(r.Context(), w)
}
