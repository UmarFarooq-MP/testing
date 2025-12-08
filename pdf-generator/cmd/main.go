package main

import (
	"fmt"
	"net/http"
	"pdf-generator/internal/infra/client"
	httpHandler "pdf-generator/internal/infra/http"
	"pdf-generator/internal/service"
)

func main() {
	nodeURL := "http://localhost:3000"

	client := client.NewStudentsClient(nodeURL)
	reportService := service.NewReportService(client)
	handler := httpHandler.NewHandler(reportService)
	router := httpHandler.NewRouter(handler)

	fmt.Println("PDF Generator Service running on :8081")
	http.ListenAndServe(":8081", router)
}
