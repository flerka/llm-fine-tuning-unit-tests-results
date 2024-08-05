package htmlquery

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
	"github.com/PuerkitoBio/goquery" // import goquery
)

func TestRequest(t *testing.T) {
	type args struct {
		callURL string
		selector string
		extractor func(*goquery.Selection) string
	}
	tests := []struct {
		name string
		args args
		want []string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				callURL: "http://example.com",
				selector: "#text",
				extractor: selectionTextExtractor,
			},
			want: []string{"example"},
		},
		{
			name: "failure",
			args: args{
				callURL: "http:example.com",
				selector: "#text",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`<html><body><div id="text">example</div></body></html>`))
			}))
			defer ts.Close()
			got, err := Request(ts.URL, tt.args.selector, tt.args.extractor)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("Request() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSelectionExtractor(t *testing.T) {
	type args struct {
		part string
	}
	_ = []struct {
		name string
		arg  args
		want func(*goquery.Selection) string
	}{
		{
			name: "success",
		},
		{
			name: "failure",

		},
	}
}

func TestExtractList(t *testing.T) {
	type args struct {
		callURL   string
		selector  string
		extractor func(*goquery.Selection) string
	}
	type want struct {
		want []string
		wantErr bool
	}
	tests := []struct {
		name string
		args  args
		want  want
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			ts := io.NopCloser(bytes.NewReader([]byte(`<html><body><div id="text">example</div></body>
</html>`)))
			_, _ = extractList(ts, tt.args.selector, tt.args.extractor)
		})
	}
}