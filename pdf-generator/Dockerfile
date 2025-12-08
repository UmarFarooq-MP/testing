# ----------- BUILD STAGE -----------
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o pdf-generator ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN adduser -D appuser
WORKDIR /app
COPY --from=builder /app/pdf-generator .
USER appuser
EXPOSE 8081
ENTRYPOINT ["./pdf-generator"]
