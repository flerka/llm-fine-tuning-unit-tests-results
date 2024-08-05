package internal

import (
	"net/http"
	"testing"
	"net/url"
	"github.com/stretchr/testify/assert"
)

func TestNewVariant(t *testing.T) {
	r := &http.Request{
		Method: "GET",
		URL:  &url.URL{Scheme: "http://", Host: "example.com", Path: "/foo"},
	}
	v := NewVariant(r)
	assert.NotNil(t, v)
}

func TestVariantHeader(t *testing.T) {
	r := &http.Request{
	}
	v := NewVariant(r)
	assert.Empty(t, v.VariantHeader())
}

func TestSetResponseHeader(t *testing.T) {
	r := &http.Request{

	}
	v := NewVariant(r)
	v.SetResponseHeader(http.Header{
		"Vary": []string{"Accept-Encoding"},
	})
	assert.Equal(t, []string{"Accept-Encoding"}, v.headerNames)
}

func TestMatches(t *testing.T) {
	r := &http.Request{
	        Method: "GET",
			URL:  &url.URL{Scheme: "http://", Host: "example.com", Path: "/foo"},
	        Header: http.Header{
	            "Accept-Encoding": []string{"gzip"},
	        },
	}
	v := NewVariant(r)
	v.SetResponseHeader(http.Header{
"Accept-Encoding": []string{ "gzip"}})
	assert.True(t, v.Matches(http.Header{
		"Accept-Encoding": []string{"gzip"},
	}))
	assert.False(t, v.Matches(http.Header{
		"Accept-Encoding": []string{"gzip"},
	}))
}

func TestVariantHeader_Vary(t *testing.T) {
	v := NewVariant(&http.Request{})
	v.SetResponseHeader(http.Header{
"Vary": []string{"Accept-Encoding"}})
	assert.Equal(t, http.Header{
		"Vary": []string{"Accept-Encoding"},
	}, v.VariantHeader())
}

func TestVariantHeader_Vary_Multiple(t *testing.T) {
	v := NewVariant(&http.Request{Header: http.Header{"Vary": []string{"Accept-Encoding"}}})
	v.SetResponseHeader(http.Header{
		"Vary": []string{"Accept-Encoding, Content-Encoding"}})
	assert.Equal(t, http.Header{
		"Accept-Encoding": []string{"Accept-Encoding"},
		"Content-Encoding": []string{"Content-Encoding"},
	}, v.VariantHeader())
}

func TestVariants_Vary_Empty(t *testing.T) {
	v := NewVariant(&http.Request{})
	v.SetResponseHeader(http.Header{
		"Vary": []string{""}})
	assert.Equal(t, http.Header{}, v.VariantHeader()) // empty header
}

func TestVariants_Vary_Empty_Multiple(t *testing.T) {
	v := NewVariant(&
		http.Request{
			Header: http.Header{
				"Vary": []string{"Accept-Encoding"},
			},
		})
		v.SetResponseHeader(http.Header{
			"Vary": []string{""}})
		assert.Equal(t,
		http.Header{
			"Accept-Encoding": []string{"Accept-Encoding"},
		},
		v.VariantHeader())
}

func TestVariants_Vary_Empty_Multiple_Duplicate(t *testing.T) {
	v := NewVariant(&http.Request{
			Header: http.Header{
				"Vary": []string{"Accept-Encoding"},
			},
		})
	v2 := NewVariant(&http.Request{
			Header: http.Header{
				"Vary": []string{"Accept-Encoding"},
			},
		})
	v.SetResponseHeader(http.Header{
			"Vary": []string{"Accept-Encoding"}})
	v2.SetResponseHeader(http.Header{
				"Vary": []string{"Accept-Encoding"}})
		assert.Equal(t, http.Header{
			"Accept-Encoding": []string{"Accept-Encoding"}, // v1
		},
		v.VariantHeader())
	assert.Equal(t, http.Header{
			"Vary": []string{"Accept-Encoding"}, // v2
		},
		v2.VariantHeader())
}

func TestVariants_Vary_Empty_Duplicate(t *testing.T) {
	v := NewVariant(
		&http.Request{
			Header: http.Header{
				"Vary": []string{"Accept-Encoding"},	
			},
		})
	assert.Equal(t,
		http.Header{
			"Vary": []string{"Accept-Encoding"},
		},
		v.VariantHeader(),
	)
}

func TestVariants_Vary_Empty_Duplicate_Multiple(t *testing.T) {
	v := NewVariant(
			&http.Request{
				Header: http.Header{
					"Vary": []string{"Accept-Encoding"},
				},
			})
	assert.Equal(t,
		http.Header{
			"Vary": []string{"Accept-Encoding"},
		},

		v.VariantHeader(),
	)
}

func TestVariant_Vary_Empty(t *testing.T) {
	v := NewVariant(&http.Request{})
	v.SetResponseHeader(http.Header{
		"Vary": []string{""}})
assert.Equal(t, http.Header{}, v.VariantHeader(),
	)
}

func TestMatches_Vary_Empty(t *testing.T) {
	v := NewVariant(&http.Request{})
	v.SetResponseHeader(http.Header{
		"Vary": []string{"Accept-Encoding"}})
	assert.True(t, v.Matches(http.Header{
		"Accept-Encoding": []string{"gzip"},
	}))
	v.SetResponseHeader(http.Header{
		"Vary": []string{""}})
	assert.False(t, v.Matches(http.Header{
		"Accept-Encodig": []string{"gzip"},
	}))
}