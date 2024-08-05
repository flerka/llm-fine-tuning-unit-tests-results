package internal

import (
	"net/http"
	"testing"
)

func TestMatches_AllHeadersMatch(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-http-client")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept", "User-Agent"},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "application/json")
	responseHeader.Set("User-Agent", "Go-http-client")

	if !v.Matches(responseHeader) {
		t.Errorf("expected headers to match")
	}
}

func TestMatches_HeaderMismatch(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-http-client")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept", "User-Agent"},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "application/json")
	responseHeader.Set("User-Agent", "Mozilla")

	if v.Matches(responseHeader) {
		t.Errorf("expected headers to mismatch")
	}
}

func TestMatches_SingleHeaderMatch(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept"},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "application/json")

	if !v.Matches(responseHeader) {
		t.Errorf("expected headers to match")
	}
}

func TestMatches_NoHeadersToMatch(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)

	v := Variant{
		r:           req,
		headerNames: []string{},
	}

	responseHeader := http.Header{}

	if !v.Matches(responseHeader) {
		t.Errorf("expected headers to match when no headers specified")
	}
}

func TestMatches_EmptyRequestHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept"},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "application/json")

	if v.Matches(responseHeader) {
		t.Errorf("expected headers to mismatch when request header is empty")
	}
}

func TestMatches_EmptyResponseHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept"},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "")

	if v.Matches(responseHeader) {
		t.Errorf("expected headers to mismatch when response header is empty")
	}
}

func TestMatches_MissingResponseHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept"},
	}

	responseHeader := http.Header{}

	if v.Matches(responseHeader) {
		t.Errorf("expected headers to mismatch when response header is missing")
	}
}

func TestMatches_ExtraResponseHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept"},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "application/json")
	responseHeader.Set("User-Agent", "Go-http-client")

	if !v.Matches(responseHeader) {
		t.Errorf("expected headers to match even with extra response header")
	}
}

func TestMatches_EmptyHeaderNames(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")

	v := Variant{
		r:           req,
		headerNames: []string{},
	}

	responseHeader := http.Header{}
	responseHeader.Set("Accept", "application/json")

	if !v.Matches(responseHeader) {
		t.Errorf("expected headers to match when no header names are specified")
	}
}