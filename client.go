package vananamassignment

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPClient interface for making HTTP requests
// This allows for easy testing by mocking the HTTP client
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// Client represents the HTTP client for fetching city data
type Client struct {
	httpClient HTTPClient
	baseURL    string
}

// NewClient creates a new Client instance
func NewClient(httpClient HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    "http://example.com",
	}
}

// CityResponse represents the response structure
type CityResponse struct {
	Body    string
	Status  int
	Headers http.Header
}

// GetCity fetches city information for the given city name
func (c *Client) GetCity(cityName string) (*CityResponse, error) {
	url := fmt.Sprintf("%s/cities/%s", c.baseURL, cityName)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch city data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &CityResponse{
		Body:    string(body),
		Status:  resp.StatusCode,
		Headers: resp.Header,
	}, nil
}

// GetBangalore fetches information for Bangalore
func (c *Client) GetBangalore() (*CityResponse, error) {
	return c.GetCity("Bangalore")
}
