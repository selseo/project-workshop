package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	// Setup the app
	app := setupApp()

	// Create a new request
	req := httptest.NewRequest("GET", "/", nil)

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	// Check status code
	assert.Equal(t, 200, resp.StatusCode, "Status code should be 200")

	// Check response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check if body equals "Hello World"
	assert.Equal(t, "Hello World", string(body), "Response body should be 'Hello World'")
}
