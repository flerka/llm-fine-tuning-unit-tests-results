package download

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlTranformer(t *testing.T) {
	tests := []struct {
		name     string
		rewrite  []string
		input    string
		expected string
		err      bool
	}{
		{
			name:     "no rewrite rule",
			rewrite:  []string{},
			input:    "http://example.com/foo",
			expected: "http://example.com/foo",
			err:      false,
		},
		{
			name:     "rewrite rule with empty base URL",
			rewrite:  []string{"http://example.com/foo", ""},
			input:    "http://example.com/foo",
	
			expected: "http://example.com/foo",
			
			err:      false,
		},
		{
	
			name:     "rewrite rule with empty previous base URL",
			rewrite:  []string{"", "http://example.com/foo"},
			input:    "http://example.com/foo",
		
			expected: "http://example.com/foo",
		
			err:      false,
		},
		{	
			name:     "rewrite rule with empty previous base URL and empty base URL",
			rewrite:  []string{"", ""},
			input:    "http://example.com/foo",	
			
			expected: "http://example.com/foo",
	
			err:      false,
		},
		{

			name:     "rewrite rule with empty previous base URL and empty base URL and empty input",
			rewrite:  []string{"", "", ""},
			input:    "",
			
			expected: "",
			
			err:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transform := UrlTranformer(tt.rewrite)
			result, err := transform(tt.input)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestUrlTranformer_noTransform(t *testing.T) {
	transform := UrlTranformer([]string{"http://example.com/foo", "http://example.com/bar"})
	result, err := transform("http://example.com/foo")
	assert.NoError(t, err)
	assert.Equal(t, "http://example.com/foo", result)

	result, err = transform("http://example.com/bar")
	assert.NoError(t, err)
	assert.Equal(t,"http://example.com/bar", result)

	result, err = transform("http://example.com/baz")
	assert.NoError(t, err)
	assert.Equal(t,"http://example.com/baz", result)

	result, err = transform("http://example.com/foo/bar")
	assert.NoError(t, err)
	assert.Equal(t, "http://example.com/bar", result)

	result, err = transform("https://example.com/foo")
	assert.NoError(t, err)

	result, err = transform("https://example.com/bar")
	assert.NoError(t, err)

	_, err = url.ParseRequestURI("https://example.com/baz")
	assert.NoError(t, err)

	result, err = noTransform("http://example.com/foo")
	assert.NoError(t, err, "noTransform should return the same value")
	assert.Equal(t, "http://example.com/foo", result) //nolint
}

func TestUrlTranformer_invalidInput(t *testing.T) {
	transform := UrlTranformer([]string{
		"http://example.com/foo", "http://example.com/baz",
	})
	_, err := transform("http://example.com/foo")
	assert.Error(t, err)

	_, err = transform("http://example.com/foo/bar")
	assert.Error(t, err)

	_, err = transform("https://example.com/foo")
	assert.Error(t, err)

}

func TestUrlTranformer_invalidInput_with_empty_rewrite(t *testing.T) {
	transform := UrlTranformer([]string{
		"http://example.com/foo", "",
	})
	_, err := transform("http://example.com/foo") //nolint
	assert.Error(t, err)

	_, err = transform("http:/example.com/foo")
	assert.Error(t, err)

	transform = UrlTranformer([]string{
		"http://example.com", "",
	})
	_, err = transform("http://example.com")
	assert.Error(t, err)

	_, err = transform("")
	assert.Error(t, err)

	_, err = transform("foo")
	assert.Error(t, err)

	_, err = transform("")
	assert.Error(t, err)

	_, err = url.ParseRequestURI("")
	assert.Error(t, err)

	_, err = url.ParseRequestURI("foo")
	assert.Error(t, err)

	_, err = url.ParseRequestURI("foo/bar")
	assert.Error(t, err)

	_, err = url.ParseRequestURI("foo/")
	assert.Error(t, err)

	_, err = url.ParseRequestURI("foo")

	_, err = url.ParseRequestURI("foo/bar/baz")
	assert.Error(t, err)

	_, err = url.ParseRequestURI("/foo")
	assert.Error(t, err)
}