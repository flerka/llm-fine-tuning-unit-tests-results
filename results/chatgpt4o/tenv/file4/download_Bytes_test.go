package download

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mock display function
func mockDisplay(msg string) {}

// Mock http.Client
var httpClient = &http.Client{}

// Mock HTTP Server
func startMockServer(responseStatus int, responseBody string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseStatus)
		io.WriteString(w, responseBody)
	})
	return httptest.NewServer(handler)
}

func TestBytes_Success(t *testing.T) {
	server := startMockServer(http.StatusOK, "mock response")
	defer server.Close()

	data, err := Bytes(server.URL, mockDisplay)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if string(data) != "mock response" {
		t.Fatalf("Expected 'mock response', got %s", string(data))
	}
}

func TestBytes_InvalidURL(t *testing.T) {
	_, err := Bytes(":", mockDisplay)
	if err == nil {
		t.Fatalf("Expected error for invalid URL, got none")
	}
}

func TestBytes_404NotFound(t *testing.T) {
	server := startMockServer(http.StatusNotFound, "404 page not found")
	defer server.Close()

	_, err := Bytes(server.URL, mockDisplay)
	if err == nil {
		t.Fatalf("Expected error for 404 status code, got none")
	}
}

func TestBytes_EmptyResponse(t *testing.T) {
	server := startMockServer(http.StatusOK, "")
	defer server.Close()

	data, err := Bytes(server.URL, mockDisplay)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(data) != 0 {
		t.Fatalf("Expected empty response, got %s", string(data))
	}
}

func TestBytes_NetworkError(t *testing.T) {
	_, err := Bytes("http://nonexistent.url", mockDisplay)
	if err == nil {
		t.Fatalf("Expected network error, got none")
	}
}

func TestBytes_InvalidScheme(t *testing.T) {
	_, err := Bytes("ftp://example.com", mockDisplay)
	if err == nil {
		t.Fatalf("Expected error for invalid URL scheme, got none")
	}
}

func TestBytes_Timeout(t *testing.T) {
	// Simulate a timeout
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	server.Config.WriteTimeout = 1 // set a short timeout for the test
	server.Start()
	defer server.Close()

	_, err := Bytes(server.URL, mockDisplay)
	if err == nil {
		t.Fatalf("Expected timeout error, got none")
	}
}

func TestBytes_DisplayFunctionCalled(t *testing.T) {
	displayCalled := false
	displayFunc := func(msg string) {
		displayCalled = true
	}

	server := startMockServer(http.StatusOK, "mock response")
	defer server.Close()

	_, err := Bytes(server.URL, displayFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !displayCalled {
		t.Fatalf("Expected display function to be called, but it was not")
	}
}

func TestBytes_ResponseBodyReadError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		server.CloseClientConnections() // Close the connection to simulate read error
	}))
	defer server.Close()

	_, err := Bytes(server.URL, mockDisplay)
	if err == nil {
		t.Fatalf("Expected read error, got none")
	}
}