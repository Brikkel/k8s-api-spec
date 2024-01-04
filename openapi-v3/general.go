package openapiv3

import (
	"fmt"
	"io"
	"net/http"
)

// Get the Json object from a provided endpoint
func getJson(endpointURL string) ([]byte, error) {
	// Make an HTTP GET request
	resp, err := http.Get(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	return body, nil
}
