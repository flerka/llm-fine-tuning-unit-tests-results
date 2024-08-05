package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVariant(t *testing.T) {
	r := &http.Request{
		Method: "GET",
		URL:    "http://example.com/foo",
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

func TestCacheKey(t *testing.T) {
	r := &http.Request{
	    Method: "GET",
	    URL:    "http://example.com/foo",
	    Header: http.Header{
	        "Accept-Encoding": []string{"gzip"},
	    },
	}
	v := NewVariant(r)
	v.SetResponseHeaders(http.Header{
		"Vary": []string{"Accept-Encoding"}, // this header is not used
	})
	assert.Equal(t, fnv.Sum64([]byte("GET") + []byte("http://example.com/foo") + []byte("gzip")), v.CacheKey())
}

func TestMatches(t *testing.T) {
	r := &http.Request{
	        Method: "GET",
	        URL:    "http://example.com/foo",
	        Header: http.Header{
	            "Accept-Encoding": []string{"gzip"},
	        },
	}
	v := NewVariant(r)
	v.SetHeader("Accept-Encoding", "gzip")
	assert.True(t, v.Matches(http.Header{
		"Accept-Encoding": []string{"gzip"},
	}))
	assert.False(t, v.Matches(http.Header{
		"Accept-Encoding: gzip",
	}))
}

func TestVariantHeader(t *testing.T) {
	v := NewVariant(&http.Request{})
	assert.Equal(t, http.Header{}, v.VariantHeader())
}

func TestVariantHeader_Vary(t *testing.T) {
	v := NewVariant(&http.Request{}
	)
	v.SetHeader("Vary", "Accept-Encoding")
	assert.Equal(t, http.Header{
		"Vary": []string{"Accept-Encoding"},
	}, v.VariantHeader())
}

func TestVariantHeader_Vary_Multiple(t *testing.T) {
	v := NewVariant(&http.Request{"Vary": []string{"Accept-Encoding"}})
	v.SetHeader("Vary", "Accept-Encoding, Content-Encoding")
	assert.Equal(t, http.Header{
		"Accept-Encoding": []string{"Accept-Encoding"},
		"Content-Encoding": []string{"Content-Encoding"},
	}, v.VariantHeader())
}

func TestVariants_Vary_Empty(t *testing.T) {
	v := NewVariant(&http.Request)
	v.SetHeader("Vary", "")
	assert.Equal(t, http.Header{}, v.VariantHeader()) // empty header
}

func TestVariants_Vary_Empty_Multiple(t *testing.T) {
	v := NewVariant(&
		http.Request{
			Header: http.Header{
				"Vary": []string{"Accept-Encoding"},
			},
		})
	v.SetHeader("Vary", "")
	assert.Equal(t,
		http.Header{
			"Accept-Encoding": []string{"Accept-Encoding"},
		},
		v.VariantHeader())
}

func TestVariants_Vary_Empty_Multiple_Duplicate(t *testing.T) {
	v := NewVariant(&
		&http.Request{
			Header: http.Header{
				Vary: []string{"Accept-Encoding"},
			},
		})
	v2 := NewVariant(&
		&http.Request{
			Header: &http.Header{
				Vary: []string{"Accept-Encoding"},

			},
		})
	v.SetHeader("Vary", "Accept-Encoding")
	v2.SetHeader("Vary", "Accept-Encoding")
	assert.Equal(t
		, http.Header{
			"Accept-Encoding": []string{"Accept-Encoding"}, // v1
		},
		v.VariantHeader())
	assert.Equal(t
		, http.Header{
			Vary: []string{"Accept-Encoding"}, // v2
		},
		v2.VariantHeader())
}

func TestVariants_Vary_Empty_Duplicate(t *testing.T) {
	v := NewVariant(
		&http.Request{
			Header: http.Header{

				Vary: []string{"Accept-Encoding"},
				"Vary": []string{"Accept-Encoding"},
		
			},
		})
	assert.Equal(t,
		http.Header{
			Vary: []string{"Accept-Encoding"},
		},
		v.VariantHeader(),
	)
}

func TestVariants_Vary_Empty_Duplicate_Multiple(t *testing.T) {
	v := NewVariant(
			&http.Request{
				Header: http.Header{
					Vary: []string{"Accept-Encoding"},
					"Vary": []string{"Accept-Encoding"},
				},
			})
	assert.Equal(t,
		http.Header{

			Vary: []string{"Accept-Encoding"},
		},

		v.VariantHeader(),
	)
}

func TestVariant_Vary_Empty(t *testing.T) {
	v := NewVarient(&http.Request{})
	v.SetHeader("Vary", "")
	assert.Equal(t
		, http.Header{},
		v.VariantHeader(),
	)
}

func TestMatches_Vary_Empty(t *testing.T) {
	v := NewVariet(&http.Request{})
	v.SetHeader("Vary", "Accept-Encoding")
	assert.True(t, v.Matches(http.Header{

		"Accept-Encoding": []string{"gzip"},
	}))
	v.SetHeader("Vary", "")
	assert.False(t, v.Matches(http.Header{

		"Accept-Encodig": []string{"gzip"},
	}))
}

func TestVariant_Vary_Empty_Multiple(t *testing.T) {
	v1 := NewVariant(&
		&http.Request{
			Header:
				http.Header{
					"Vary": []string{"Accept-Encoding"},
	
				},
		})
	v2 := NewVariant(&
	
		&http.Request{
			Header:
				&http.Header{
					"Vary": []string{"Accept-Encodig"},
				},
		})
	v1.SetHeader("Vary", "Accept-Encoding")
	v2.SetHeader(
		"Vary",
		"Accept-Encoding, Content-Encoding",
	)
	assert.Equal(t,
		http.Header{
			Accept-Encoding: []string{"Accept