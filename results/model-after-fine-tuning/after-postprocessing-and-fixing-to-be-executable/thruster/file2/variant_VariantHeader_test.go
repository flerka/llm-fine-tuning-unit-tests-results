package internal

import (
	"testing"
	"net/http"
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
		headerNames:  []string{"Content-Type", "Accept"},
	}
	assert.True(t, variant.Matches( http.Header{}))
}

func TestValidate_EmptyRequest(t *testing.T) {
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
	assert.True(t,  variant.Matches( http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}}))
}

func TestValidate_EmptyRequest_2(t *testing.T) {
	r := &http.Request{Header: http.Header{}}

	variant := Variant{
		r:           r,

		headerNames: []string{"Content-Type", "Accept"},
	}
	assert.True(t,  variant.Matches( http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}}))
}

func TestValidate_EmptyRequest_4(t *testing.T) {
	r := &http.Request{Header: http.Header{"Content-Type": []string{"application/json"}}}

	variant := Variant{
		r:           r,
		headerNames: []string{"Content-Type", "Accept"},
	}
	assert.True(t,  variant.Matches( http.Header{"Content-Type": []string{""}, "Accept": []string{""}}))
}

func TestValidate_EmptyRequest_5(t *testing.T) {
	r := &http.Request{Header: map[string][]string{"Content-Type": []string{"application/json"}}}

	variant := Variant{
		r:           r,
		headerNames: []string{},
	}
	assert.True(t,  variant.Matches(http.Header{}))
}