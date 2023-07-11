package handlers

import (
	"fmt"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Route() string {
	return "/health"
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (h *HealthHandler) Methods() []string {
	return []string{http.MethodGet}
}
