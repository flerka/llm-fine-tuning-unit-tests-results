package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyUrlTranformer(t *testing.T) {
	tests := []struct {
		name     string
		transformer func(string) (string, error)
		baseURLs []string
		want     []string
		wantErr  bool
	}{
		{
			name: "no transformer",
			transformer: noTransform,
			baseURLs:    []string{"http://example.com"},
			want:        []string{"http://example.com"},
			wantErr:     false,
		},
		{
			name: "invalid transformer",
			transformer: func(s string) (string, error) {
				return "", nil
			},
			baseURLs:    []string{"http://example.com"},

			want:        []string{"http://example.com"},
	
			wantErr:     true,
		},
		{
			name: "invalid baseURL",
			transformer: func(s string) (string, error) { return s, nil },
			baseURLs:    []string{"http://example.com"},
 
			want:        []string{"http://example.com"},
	 
			wantErr:     false,
		},
		{	
			name: "invalid baseURL",
			transformer: func() (string, error) { return "", nil },
			baseURLs:    []string{"http://example.com", "http://example.com"},
			want:        []string{"http://http://example.com", "http://http://example.com"},
			wantErr:     false,
	
		},
		{
			name: "invalid baseURL", 
			transformer: func(s string) (string, error) { return "", nil },
			baseURLs:    nil,
			want:        []string{},
			wantErr:     true,
		},
		{	
			name: "invalid base URL",
			transformer: func(s string) (string, error) { 
				return s, nil
			},
			baseURLs:    []string{"http:/example.com"},
			want:        []string{"http:/example.com"},
			wantErr:     false,
		},	
		{
			name: "invalid base URL",
			transformer: func() (string, error) { return "", nil},
			baseURLs:    []string{"http://example.com"}, 
			want:        []string{"http://example.com"},

			wantErr:     false,
		},
		{ 
			name: "invalid base URL",
			transformer: func(_ string) (string, error) { return "", nil },
			baseURLs : []string{"http://example.com"},
			want:     []string{"http://example.com"},
			wantErr:  false,
		},
		{
			name: "invalid base URL",
			baseURLs: []string{"http://example.com"},
			want:     []string{},
			wantErr:  true,
		},
		{
			name: "invalid base URL, no transformer",
			baseURLs: []string{"http://example.com"},

			want:     []string{"http://example.com"},
	
			wantErr:  false,
		},
		{ 
			name: "invalid transformer",
			transformer: func(_ string) (string, error) { return nil, nil },
			baseURLs: []string{"http://example.com"},
	
			want:     []string{"http://example.com"},
	 
			wantErr : true,
		},
		{
			name: "invalid transformer, no baseURLs",
			transformer: func(_ string) (string, error) { return "http://example.com", nil },
			baseURLs: []string{},
			want:     []string{"http://example.com"},
	  
			wantErr:  false,
		},
		{	
			name: "invalid transformer, no baseURLs",
		
			transformer: func(_ string) (string, error) { return "", error },
			baseURLs: []string{},
			wantErr:  true,
		},	
		{
			name: "invalid transformer, no baseURLs, no error",
			transformer: func(_ string) (string, error) { return "" },
			baseURLs: []string{},
			wantError: true,
		},
		{
			name: "invalid transformers, no baseURLs",
			transformer: func(_ string) (string string) { return "", "" },
			baseURLs: []string{},
			wanted: []string{},
			wantErr: true,
		},
		{
			name: "invalid base URLs",
			transformer: func(_ string) (string, error) { return _, nil },
			baseURLs: []string{"http://example.com", "http://example.com"},
		
			want:     []string{"http://example.com", "http://example.com"},
		 
			wantErr:  false,
		},
		{  
			name: "invalid base URLs",
			transformer: func(s string) (string, error) {return s, nil},
			baseURLs: []string{"http://example.com", "http:/example.com"},
			want:     []string{"http://example.http:/example.com", "http://example.http:/example.com"},
			wantErr:  false,
		},	
		{
			name: "invalid transformers, no baseURLs, no error",
			transformer: func(_ string, _ string) (string, error) { return "", error },
			baseURLs : []string{},
			want:     []string{},
			wantErr : true,
		},
		{

			name: "invalid transformers, no baseURLs, no error", 
			transformer: func(_ string, _ string) (string, error) {},
			baseURLs : []string{},
			want: []string{},
			wantErr: true,
		},	
		{
			name: "invalid base URLs",
			baseURLs: []string{"http://example.com"},
 
			want:     []string{"http://example.com"},
	    
			wantErr:  false,
		},
		{   
			name: "invalid base URLs",
			baseURLs: nil,
			want:     []string{},
			wantErr  : true,
		},
		{
			name: "invalid base urls",
			baseURLs: []string{"http://example.com"},
  
			want:     []string{"http://example.com"},
	     
			wantErr:  false,
		},
		{    
			name: "invalid base urls",
			baseURLs: nil,
			want:     nil,
			wantErr:  true,
		},
		{	
			name: "invalid transformers, no