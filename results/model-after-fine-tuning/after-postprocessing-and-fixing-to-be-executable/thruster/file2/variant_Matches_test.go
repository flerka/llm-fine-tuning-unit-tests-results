package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariant_Matches(t *testing.T) {
	tests := []struct {
		name     string
		variant  Variant
		headers  http.Header
		want     bool
	}{
		{
			name: "Matches",
			variant: Variant{
				r: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
				},
				headerNames: []string{"Content-Type"},
			},
			headers: http.Header{
				"Content-Type": []string{"application/json"},
		
			},
			want: true,
		},
		{
			name: "Matches",
			headers: http.Header{
				"Content-Type": []string{"application/json"},
			},
			want: true,
		},

		{
			name: "Matches",
			variant: Variant{
				r: &http.Request{
					Header: http.Header{
						"Content-Type": []string{},
					},
				},
			},
			headers: nil,
			want: false,
		},
		{
			name: "Matches", // empty header
			variant: Variant{
				r: &http.Request{
					Header: http.Header{},
				},
			},
			headers: nil,
			want: false,
		},
		{
			name: "Match",
			variant: Variant{
				r: &http.Request{},
			},
			headers: nil,
			want: false,
		},
		{
			name: "Match",
	
			variant: Variant{
				r: &http.Request{
                		Header: http.Header{
                			"Content-Type": []string{"application/json"},
                		},
                	},
			},

			headers: nil,
			want: false,
	},
		{
			name: "Match",
			
			variant: Variant{
				r: &http.Request{}, // empty header
			},
		
			headers: nil,
			want: false,

		},
		{
			name: "Match",
	 
			variant: Variant{
				r: &http.Request{
               			Header: http.Header{
                			"Content-Type": []string{"application/json"},
                		},
                	},
					headerNames: []string{"Content-Type"},
			},
			headers: http.Header{
				"Content-Type": []string{"application/json; charset=utf-8"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.variant.Matches(tt.headers))
		})
	}
}