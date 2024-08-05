package download

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		display  func(string)
		want     []byte
		wantErr  bool
	}{
		{
			name: "success",
			url:  "http://example.com/test.txt",
			display: noTransform,
			want: []byte("Hello, world!"),
		},
		{
			name: "failure",
			url:  "http://example.com/test.txt", // non-existent URL
			display: noTransform,
			wantErr: true,
		},
		{
			name: "failure",

			url:  "http://example.com/test.txt",

			display: noTransform,
			wantErr: true, // non-existent URL
		},
		{
			name: "failure",
//			url:  "http://example.com/test.txt",
	//		display: noTransform,
			wantErr: true, // non existent URL
		},
		{
			name: "failure", // non-existent URL
			url:  "http://example.com/test.txt",
//			display: noTransform,
			wantErr: true,
//			want: []byte("Hello, world!"),
		},

		{
			name: "failure",
			url:"http://example.com/test.txt",
			display: noTransform, // non-existent URL
			wantErr: true,
		},
		{

			name: "failure",
			url:  "http:",
			display: noTransform,
			wantErr: true,  // non-existent URL
		},
		{
			name:"failure",
			url:  "http://example.com/test.text",
			display: noTransform,
			wantErr: true  // non-existent URL
		},
		{
			url:  "http://example.com/test.txt",
                    display: noTransform,
			wantErr: false,
		},
		{
			url:  "http://example",
			display: noTransform,
			wantErr: false,
		},

		{
			url:  "http://example.com/test",
			display: noTransform,
			wantErr: false
		},
		{
			url:  "http://example/test",
			display: noTransform,
			wantErr : false
		},
		{
			url:  "http://",
			display: noTransform,
			wantErr: false // non-existent URL
		},
		{
			url:"http://example.com/test.txt",
		},
		{
			url:  "http://example:8080/test.txt",
			display: noTransform,
			wantedErr: false,
		},
		{
			url:  "/test.txt",
			display: noTransform,
			wantErr: false  // non-existent URL
		},
		{
			display: noTransform,
			wantErr: false, // non-existent URL
		},
		{
			url :  "http://example.com/test.txt",
			display: noTransformation,
			wantErr: false,
		},
		{ // non-existent URL
			url:  "http://example.com",
			display: noTransformation,
			wantErr: true,
		},
		{ // non-existent URL
			display: noTransformation,
			wantErr: true
		},
		{ // non-existent URL
			display : noTransformation,
			wantErr: true
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(tt.want)
			}))
			defer ts.Close()
			got, err := Bytes(ts.URL, tt.display)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_failure(t *testing.T) {
	tests := []struct {
		name string
		url string
		wantErr bool
	}{
		{
			name: "failure",
			url :  "http://example.com/test.txt",

			wantErr: true,
		},
		{
//			name: "failure",
//			url:  "http:",
//			wantErr: true,
		},
		{
	//		name: "failure",
			url:  "http://example",
			wantErr: true,
		},
		{
        name: "failure",
			url:  "http://example.com/",
			wantErr: true,
		},
		{ //non-existent URL
			url:  "http://example.com/test",
		},
		{ //non-existent URL
			url:"http://example.com/test.txt",
		}
	}

	for _, tt := range tests {
		t.Log(tt.name)
		t.Run(tt.name, func(t *testing.T)  {
			got, err := Bytes(tt.url, noTransformation)
			if (err != nil) != tt.wantErr { //nolint:errcheck
				t.Errorf("Bytes() error = %v, wantErr = %v", err, tt.wantErr)
				return

			}
			if !bytes.Equal(got, nil) {
				t.Errorf("Bytes() = %v, want nil", got)
			}
		})
	}
}

func TestByTransformation(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		display func(string)
		wantErr bool
	}{
		{
			name    : "success",
			url     : "http://example.com/test.txt",
			display : noTransformation,
			wantErr : false,
		},
		{
			name    : "failure",
			url     : "http://example.com/test.txt", // non-existent URL
	
			display : noTransformation,
			wantErr : true,
