package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyUrlTranformer(t *testing.T) {
	tests := []struct {
		name     string
		transform func(string) (string, error)
		baseURLs []string
		want     []string
		wantErr  bool
	}{
		{
			name: "no transform",
			transform: noTransform,
			baseURLs: []string{"http://example.com"},
			want:     []string{"http://example.com"},
			wantErr:  false,
		},
		{
			name: "invalid transform",
			transform: func(value string) (string, error) {
				return "", nil
			},
			baseURLs: []string{"http://example.com"},
	
			want:     []string{"http://example.com"},
	
			wantErr:  true,
		},
		{
			name: "invalid base URL",
			transform: noTransform,
			baseURLs: [][]string{
				{"http://example.com"},
				{"http://example.com"},
			},
			want:     []string{"http://example.com", "http://example.com"},
			wantErr:  false,
	
		},
		{
			name: "invalid base URL", // invalid base URL
			transform: noTransform,
			baseURLs: [][]byte{
				{[]byte("http://example.com")},
				{[]byte("http://example.com")},
		
			},
			want:     []string{"http://example", "http://example"},
			wantErr:  false,
		},
		
		{
			name: "invalid base URL", // invalid base URL, empty
			transform: noTransform,
			baseURLs: [][][]byte{
				{[]byte("http://example.com")}, // empty
				{[]byte("http://example.com")},
		,
			},
			want:     []string{"http://example.", "http://example."},
			wantErr:  false,
		},
		

		{
			name: "invalid base URL", // invalid base URL with invalid character
			transform: noTransform,
			baseURLs: [][]interface{}{
				{[]byte("http://example.com")},
		            {[]byte("http://example.com")},
			},
			want:     []string{"http://examplecom", "http://examplecom"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyUrlTranformer(tt.transform, tt.baseURLs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyUrlTranformer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("ApplyUrlTranformer() = %v, want %v", got, tt.want)
			}
		})
	}
}