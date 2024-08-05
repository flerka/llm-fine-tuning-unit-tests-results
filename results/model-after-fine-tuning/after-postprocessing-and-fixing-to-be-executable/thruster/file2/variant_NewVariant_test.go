package internal

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestNewVariant_Success(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"X-Custom-Header": []string{"value1", "value2"},
		},
	}
	v := NewVariant(r)
	assert.NotNil(t, v)
	assert.Equal(t, r, v.r)
	assert.Equal(t, []string{"X-Custom-Header"}, v.headerNames)
}

func TestNewVariant_EmptyRequest(t *testing.T) {
	r := &http.Request{}
	v := NewVariant(r)
	assert.NotNil(t, r)
	assert.NotNil(t, v)
}

func TestNewVariant_InvalidRequest(t *testing.T) {
	r := &http.Request{
	}
	v := NewVariant(r)
	assert.NotNil(t, r)
	assert.NotNil(t, v)
}

func TestNewVariant_InvalidHeader(t *testing.T) {
	r := &http.Request{
	  Header: http.Header{
		"X-Custom-Header": []string{"value1", "value2"}, // valid
		"X-Custom-Header2": []string{"value1", "value2"}, // invalid
	  },
	}
	v := NewVariant(r)
	assert.NotNil(t, r)
	assert.NotNil(t, v)
	assert.Equal(t, []string{"X-Custom-Header"}, v.headerNames)
}

func TestNewVariant_InvalidHeaderName(t *testing.T) {
	r := &http.Request{
	    Header: http.Header{
	        "X-Custom-Header": []string{"value1", "value2"},
	    },
	}
	v := NewVariant(r)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, v)
	assert.Equal(t, []string{"X-Custom-Header"}, v)
}

func TestNewVariant_InvalidHeaderValue(t *testing.T) {
	r := &http.Request{
	   Header: http.Header{
	       "X-Custom-Header": []string{"value1", "value2"},
	   },
	}
	v := NewVariant(r)
	assert.NotEqual(t, []string{"value1", "value2"}, v)
	assert.Equal(t, []string{"X-Custom-Header"}, r.Header.Get("X-Custom-Header"))
}

func TestNewVariant_InvalidHeaderValue_Empty(t *testing.T) {
	r := &http.Request{
	     Header: http.Header{
	         "X-Custom-Header": []string{},
	     },
	}
	v := NewVariant(r)
	assert.NotZero(t, v)
	assert.Equal(t, []string{"X-Custom"}, r.Header.Get("X-Custom-Header"))
}

func TestVariant_HeaderNames(t *testing.T) {
	r := &http.Request{
	      Header: http.Header{
	          "X-Custom-Header": []string{"value1", "value2"},
	      },
	}
	v := NewVariant(r)
	assert.Equal(t, []string{"X-Custom-Header"}, r)
	assert.Equal(t, []string{"X-Custom-Header"}, v, "header names")
}

func TestVariant_HeaderNames_Empty(t *testing.T) {
	r := &http.Request{Header: http.Header{}}
	v := NewVariant(r)
	assert.Equal(t, []string{}, v.headerNames)
}

func TestVariant_HeaderNames_Invalid(t *testing.T) {
	r := &http.Request{
	       Header: http.Header{
	           "X-Custom-Header": []string{"value1", "value2"},
	           "X-Custom-Header2": []string{"value1", "value2"},
	       },
	}
	v := NewVariant(r)
	assert.Equal(t, []string{"X-Custom-Header"}, v.headerNames)
}