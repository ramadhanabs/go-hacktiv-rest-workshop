# My First API

A simple REST API for managing books, built with Go's standard `net/http` library.

## Features

- **GET /books** — Retrieve all books
- **POST /books** — Add a new book with input validation
- Go 1.22+ method routing (automatic 405 for wrong methods)
- Consistent JSON response wrapper (`APIResponse`)
- Struct-based validation using `go-playground/validator`
- Auto-generated Swagger documentation

## Prerequisites

- Go 1.22+

## Getting Started

```bash
cd my-api
go mod tidy
go run main.go
# Server running di :8080
# Swagger UI → http://localhost:8080/swagger/
```

## API Usage

### Get all books

```bash
curl http://localhost:8080/books
```

### Add a new book

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "Atomic Habits", "author": "James Clear"}'
```

## Response Format

All responses use a consistent wrapper:

```json
{
  "status": "success",
  "message": "Books retrieved",
  "data": [...]
}
```

## Swagger

Swagger UI is available at `http://localhost:8080/swagger/` when the server is running.

To regenerate docs after modifying annotations:

```bash
swag init
```

## Project Structure

```
my-api/
├── docs/       # Generated Swagger docs
├── go.mod      # Module & dependencies
├── go.sum      # Dependency checksums
└── main.go     # API entry point
```
