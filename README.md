# Go Hacktiv REST Workshop

Build Your First Backend API in 90 Minutes — Back End Golang Workshop by Ramadhana Bagus.

## What's Inside

### [slidev/](./slidev)

Presentation slides built with [Slidev](https://sli.dev), covering:

- REST API concepts (request-response, HTTP methods, status codes, JSON)
- Comparison with GraphQL, WebSocket, gRPC/SSE
- Building a Go REST API from scratch using `net/http`
- Input validation with struct tags
- Swagger documentation
- Unit testing & performance benchmarking

### [my-api/](./my-api)

A working Go REST API extracted from the workshop slides:

- **GET /books** — Retrieve all books
- **POST /books** — Add a new book with validation
- Go 1.22+ method routing
- Response wrapper pattern
- Swagger UI at `/swagger/`

## Quick Start

### Run the slides

```bash
cd slidev
npm install
npm run dev
```

### Run the API

```bash
cd my-api
go mod tidy
go run main.go
# Server running di :8080
# Swagger UI → http://localhost:8080/swagger/
```

## Speaker

**Ramadhana Bagus** — Full-stack Engineer

- [LinkedIn](https://www.linkedin.com/in/ramadhanabagus/)
- [GitHub](https://github.com/ramadhanabs)
