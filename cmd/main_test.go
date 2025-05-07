package main

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
func TestDbStatus(t *testing.T) {
	// Setup the app
	app := setupApp()

	// Create a new request
	req := httptest.NewRequest("GET", "/db-status", nil)

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	// Check status code - this will likely be 500 if db is nil in tests
	assert.Equal(t, 500, resp.StatusCode, "Status code should be 500 when database is not connected")

	// Check response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check if body indicates database not connected
	assert.Equal(t, "Database not connected", string(body), "Response body should indicate database is not connected")
}

func TestCustomerAccounts(t *testing.T) {

	// Setup a mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Set the global db variable to our mock
	db = mockDB

	// Setup mock expectations
	rows := sqlmock.NewRows([]string{"account_number", "account_type", "status"}).
		AddRow("ACC123456", "CHECKING", "ACTIVE").
		AddRow("ACC789012", "SAVINGS", "ACTIVE")

	mock.ExpectQuery("SELECT account_number, account_type, status FROM accounts WHERE customer_id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	// Setup the app
	app := setupApp()

	// Create a new request
	req := httptest.NewRequest("GET", "/customers/me/accounts", nil)

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	// Check status code - should be 200 now that we have a mock DB
	assert.Equal(t, 200, resp.StatusCode, "Status code should be 200")

	// Check response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Parse the JSON response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	// Check if accounts exist in the response
	accounts, exists := response["accounts"]
	assert.True(t, exists, "Response should contain accounts field")

	// Check the accounts data
	accountsArray, ok := accounts.([]interface{})
	assert.True(t, ok, "Accounts should be an array")
	assert.Equal(t, 2, len(accountsArray), "Should have 2 accounts")

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestInvalidRoute(t *testing.T) {
	// Setup the app
	app := setupApp()

	// Create a new request to a non-existent route
	req := httptest.NewRequest("GET", "/non-existent-route", nil)

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	// Check status code
	assert.Equal(t, 404, resp.StatusCode, "Status code should be 404 for non-existent route")
}

func TestMethodNotAllowed(t *testing.T) {
	// Setup the app
	app := setupApp()

	// Create a POST request to a GET-only endpoint
	req := httptest.NewRequest("POST", "/", nil)

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	// Check status code
	assert.Equal(t, 405, resp.StatusCode, "Status code should be 405 Method Not Allowed")
}
