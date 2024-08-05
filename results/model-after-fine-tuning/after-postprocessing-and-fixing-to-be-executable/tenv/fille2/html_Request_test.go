package htmlquery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery" // import goquery
)

func TestRequest(t *testing.T) {
	tests := []struct {
		name     string
		callURL  string
		selector string
		extractor func(*goquery.Selection) string
		want     []string
		wantErr  bool
	}{
		{
			name: "success",
			callURL: "http://example.com",
			selector: ".title",
			extractor: func(s *goquery.Selection) string {
				return s.Text()
			},
			want: []string{"Example Domain"},
		},
		{
			name: "failure",
			callURL: "http://example.com",
			wantErr: true,
		},
		{
			name: "empty",
			callURL: "http://example.com",
			extractor: func(s *goquery.Selection) string { return "" },
			want: []string{},
		},
		{
			name: "no selector",
			callURL: "http://example.com",
			selector: ".title",
			extractor: func(s *goquery.Selection) string { return s.Text() },
			want: []string{"Example Domain"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`<html><body><div class="title">Example Domain</div></body></html>`))
			}))
			defer ts.Close()
			got, err := Request(ts.URL, tt.selector, tt.extractor)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range tt.want {
			 if	got[i] != v {
				t.Errorf("Request() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}