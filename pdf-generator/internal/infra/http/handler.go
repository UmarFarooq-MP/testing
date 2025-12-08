package http

import (
	"net/http"

	"pdf-generator/internal/domain"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc domain.Report
}

func NewHandler(svc domain.Report) *Handler {
	return &Handler{svc: svc}
}

func (h Handler) GetStudentReport(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	pdfData, err := h.svc.GenerateStudentReport(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=student_report.pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfData)
}
