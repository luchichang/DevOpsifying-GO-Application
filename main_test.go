package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test for homeHandler
func TestHomeHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use httptest to create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	// Call the handler with our ResponseRecorder and request
	handler.ServeHTTP(rr, req)

	// Check if the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

// Test for aboutHandler
func TestAboutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aboutHandler)

	handler.ServeHTTP(rr, req)

	// Check for correct status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check for correct response body
	//expected := "This is the About Page!"
	//	if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	//	}
}

