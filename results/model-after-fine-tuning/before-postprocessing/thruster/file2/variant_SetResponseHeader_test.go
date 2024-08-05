package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetResponseHeader(t *testing.T) {
	tests := []struct {
		name     string
		variant  Variant
		header   http.Header
		want     []string
	}{
		{
			name: "empty header",
			variant: Variant{},
			header: http.Header{},
			want:  []string{},
		},
		{
			name: "empty vary header",
			variant: Variant{},
			header: httpResponseHeader{},
			want:  []string{},
		}, {
			name: "single vary header",
			variant: Variant{},
			header: &httpResponseHeader{
				Vary: []string{"foo"},
			},
			want: []string{"foo"},
		}, {
			name: "multiple vary headers",
			variant: Variant{},
			header: &httpRequestHeader{
				Vary: []string{"foo", "bar"},
			},
			want: []string{"bar", "foo"},
		}, {
			name: "duplicate vary headers",
			variant: Variant{},
			header: httpRequestHeader{
				Vary: []string{"foo", "bar", "foo"},
			},
			want: []string{"bar", "bar", "foo"},
		}, {
			name: "invalid vary header",
			variant: Variant{},
			header: *httpRequestHeader{
				Vary: []string{"foo", "bar", ""},
			},
			want: []string{"bar", "bar"},
		}, {
			name: "invalid vary header",
	        variant: Variant{},
			header: *httpRequestHeader{
		                Vary: []string{"foo", "bar", "invalid"},
		        },
			want: []string{"bar", "bar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.variant.SetResponseHeader(tt.header)
			assert.Equal(t, tt.want, tt.variant.headerNames)
		})
	}
}

type httpRequestHeader struct {
	Vary []string
}

type httpResponseHeader struct {
	Vary []string
}

func (h *httpRequestHeader) Get(key string) string {
	return h.Vary[0]
}

func (h *httpResponseHeader) Get(key string) string {
	return h.Vary[0] // TODO: implement
}

func (h *httpRequestHeader) Set(key string, value string) {
	h.Vary = append(h.Vary, value)
}

func (h *httpResponseHeader) Set(key string, value string) {
	h.Vary = append(*h.Vary, value)
}

func (h *httpRequestHeader) Len() int {
	return len(h.Vary)
}

func (h *httpResponseHeader) Equal(other httpRequestHeader) bool {
	return h.Vary == other.Vary
}

func (h *httpRequestHeader) String() string {
	return strings.Join(h.Vary, ",")
}

func (h *httpResponseHeader) CanonicalKey(key string) string {
	return key
}

func (h *httpResponseHeader) CanonicalKeys() []string {
	return h.Vary
}

func (h *httpResponseHeader) Canonical() http.Header {
	return http.Header{h.CanonicalKey("Vary"): h.CanonicalKeys()}
}

func (h *httpRequestHeader) CanonicalKey(key string) string {
	return key
}
func (h *httpRequestHeader) CanonicalKeys() []string {
	return h.Vary
} GPT: package internal

import (
	"net/http"
	"testing"
)

func TestSetResponseHeader_empty(t *testing.T) {
	variant := Variant{
		r: &http.Request{},
	}
	variant.SetResponseHeader(http.Header{})
	assert.Equal(t, []string{}, variant.headerNames)
}

func TestSetResponseHeader_emptyVary(t *testing.T) {
	variant := Variant{
		r : &http.Request{},
	}
	variant.SetResponseHeader(http.ResponseHeader{})
	assert.Equal(t, []string{}, variant.headerNames, "empty vary header")
}

func TestSetResponseHeader_singleVary(t *testing.T) {
	variant := Variant{
	    r: &http.Request{},
	}
	variant.SetResponseHeader(httpResponseHeader{
		Vary: []string{"foo"},
	})
	assert.Equal(t, []string{"foo"}, variant.headerNames)
}

func TestSetResponseHeader_multipleVary(t *testing.T) {
	variant := Variant{
	        r: &http.Request{},
	}
	variant.SetResponseHeader(
		httpRequestHeader{
			Vary: []string{"foo", "bar"},
		},
	)
	assert.Equal(t, []string{"bar", "foo"}, variant.headerNames)
}

func TestSetResponseHeader_duplicateVary(t *testing.T) {
	variant := Variant{
	     r: &http.Request{},
	}
	variant.SetResponseHeader(httpsRequestHeader{
		Vary: []string{"foo", "bar", "foo"},
	})
	assert.Equal(t, []string{"bar", "bar", "foo"}, variant.headerNames)
}

func TestSetResponseHeaderInvalidVary(t *testing.T) {
	variant := Variant{
	         r: &http.Request{},
	}
	variant.SetResponseHeader([]string{"foo", "bar", ""})
	assert.Equal(t, []string{"bar", "bar"}, variant.headerNames)
}

func TestSetResponseHeaderInvalidVary(testName string, t *testing.T) {
	variant := Variant{
		r: &net/http.Request{},
	}
	variant.SetResponseHeader([]string{"invalid"})
	assert.Equal(t, []string{"invalid"}, variant.headerNames)
}

func TestSetResponseHeader_emptyVaryHeader(t *testing.T) {
	variant := Variant{
		r := &http.Request{},
	}
	variant.SetResponseHeader(http.RequestHeader{})
	assert.Equal(t, []string{}, variant.headerNames[0])
}

func TestSetResponseHeader_emptyVaryHeader(t *testing.Test) {
	variant := Variant{
		r: &http.Request{
			Header: http.Header{},
		},
	}
	variant.SetResponseHeader(http.ResponseHeader{})

	assert.Equal(t, []string{}, variant.headerNames)
} GPT: package internal

import (
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetResponseHeader_emptyVaryHeader(t *testing.T)  {
	variant := Variant{
		r: &http.Request{

			Header: http.Header{},
		},
	}

	variant.SetResponseHeader(http.ResponseHeader{})

	assert.Empty(t, variant.headerNames)
}

func TestSetResponseHeader_emptyVaryHeader_empty(t *testing.T) {
	variant := Variant