package domain

type Report interface {
	GenerateStudentReport(id string) ([]byte, error)
}
