package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Mock request with query parameter "name"
	req, err := http.NewRequest("GET", "/city?name=London", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the mock request and ResponseRecorder
	getWeather(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode the response body
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(rr.Body).Decode(&weatherResponse); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Verify the city name in the response
	expectedCity := "London"
	if weatherResponse.City != expectedCity {
		t.Errorf("Unexpected city name in response: got %s want %s",
			weatherResponse.City, expectedCity)
	}

}

func TestPostWeather(t *testing.T) {
	// Create a request body with JSON payload
	requestBody := map[string]string{"name": "London"}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a mock request with the JSON payload
	req, err := http.NewRequest("POST", "/city", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the mock request and ResponseRecorder
	postWeather(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(rr.Body).Decode(&weatherResponse); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Verify the city name in the response
	expectedCity := "London"
	if weatherResponse.City != expectedCity {
		t.Errorf("Unexpected city name in response: got %s want %s", weatherResponse.City, expectedCity)
	}
}

var test = "testing"
