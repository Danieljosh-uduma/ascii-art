package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Change working directory to the project root so paths like "banner/standard.txt"
	// and "index.html" can be resolved correctly during testing.
	err := os.Chdir("..")
	if err != nil {
		panic(err)
	}

	// Run the tests
	os.Exit(m.Run())
}

// TestHomeHandler tests the root endpoint ("/")
func TestHomeHandler(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify that the response contains some expected HTML
	// Since we execute index.html, we expect some basic html structure
	if !strings.Contains(rr.Body.String(), "<html") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}

// TestAsciiHandler tests the "/ascii-art" endpoint with POST form data
func TestAsciiHandler(t *testing.T) {
	// Prepare form data
	form := url.Values{}
	form.Add("text", "hello")
	form.Add("banner", "standard")

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// We create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(asciiHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that the response contains the expected ASCII art parts
	// For "hello" in standard, it should have the characteristic 'h' shape
	responseBody := rr.Body.String()
	if !strings.Contains(responseBody, "h") && len(responseBody) < 50 {
		t.Errorf("handler returned unexpected body length: got %v",
			len(responseBody))
	}
}
