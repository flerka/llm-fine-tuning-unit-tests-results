package htmlquery

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery" // import goquery
)

func TestSelectionExtractor(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		expected string
	}{
		{
			name:     "extracts text from a single element",
			selector: "#text",
			expected: "text",
		},
		{
			name:     "extracts text from a single element with spaces",
			selector: "#text",
			expected: " text ",
		},
		{
			name:     "extracts text with attributes",
			selector: "#text",
			expected: "text", // no attributes
		},
		{
			name:     "extracts text without attributes",
			selector: "#text",
			expected: "text ", // no attributes
		},
		{
			name:     "returns empty string when no text is found",
			selector: "#text",
			expected: "",
		},
		{
			name:     "returns empty string for invalid selector",
			selector: "#text",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extractor := SelectionExtractor(tt.selector)
			extracted := extractor(nil)
			if extracted != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, extracted)
			}
		})
	}
}

func TestExtractList(t *testing.T) {
	tests := []struct {
		name    string
		selector string
		extractor func(*goquery.Selection) string
		expected []string
	}{
		{
			name:    "extracts text from a single element",
			selector: "#text",

			extractor: selectionTextExtractor,
			expected: []string{"text"},
		},
		{
			name:    "extracts text from a single element with spaces",
			selector: "#text ",

			extractor: selectionTextExtractor,
			expected:  []string{" text "},
		},
		{
			name:    "extracts text with attributes",
			selector: "#text",

			extractor: selectionTextExtractor,
			expected: []string{"text"}, // no attributes
		},
		{
			name:    "extracts text without attributes",
			selector: "#text",

			extractor: SelectionExtractor("#text"),
			expected: []string{"text"}, // no attributes
		},

		{
			name:    "returns empty list when no text is found",
			selector: "#text",

			extractor: selectionTextExtractor,
			expected: []string{},
		},
		{
			name:    "returns empty list for invalid selector",
			selector: "#text",

			extractor: SelectionExtractor("#invalid"),
			expected: []string{},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(`<div id="text">text</div>`))
			})
			go func() {
				http.ListenAndServe(":8080", nil)
			}()

			extracted, err := extractList(b, tt.selector, tt.extractor)
			if err != nil {
				t.Errorf("extractList returned error: %v", err)
			}
			for i, v := range tt.expected {
				if extracted[i] != v {
					t.Errorf("extracted: %v, expected: %v", extracted, tt.expected)
				}
			}
			
		})
	}
}