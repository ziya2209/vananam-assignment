# Go HTTP Client Assignment

A Go module that demonstrates HTTP client functionality with comprehensive unit tests and CI/CD integration.

## What the project does

This project implements an HTTP client in Go that fetches city information from a remote endpoint (`http://example.com/cities/Bangalore`). The implementation focuses on:

- **Testability**: Code is structured with interfaces to allow easy mocking
- **Testing**: Unit tests using Go's testing framework with `httpmatter` for mocking external HTTP calls
- **CI/CD**: Automated testing via GitHub Actions

The client provides:
- `GetCity(cityName string)`: Fetches information for any city
- `GetBangalore()`: Convenience method to fetch Bangalore-specific data

**Note**: The endpoint is not real and all HTTP calls are mocked in tests to avoid actual network requests.

## Project Structure

```
.
├── client.go              # Main HTTP client implementation
├── client_test.go         # Unit tests with mocked HTTP calls
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
├── testdata/              # HTTP fixture files for mocking
│   └── cities/
│       ├── get_bangalore_request.http
│       ├── get_bangalore_success_response.http
│       ├── get_bangalore_error_response.http
│       ├── get_city_request.http
│       └── get_city_success_response.http
├── .github/
│   └── workflows/
│       └── ci.yml         # GitHub Actions CI workflow
└── README.md              # This file
```

## How to run tests locally

### Prerequisites

- Go 1.24.6 or later
- Internet connection (for downloading dependencies on first run)

### Running Tests

To run all tests:

```bash
go test ./...
```

To run tests with verbose output:

```bash
go test ./... -v
```

To run tests with coverage:

```bash
go test ./... -cover
```

### Test Coverage

The test suite includes:
- ✅ Successful HTTP response handling
- ✅ Error response handling (500 status codes)
- ✅ Multiple city name support
- ✅ Response body and status code validation

All tests use `httpmatter` to mock HTTP calls, ensuring no real network requests are made during testing.

## Dependencies

- `github.com/therewardstore/httpmatter`: HTTP mocking library for Go tests

## CI/CD

This repository includes a GitHub Actions workflow (`.github/workflows/ci.yml`) that:
- Runs on every push to the repository
- Executes `go test ./...`
- Fails the build if any tests fail

The workflow ensures code quality and test coverage on every commit.
