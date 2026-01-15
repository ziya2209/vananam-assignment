package vananamassignment

import (
	"net/http"
	"testing"

	"github.com/therewardstore/httpmatter"
)

func init() {
	_ = httpmatter.Init(&httpmatter.Config{
		BaseDir:       "testdata",
		FileExtension: ".http",
	})
}

func TestClient_GetBangalore_Success(t *testing.T) {
	// Set up httpmatter to mock the HTTP response
	h := httpmatter.NewHTTP(t, "cities").
		Add("get_bangalore_request", "get_bangalore_success_response").
		Respond(nil)

	h.Init()
	defer h.Destroy()

	// Create client with real http.Client (will be mocked by httpmatter)
	client := NewClient(&http.Client{})
	
	// Call the method under test
	response, err := client.GetBangalore()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Validate response
	if response.Status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, response.Status)
	}

	if response.Body == "" {
		t.Error("expected non-empty response body")
	}
}

func TestClient_GetBangalore_Failure(t *testing.T) {
	// Set up httpmatter to mock a failure response
	h := httpmatter.NewHTTP(t, "cities").
		Add("get_bangalore_request", "get_bangalore_error_response").
		Respond(nil)

	h.Init()
	defer h.Destroy()

	// Create client with real http.Client (will be mocked by httpmatter)
	client := NewClient(&http.Client{})
	
	// Call the method under test
	response, err := client.GetBangalore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Validate that we got an error status code
	if response.Status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, response.Status)
	}

	// Validate error response body
	if response.Body == "" {
		t.Error("expected non-empty error response body")
	}
}

func TestClient_GetCity_Success(t *testing.T) {
	// Set up httpmatter to mock the HTTP response
	h := httpmatter.NewHTTP(t, "cities").
		Add("get_city_request", "get_city_success_response").
		Respond(nil)

	h.Init()
	defer h.Destroy()

	// Create client with real http.Client (will be mocked by httpmatter)
	client := NewClient(&http.Client{})
	
	// Call the method under test
	response, err := client.GetCity("Mumbai")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Validate response
	if response.Status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, response.Status)
	}
}
