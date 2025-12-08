package test

import (
	"pdf-generator/internal/domain"
	"pdf-generator/internal/infra/pdf"
	"testing"
)

func TestGenerateStudentPDF(t *testing.T) {
	student := domain.Student{
		ID:     1,
		Name:   "John Doe",
		Course: "Mathematics",
		Grade:  "A",
		GPA:    3.9,
	}

	pdfBytes, err := pdf.GenerateStudentPDF(student)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pdfBytes) == 0 {
		t.Fatalf("expected non-empty PDF data")
	}

	if string(pdfBytes[:4]) != "%PDF" {
		t.Fatalf("expected PDF header, got %s", string(pdfBytes[:4]))
	}
}
