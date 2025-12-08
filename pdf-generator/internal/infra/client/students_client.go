package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pdf-generator/internal/domain"
	"time"
)

type StudentsClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewStudentsClient(baseURL string) *StudentsClient {
	return &StudentsClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *StudentsClient) GetStudent(id string) (domain.Student, error) {
	url := fmt.Sprintf("%s/api/v1/students/%s", c.baseURL, id)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return domain.Student{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Student{}, fmt.Errorf("node api returned status %d", resp.StatusCode)
	}

	var student domain.Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return domain.Student{}, err
	}

	return student, nil
}
