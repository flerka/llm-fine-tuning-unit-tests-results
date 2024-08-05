package internal

import (
	"net/http"
	"testing"
	"strings"
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
			header:  http.Header{},
			want:  []string{},
		}, {
			name: "single vary header",
			variant: Variant{},
			header:  http.Header{
				"Vary": []string{"foo"},
			},
			want: []string{"foo"},
		}, {
			name: "multiple vary headers",
			variant: Variant{},
			header:  http.Header{
				"Vary": []string{"foo", "bar"},
			},
			want: []string{"bar", "foo"},
		}, {
			name: "duplicate vary headers",
			variant: Variant{},
			header:  http.Header{
				"Vary": []string{"foo", "bar", "foo"},
			},
			want: []string{"bar", "bar", "foo"},
		}, {
			name: "invalid vary header",
			variant: Variant{},
			header:  http.Header{
				"Vary": []string{"foo", "bar", ""},
			},
			want: []string{"bar", "bar"},
		}, {
			name: "invalid vary header",
	        variant: Variant{},
			header:  http.Header{
		                "Vary": []string{"foo", "bar", "invalid"},
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

func (h *httpRequestHeader) Len() int {
	return len(h.Vary)
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
}