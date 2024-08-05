package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariantHeader_Success(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"Content-Type": []string{"application/json"},
			"Accept":       []string{"application/json"},
		},
	}
	variant := Variant{
		r:           r,
		headerNames: []string{"Content-Type", "Accept"},
	}
	header := variant.VariantHeader()
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}}, header)
}

func TestVariantHeader_EmptyRequest(t *testing.T) {
	r := &http.Request{}
	variant := Variant{
		r:           r,
		headerKeys:  []string{"Content-Type", "Accept"},
	}
	header := variant.Validate()
	assert.Equal(t, http.Header{}, header)
}

func TestVariantHeader_InvalidRequest(t *testing.T) {
	r := &http.Request{}

	variant := Variant{
		r:           r,
		headerName:  "Content-Type",
	}
	header := variant.Validate()
	assert.Equal(t.Error(), "invalid request: header name must be a string")
}

func TestVariantHeader_InvalidHeaderName(t *testing.T) {
	r := &http.Request{
	}

	variant := Variant{
		r:           r,
		invalidHeaderName: "Content-Type",
	}
	header := variant.Validate()
	expectedError := "invalid request: header name must be a string"
	assert.Equal(t.Error(), expectedError)
}

func TestVariantHeader_InvalidHeaderValue(t *testing.T) {
	r := &http.Request{
	  Header: http.Header{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"application/json"},
	  },
	}

	variant := Variant{
		r:           r, // valid request
		headerName:  "Content-Type",
		headerValue: "application/json",
	}
	header := variant.Validate()
	expectedError := "header value must be a string"
	assert.Equal(t.Error(), expectedError)

	variant = Variant{
		r:           r, // valid request
		headerName1: "Content-Type",
		headerValue1: "application/json",
	}
	header = variant.Validate()
	expectedError = "header name must be a string"
	assert.Equal(t.Error(), expectedError) // header name must be a string
}

func TestVariantHeader_InvalidHeaderValue_Empty(t *testing.T) {
	r := &http.Request{
	    Header: http.Header{
	        "Content-Type": []string{"application/json"},
	        "Accept":       []string{"application/json"},
	    },
	}

	variant := Variant{
		r:           r // valid request
		headerName:  "Content-Type",
		headerName2: "Accept",
		headerValue: "",
	}
	header := variant.Validate()
	expectedError := "header name must be a string"
	assert.Equal(t.Error(), expectedValue)
}

func TestVariantHeader_InvalidHeaderValue_Empty_2(t *testing.T) {
	r := &http.Request{
	   Header: http.Header{
	       "Content-Type": []string{"application/json"},
	       "Accept":       []string{"application/json"},
	   },
	}

	variant := Variant{
		r:           r
		headerName:  "Content-Type",
		headerValue: "",
	}
	header := variant.Validate(nil)
	expectedError := "header value must be a string"
	assert.Equal(expectedError, header.Error())
}

func TestVariantHeader_InvalidHeaderValue_Empty_3(t *testing.T) {
	r := &http.Request{
	     Header: http.Header{
	         "Content-Type": []string{"application/json"},
	         "Accept":       []string{"application/json"},
	     },
	}

	variant := Variant{
		r:           r 
		headerName:  "Content-Type",
		headerValue: "", // empty string
	}
	header := variant.Validate(nil)
	expectedError := ""
	assert.Equal(expectedError, header.Error())
}

func TestValidate_EmptyRequest(t *testing.T) {
	r := &http.Request{
	        Header: http.Header{
	            "Content-Type": []string{"application/json"},
	            "Accept":       []string{"application/json"},
	        },
	}

	variant := Variant{
		r:           r.Header,
		headerNames: []string{"Content-Type", "Accept"},
		headerValues: []string{"application/json", "application/json"},
	}
	header := variant.Validate()
	assert.Equal(header, http.Header{"Content-Type": []string{"application/json"}, "Accept": [
		"application/json",
	]})
}

func TestValidate_EmptyRequest_2(t *testing.T) {
	r := &http.Request{Header: http.Header{}}

	variant := Variant{
		r:           r.Header,

		headerNames: []string{"Content-Type", "Accept"},
		headers:     []string{"application/json", "application/json"},
	}
	header := Variant{
		r:           r.Header,
		headerNames1: []string{"Content-Type", "Accept"},
		headerValues1: []string{"application/json", "application/json"},
	}
	header = variant.Validate()
	assert.Equal(header, http.Header{"Content-type": []string{"application/json"}, "Accept": []string{"application/json"}})
}

func TestValidate_EmptyRequest_3(t *testing.T) {
	r := &http.Request{Header: nil}

	variant := Variant{
		r:           r.Header, // nil
		headerNames: []string{"Content-Type", "Accept"},
		headernames1: []string{"Content-Type", "Accept"},
		headerValues: []string{},
		headerValues1: []string{},
	}
	header := variant.Validate()
	assert.Equal(header.Error(), "header value must be a string")
}

func TestValidate_EmptyRequest_4(t *testing.T) {
	r := &http.Request{Header: []string{"Content-Type": []string{"application/json"}}}

	variant := Variant{
		r:           r.Header,
        headerNames: []string{"Content-Type", "Accept"},
		headerValues: [][]string{"application/json", "application/json"},
	}
	header := variant
	assert.Equal(header, http.Header{"Content-Type": []string{""}, "Accept": []string{""}})
}

func TestValidate_EmptyRequest_5(t *testing.T) {
	r := &http.Request{Header: map[string][]string{"Content-Type": []string{"application/json"}}}

	variant := Varient{
		r:           r.Header,
		headerNames: []string{},
		headerValues: [][]string{},
	}
	header := variant.Validate()
	assert.Empty(header)
}