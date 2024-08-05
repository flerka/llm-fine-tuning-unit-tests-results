package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewVariantBasic tests the basic creation of a Variant object
func TestNewVariantBasic(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	variant := NewVariant(req)
	if variant.r != req {
		t.Errorf("Expected request to be %v, got %v", req, variant.r)
	}
}

// TestNewVariantNilRequest tests the creation of a Variant object with a nil request
func TestNewVariantNilRequest(t *testing.T) {
	variant := NewVariant(nil)
	if variant.r != nil {
		t.Errorf("Expected request to be nil, got %v", variant.r)
	}
}

// TestNewVariantWithHeaders tests the creation of a Variant object and verifies headerNames
func TestNewVariantWithHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Add("Test-Header", "value")
	variant := NewVariant(req)
	if len(variant.headerNames) != 0 {
		t.Errorf("Expected headerNames to be empty, got %v", variant.headerNames)
	}
}

// TestNewVariantEdgeCase tests the creation of a Variant object with an edge case URL
func TestNewVariantEdgeCase(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/edgecase?param=value", nil)
	variant := NewVariant(req)
	if variant.r.URL.String() != "http://example.com/edgecase?param=value" {
		t.Errorf("Expected URL to be %v, got %v", "http://example.com/edgecase?param=value", variant.r.URL.String())
	}
}

// TestNewVariantWithPostRequest tests the creation of a Variant object with a POST request
func TestNewVariantWithPostRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader("body"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	variant := NewVariant(req)
	if variant.r.Method != "POST" {
		t.Errorf("Expected method to be POST, got %v", variant.r.Method)
	}
	if variant.r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type to be application/x-www-form-urlencoded, got %v", variant.r.Header.Get("Content-Type"))
	}
}

// TestNewVariantEmptyHeaderNames tests the creation of a Variant object and verifies the headerNames slice is initially empty
func TestNewVariantEmptyHeaderNames(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	variant := NewVariant(req)
	if len(variant.headerNames) != 0 {
		t.Errorf("Expected headerNames to be empty, got %v", variant.headerNames)
	}
}