package htmlquery

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

// Mock server to handle HTTP requests for testing
func mockServer(responseBody string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	})
	return httptest.NewServer(handler)
}

// Test for Request function
func TestRequest_Success(t *testing.T) {
	server := mockServer("<html><body><div id='test'>content</div></body></html>", http.StatusOK)
	defer server.Close()

	extractor := SelectionExtractor("#text")
	result, err := Request(server.URL, "div#test", extractor)

	assert.NoError(t, err)
	assert.Equal(t, []string{"content"}, result)
}

func TestRequest_InvalidURL(t *testing.T) {
	_, err := Request("http://invalid-url", "div#test", SelectionExtractor("#text"))

	assert.Error(t, err)
}

func TestRequest_InvalidSelector(t *testing.T) {
	server := mockServer("<html><body><div id='test'>content</div></body></html>", http.StatusOK)
	defer server.Close()

	extractor := SelectionExtractor("#text")
	result, err := Request(server.URL, "invalid#selector", extractor)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestRequest_EmptyResponseBody(t *testing.T) {
	server := mockServer("", http.StatusOK)
	defer server.Close()

	extractor := SelectionExtractor("#text")
	result, err := Request(server.URL, "div#test", extractor)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestRequest_Non200StatusCode(t *testing.T) {
	server := mockServer("<html><body><div id='test'>content</div></body></html>", http.StatusInternalServerError)
	defer server.Close()

	extractor := SelectionExtractor("#text")
	_, err := Request(server.URL, "div#test", extractor)

	assert.Error(t, err)
}

// Test for SelectionExtractor function
func TestSelectionExtractor_TextExtraction(t *testing.T) {
	selection := goquery.NewDocumentFromReader(strings.NewReader("<div id='test'>text content</div>")).Find("div#test")
	extractor := SelectionExtractor("#text")

	result := extractor(selection)

	assert.Equal(t, "text content", result)
}

func TestSelectionExtractor_AttributeExtraction(t *testing.T) {
	selection := goquery.NewDocumentFromReader(strings.NewReader("<div id='test' data-attr='attribute content'></div>")).Find("div#test")
	extractor := SelectionExtractor("data-attr")

	result := extractor(selection)

	assert.Equal(t, "attribute content", result)
}

func TestSelectionExtractor_NonExistentAttribute(t *testing.T) {
	selection := goquery.NewDocumentFromReader(strings.NewReader("<div id='test'></div>")).Find("div#test")
	extractor := SelectionExtractor("data-attr")

	result := extractor(selection)

	assert.Equal(t, "", result)
}

func TestSelectionExtractor_EmptyAttribute(t *testing.T) {
	selection := goquery.NewDocumentFromReader(strings.NewReader("<div id='test' data-attr=''></div>")).Find("div#test")
	extractor := SelectionExtractor("data-attr")

	result := extractor(selection)

	assert.Equal(t, "", result)
}

func TestSelectionExtractor_WhitespaceAttribute(t *testing.T) {
	selection := goquery.NewDocumentFromReader(strings.NewReader("<div id='test' data-attr='  '></div>")).Find("div#test")
	extractor := SelectionExtractor("data-attr")

	result := extractor(selection)

	assert.Equal(t, "", result)
}