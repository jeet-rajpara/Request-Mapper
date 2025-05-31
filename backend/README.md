# Request Mapper Backend

This is the backend API service for the Request Mapper project, built with Go and Gin framework.

## Prerequisites

- Go 1.21 or higher
- Git

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/jeet-rajpara/Request-Mapper.git
cd backend
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## Project Structure

```
backend/
├── api/
│   ├── controller/    # HTTP request handlers
│   ├── repository/    # Data access layer
│   └── service/       # Business logic
├── error/            # Error handling 
├── main.go           # Application entry point
├── go.mod            # Go module file
└── go.sum            # Go module checksum
```

## API Endpoints

### Map Request
- **URL**: `/api/map-request`
- **Method**: `POST`
- **Description**: Maps input JSON according to the provided request map
- **Request Body**:
  ```json
  {
    "inputJSON": {
      "customer": {
        // customer data
      }
    },
    "requestMap": {
      // mapping configuration
    }
  }
  ```
- **Response**: Mapped JSON object
- **Error Codes**:
  - 400: Bad Request (invalid input)
  - 500: Internal Server Error
