package pdf

import (
	"bytes"
	"fmt"

	"pdf-generator/internal/domain"

	"github.com/jung-kurt/gofpdf"
)

func GenerateStudentPDF(student domain.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report")

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 14)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s", student.Name))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Course: %s", student.Course))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Grade: %s", student.Grade))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("GPA: %.2f", student.GPA))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
