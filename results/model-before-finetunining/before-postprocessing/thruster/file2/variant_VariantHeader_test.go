package internal

import (
    "testing"
    "net/http"
)

type Variants struct {
    VariantHeaderFunc func(*http.Request) http.Header
}

func (v Variants) VariantHeader(req *http.Request) http.Header {
    if v.VariantHeaderFunc == nil {
        return http.Header{}
    }
    return v.VariantHeaderFunc(req)
}

func TestVariantHeader_Success(t *testing.T) {
    req := &http.Request{
        Method: "GET",
        Header: http.Header{
            "X-Custom-Header": []string{"value1"},
        },
    }
    variants := Variants{
        VariantHeaderFunc: func(req *http.Request) *http.Header {
            return req.Header.Get("X-Custom-Header")
        },
    }

    header := variants.VariantHeader(req)
    if len(header) != 1 || header["X-Custom-Header"][0] != "value1" {
        t.Errorf("Expected header value to be 'value1', got '%v'", header["X-Custom0-Header"][0])
    }
}

func TestVariants_EmptyRequest(t *testing.T) {
    req := &http.MockRequest{Method: "GET", Header: http.Header{}}
    variants := Variants{VariantHeaderFunc: func(req http.Request) http.Header {}}

    header := variants.Variants(req).VariantHeader(req)
    if len(header) > 0 {
        t.Errorf("Header should be empty for an empty request")
    }
}

// Add more tests for different scenarios...