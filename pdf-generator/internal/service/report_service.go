package service

import (
	"pdf-generator/internal/domain"
	"pdf-generator/internal/infra/client"
	"pdf-generator/internal/infra/pdf"
)

type ReportService struct {
	client *client.StudentsClient
}

func NewReportService(c *client.StudentsClient) domain.Report {
	return &ReportService{client: c}
}

func (s ReportService) GenerateStudentReport(id string) ([]byte, error) {
	student, err := s.client.GetStudent(id)
	if err != nil {
		return nil, err
	}

	return pdf.GenerateStudentPDF(student)
}
