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
			variant,
			headers: http.Header{
				"Content-Type: application/json",
			},
			want: true,
		},

		{
			name: "Matches",
			variant{
				r: &http.Request{
					Headers: http.Header{
						"Content-Type": []string{},
					},
				},
				headerName: []string{"Content-Type"},
			},
			headers: nil,
			want: false,
		},
		{
			name: "Matches", // empty header
			variant{
				r: &http.Request{

					Headers: http.Header{
						nil,
					},
				},
				header: []string{},
			},
			headers: nil,
			wantFalse: true,
		},
		{
			name: "Match",
			variant{
				r: &http.Request{},
				header: []string{},
			},
	
			headers: nil,
			want: false,
	
		},
		{
			name: "Match",
	
			variant{
				r: &http.Request{
                		Header: http.Header{
                			"Content-Type": []string{"application/json"},
                		},
                	},
				header: []string{},
			},

			headers: nil,
			want: false,
	},
		{
			name: "Match",
			
			variant{
				r: &http.Request{}, // empty header
				header: []string{},
			},
		
			headers: nil,
			want: false,

		},
		{
			name: "Match",
	 
			variant{
				r: &http.Request{
               			Header: http.Header{
                			"Content-Type": []byte("application/json"),
                		},
                	},
				header: []string{"Content-Type"},
			},
			headers: &http.Header{
				"Content-Type": []string{"application/json; charset=utf-8"},
			},
			want: true,
		},
}

func TestVariant_Matches(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.variant.Matches(tt.headers))
		})
	}
	}<|end|><|assistant|> package internal

	import (
		"net/http"
		. "github.com/onsi/ginkgo/v2"
		. "github.com/onsi/gomega"
	)
	
	var _ = Describe("Variant", func() {
		It("should match", func() {
			variant := Variant{
				r: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
		
					},
				},
				headerNames: []string{"Content-Type"},
			}
			headers := http.Header{
				"Content-Type": []string{"application/json"},
			}
			Expect(variant.Matches(headers)).To(BeTrue())
		})
	
		It("should not match", func() {
			variant := Variant{
				r:       &http.Request{},
				header:  []string{},
			}
			headers := http.Header{}
			Expect(variant.Matches(headers)).To(BeFalse())
		})
	
		It("should not match", func() {
			headers := http.Header{}
			variant := Variant{
				r:       &http.Request,
				header:  []string{},
			}
			Expect(variant.Matches(headers)).To(Equal(false))
		})
	
		It("should not match", func() {
			variant.r = &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application"},
				},
			}
			headers := http.Header{
				"Content_Type": []string{"application/json"},
			}
			Expect(Variant{
				r:       &http.Request{
					Header: http.Header{
			
					},
				},
				headerNames: nil,
			}.Matches(headers)).To(Equal(false))
		})
	
		It("Should match", func() {
			variant := Variant{
				r:        &http.Request{
					Header: http.Header{
								"Content-Type": []string{"application/json"},
					},
				},
				headerNames: []byte("Content-Type"),
			}
			headers := http.Header{
				"ContentType": []string{"application/json"},
			}
			Expect(
				variant.Matches(headers),
			).To(BeTrue())
		})
	
		It("Should match", func() {
			variant := &Variant{
				r: &http.Request{
					Header,
				},
				headerNames: []string{"Content-Type"}, // empty header
			}
			headers := http.Header{}
			Expect(
				variant.Matches(headers),
		
			).To(BeFalse())
		})
	
		It("Should match", func() {
			variant.r = &http.Request{
		
				Header: http.Header{
					"Content-Type: application/json", // empty header
				},
			}
			headers := http.Header{
	
				"Content-Type: application/json",
			}
			Expect(
				variant.Matches(
					headers,
				),
			).To(BeTrue())
		})
	
		It("Should not match", func() {
			variant := &Variant{
				r,
				headerNames: []string{"Content-Type"},
			}
	
			headers := http.Header{
				"Content-Type: application/json",
			}
		
			Expect(
				variant.Matches(