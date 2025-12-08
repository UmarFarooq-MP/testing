# PDF Generator Microservice (Go + Chi + Clean Architecture)

This microservice generates **PDF reports for students** by fetching their data from an existing **Node.js backend API**.

It is built using:

- **Go 1.25+**
- **Chi Router**
- **gofpdf** for PDF generation
- **Clean Architecture**
- **In-memory PDF rendering**
- **Docker-ready**

---

## 📁 Project Structure

```
pdf-generator/
│── cmd/
│     └── main.go
│── internal/
│     ├── domain/
│     │     ├── student.go
│     │     └── report.go
│     ├── infra/
│     │     ├── http/
│     │     │     ├── handler.go
│     │     │     └── router.go
│     │     ├── client/
│     │     │     └── students_client.go
│     │     └── pdf/
│     │           └── generator.go
│     └── service/
│           └── report_service.go
│── test/
│     ├── pdf_generator_test.go
│     ├── students_client_test.go
│     └── report_service_test.go
│── go.mod
│── Dockerfile
```

---

## Running the Service

### **1. Run Locally (Without Docker)**

```
go mod tidy
go run ./cmd/main.go
```

Service runs at:

```
http://localhost:8081
```

---

## Running with Docker

### **Build Image**

```
docker build -t pdf-generator .
```

### **Run Container**

```
docker run -p 8081:8081 pdf-generator
```

---

## 🧪 API Endpoint

### **GET /api/v1/students/:id/report**

This endpoint:

1. Calls the Node.js backend at:
   ```
   GET http://localhost:3000/api/v1/students/:id
   ```
2. Receives student JSON like:
   ```json
   {
       "id": 1,
       "name": "John Doe",
       "course": "Mathematics",
       "grade": "A",
       "gpa": 3.9
   }
   ```
3. Generates a PDF with these details
4. Returns the PDF file to the client

### Example cURL:

```
curl -X GET   http://localhost:8081/api/v1/students/1/report   -o student_report.pdf
```

---

## Running Tests

All tests are inside the `test/` folder.

Run all tests:

```
go test ./...
```

Test coverage includes:

- PDF generation logic
- Student client HTTP mocking
- Report service integration

---

## Architecture Overview

This microservice follows a **clean architecture**:

- **Domain** → Pure entities and interfaces
- **Service** → Business logic for PDF generation
- **Infra**
  - HTTP handlers + routing
  - External API client for Node.js
  - PDF generator engine
- **cmd/main.go** → Composition root

---

---

## PDF Layout

The generated PDF includes:

- Student Name  
- Course  
- Grade  
- GPA  
- Title: **Student Report**

The first bytes always begin with:

```
%PDF
```

---
