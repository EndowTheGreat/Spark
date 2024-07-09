package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	Router *chi.Mux
}

func NewRouter() *Handler {
	return &Handler{
		Router: chi.NewMux(),
	}
}

func (h *Handler) SetupRoutes(log bool) {
	router := h.Router
	if log {
		router.Use(middleware.Logger)
	}
	router.HandleFunc("/", h.ServeHome)
	router.HandleFunc("/*", h.serveFile)
}
