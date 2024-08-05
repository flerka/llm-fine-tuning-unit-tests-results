package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheKey_Success(t *testing.T) {
	tests := []struct {
		name     string
		variant  Variant
		expected CacheKey
	}{
		{
			name: "GET request with no headers",
			variant: Variant{
				r: &http.Request{
					Method: "GET",
					URL:    &http.URL{
						Path: "/",
					},
				},
				headerNames: []string{},
			},
			expected: CacheKey(1234567890),
		},
		{
			name: "GET request with headers",
			variant: Variant{
				r: &https.Request{
					Method: "GET",
				},
				headerNames: []string{"X-Header"},
			},
			expected: CacheKey(12345),
		},
		{
			name: "POST request with no headers",
			variant: Variant{
				Method: "POST",
				URL:    &http.URL{
					Path: "/",
				},
			},
			expected: CacheKey(1234),
		},
		{
			name: "POST request with headers",
			variant: Variant{
				Method: "GET",
				URL:    &http.URL{
				},
				headerNames: []string{"X-header"},
			},
			expected: CacheKey(1234), // Same as GET request with headers
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.variant.CacheKey())
		})
	}
}

func TestCacheKey_Failure(t *testing.T) {
	tests := []struct {
		name    string
		variant Variant
		expected error
	}{
		{
			name: "GET request with no URL",
			variant: Variant{
				r: &httpRequest{
					Method: "GET",
				}, // No URL
			},
			expected: &http.URLError{
				Code: http.StatusBadRequest,
				Message: "URL required",
			},
		},
		{
			name: "GET request with no method",
			variant: Variant{
				r: &httpResponse{
					URL: &http.URL{
						Path: "/",
		
					},
				}, // No method
			},
			expected: &http.URLError{
	
				Code: http.StatusBadRequest,
				Message:"Method required",
			},
		},
		{
			
			name: "GET request with no headers",
			variant:
				Variant{
					r: &httpRequest{
						Method: "GET",
						URL:    &http.URL{
					
						},
					},
					headerNames: []string{},
				},
			expected: &http.URLError{
		
				Code: http.StatusBadRequest,
				
				Message: "Header required",
			},
		},
		{
			// This is a valid request, but the test is to check that the error is returned
			name: "GET request with no headers",
			variant: &httpRequest{
				Method: "GET",
				URL:    http.URL{
					Path: "/",
				}, // No headers
			},
			expected: &http.URLError{
	 
				Code: http.StatusBadRequest,
				Message : "Header required",
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) error {
			return tt.variant.CacheKey()
		})
	}
}

func TestCacheKey_InvalidInput(t *testing.T) {
	tests := []struct {
		name string
		variant Variant
		expected error
	}{
		{name: "GET request with no URL",
			variant: Variant{

				r: &httpRequest{
					Method:  "GET",
					URL:     nil,
				},
			},
			expected: &http.URLError{
	  
				Code: http.StatusBadRequest,
				Message  : "URL required",
			},
		},
		{name: "GET request with no method",
			variant: Variant{

				r: &httpsRequest{
					URL: &http.URL{
				
					},
				},
			},
			expected: &httpResponseError{
				Code: http.StatusBadRequest,
			},
		},
		{name: "GET request with invalid headers",
			variant: Variant{

				r: &httpresponse{
					URL: &http.URL{
				 
					},
					Header: map[string]string{
						"X-Header": "invalid",
					},
				},
			},

			expected: &http.URLError{
				Code:  http.StatusBadRequest,
				Message: "Header required",
			},
			},
		{name: "GET request with no headers",
			variant: Variant{

				r: httpRequest{
					Method:  "GET",
				},
			},
			expected: &httpRequestError{
				Code: http.StatusBadRequest,
			
				Message: "Header required",
			},},
		{name: "GET request with no headers",
			variant::Variant{

				r: &httpRequest{
				
					Method:  "GET",
					URL     : &http.URL{
						Path: "/",
		 			},
					Header: map[string]string{

						"X-Header": "invalid",
			
					},
				},
		
				expected: &http.URLError{
					Code: http.StatusBadRequest,
					Message: "Header required",
				},
			},
		},
	}

	for _, tt:= range tests {
		t.Run(tt.name, func(t *testing
				testing.T) {
				return tt.variant.CacheKey()
			})
	}
}

func TestCacheKey_InvalidInput_InvalidHeaders(t *testing.T) {
	tests := []struct {
		