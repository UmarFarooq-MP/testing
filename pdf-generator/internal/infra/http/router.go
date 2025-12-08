package http

import "github.com/go-chi/chi/v5"

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/api/v1/students/{id}/report", h.GetStudentReport)
	return r
}
