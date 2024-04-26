package api

import (
	"slackbot/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Mux     *chi.Mux
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	h := &Handler{service: service}
	h.Mux = h.SetupRoutes()

	return h
}
