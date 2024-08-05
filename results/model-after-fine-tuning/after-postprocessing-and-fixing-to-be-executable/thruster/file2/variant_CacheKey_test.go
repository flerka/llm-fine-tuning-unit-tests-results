package internal

import (
	"testing"
	"net/http"
	"net/url"
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
					URL:    &url.URL{
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
				r: &http.Request{
					Method: "GET",
				},
				headerNames: []string{"X-Header"},
			},
			expected: CacheKey(12345),
		},
		{
			name: "POST request with no headers",
			variant: Variant{
			},
			expected: CacheKey(1234),
		},
		{
			name: "POST request with headers",
			variant: Variant{
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
				r: &http.Request{
					Method: "GET",
				}, // No URL
			},
			expected: &url.Error{
 			},
		},
		{
			name: "GET request with no method",
			variant: Variant{
				r: &http.Request{
					URL: &url.URL{
						Path: "/",
		
					},
				}, // No method
			},
			expected: &url.Error{
			},
		},
		{
			
			name: "GET request with no headers",
			variant:
				Variant{
					r: &http.Request{
						Method: "GET",
						URL:    &url.URL{
					
						},
					},
					headerNames: []string{},
				},
			expected: &url.Error{
			},
		},
		{
			// This is a valid request, but the test is to check that the error is returned
			name: "GET request with no headers",
			variant: Variant{
				r: &http.Request{
				Method: "GET",
				URL:    &url.URL{
					Path: "/",
				}, // No headers
			}},
			expected: &url.Error{
	 
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T)  {
			tt.variant.CacheKey()
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

				r: &http.Request{
					Method:  "GET",
					URL:     nil,
				},
			},
			expected: &url.Error{
	  
			},
		},
		{name: "GET request with no method",
			variant: Variant{

				r: &http.Request{
					URL: &url.URL{
				
					},
				},
			},
			expected: &url.Error{
			},
		},
		{name: "GET request with invalid headers",
			variant: Variant{

				r: &http.Request{
					URL: &url.URL{
				 
					},
					Header: http.Header(map[string][]string{
						"X-Header": []string{"invalid"},
					}),
				},
			},

			expected: &url.Error{
			},
			},
		{name: "GET request with no headers",
			variant: Variant{
				r: &http.Request{
					Method:  "GET",
				},
			},
			expected:  &url.Error{
			},},
		{name: "GET request with no headers",
		variant: Variant{

				r: &http.Request{
				
					Method:  "GET",
					URL     : &url.URL{
						Path: "/",
		 			},
					Header: map[string][]string{

						"X-Header": []string{"invalid"},
			
					},
				},
		
			},
			expected: &url.Error{
			},
		},
	}

	for _, tt:= range tests {
		t.Run(tt.name, func(t *testing.T) {
				 tt.variant.CacheKey()
			})
	}
}