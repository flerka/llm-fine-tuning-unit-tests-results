package internal

import (
	"net/http"
	"testing"
)

func TestNewVariant(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	v := NewVariant(req)
	if v == nil {
		t.Errorf("NewVariant returned nil")
	}
	if v.r != req {
		t.Errorf("NewVariant did not set request correctly")
	}
}

func TestSetResponseHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	v := NewVariant(req)

	header := http.Header{}
	header.Set("Vary", "User-Agent, Accept-Encoding")
	v.SetResponseHeader(header)

	expectedHeaders := []string{"Accept-Encoding", "User-Agent"}
	for i, name := range expectedHeaders {
		if v.headerNames[i] != name {
			t.Errorf("SetResponseHeader failed, expected %s, got %s", name, v.headerNames[i])
		}
	}
}

func TestCacheKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("User-Agent", "Go-http-client/1.1")
	req.Header.Set("Accept-Encoding", "gzip")

	v := NewVariant(req)
	v.headerNames = []string{"User-Agent", "Accept-Encoding"}

	expectedKey := v.CacheKey()
	if expectedKey == 0 {
		t.Errorf("CacheKey returned zero value")
	}
}

func TestMatches(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("User-Agent", "Go-http-client/1.1")
	req.Header.Set("Accept-Encoding", "gzip")

	v := NewVariant(req)
	v.headerNames = []string{"User-Agent", "Accept-Encoding"}

	responseHeader := http.Header{}
	responseHeader.Set("User-Agent", "Go-http-client/1.1")
	responseHeader.Set("Accept-Encoding", "gzip")

	if !v.Matches(responseHeader) {
		t.Errorf("Matches should return true for matching headers")
	}

	responseHeader.Set("Accept-Encoding", "deflate")
	if v.Matches(responseHeader) {
		t.Errorf("Matches should return false for non-matching headers")
	}
}

func TestVariantHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("User-Agent", "Go-http-client/1.1")
	req.Header.Set("Accept-Encoding", "gzip")

	v := NewVariant(req)
	v.headerNames = []string{"User-Agent", "Accept-Encoding"}

	expectedHeader := http.Header{}
	expectedHeader.Set("User-Agent", "Go-http-client/1.1")
	expectedHeader.Set("Accept-Encoding", "gzip")

	resultHeader := v.VariantHeader()
	for k, v := range expectedHeader {
		if resultHeader.Get(k) != v[0] {
			t.Errorf("VariantHeader mismatch, expected %s for header %s, got %s", v[0], k, resultHeader.Get(k))
		}
	}
}