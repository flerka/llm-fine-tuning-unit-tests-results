package htmlquery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// Mock server handler
func mockServer(response string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	})
}

// Test for a successful response
func TestRequest_Success(t *testing.T) {
	mockResp := `<html><body><div class="item">Item 1</div><div class="item">Item 2</div></body></html>`
	server := httptest.NewServer(mockServer(mockResp))
	defer server.Close()

	extractor := selectionTextExtractor
	results, err := Request(server.URL, ".item", extractor)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{"Item 1", "Item 2"}
	if len(results) != len(expected) {
		t.Fatalf("Expected %v results, got %v", len(expected), len(results))
	}

	for i, v := range results {
		if v != expected[i] {
			t.Errorf("Expected result %v, got %v", expected[i], v)
		}
	}
}

// Test for a 404 response
func TestRequest_404Response(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	extractor := selectionTextExtractor
	_, err := Request(server.URL, ".item", extractor)

	if err == nil {
		t.Fatalf("Expected error for 404 response, got nil")
	}
}

// Test for an invalid URL
func TestRequest_InvalidURL(t *testing.T) {
	invalidURL := "http://invalid.url"
	extractor := selectionTextExtractor
	_, err := Request(invalidURL, ".item", extractor)

	if err == nil {
		t.Fatalf("Expected error for invalid URL, got nil")
	}
}

// Test for no matching elements
func TestRequest_NoMatchingElements(t *testing.T) {
	mockResp := `<html><body><div class="not-item">Item 1</div><div class="not-item">Item 2</div></body></html>`
	server := httptest.NewServer(mockServer(mockResp))
	defer server.Close()

	extractor := selectionTextExtractor
	results, err := Request(server.URL, ".item", extractor)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(results) != 0 {
		t.Fatalf("Expected 0 results, got %v", len(results))
	}
}

// Test for empty response body
func TestRequest_EmptyResponseBody(t *testing.T) {
	mockResp := ``
	server := httptest.NewServer(mockServer(mockResp))
	defer server.Close()

	extractor := selectionTextExtractor
	results, err := Request(server.URL, ".item", extractor)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(results) != 0 {
		t.Fatalf("Expected 0 results, got %v", len(results))
	}
}

// Test for malformed HTML
func TestRequest_MalformedHTML(t *testing.T) {
	mockResp := `<html><body><div class="item">Item 1<div><div class="item">Item 2</body></html>`
	server := httptest.NewServer(mockServer(mockResp))
	defer server.Close()

	extractor := selectionTextExtractor
	results, err := Request(server.URL, ".item", extractor)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{"Item 1", "Item 2"}
	if len(results) != len(expected) {
		t.Fatalf("Expected %v results, got %v", len(expected), len(results))
	}

	for i, v := range results {
		if v != expected[i] {
			t.Errorf("Expected result %v, got %v", expected[i], v)
		}
	}
}