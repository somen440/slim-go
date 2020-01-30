package handler

import (
	"fmt"
	"net/http"
)

type HttpErrorHandler struct {
	status int
}

func NewForbiddenHttpErrorHandler() *HttpErrorHandler {
	return &HttpErrorHandler{
		status: http.StatusForbidden,
	}
}

func (h *HttpErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(h.status)
	fmt.Fprintf(w, "Forbidden")
}
