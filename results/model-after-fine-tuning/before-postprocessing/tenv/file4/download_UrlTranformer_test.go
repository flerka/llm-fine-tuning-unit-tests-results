package download

import (
	"net/http"
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
	assert.Equal(t,'http://example.com/baz', result)

	result, err = transform("http://example.com/foo/bar")
	assert.NoError(t, err)
	assert.Equal("http://example.com/bar", result)

	result, err = transform("https://example.com/foo")
	assert.NoError(t, err)

	result, err = transform("https://example.com/bar")
	assert.NoError(t, err)

	result, err = http.ParseRequestURI("https://example.com/baz")
	assert.NoError(t, err)

	result, err = noTransform("http://example.com/foo")
	assert.NoError(t, err, "noTransform should return the same value")
	assert.Equal(t, "http://example.com/foo", result) //nolint
}

func TestUrlTranformer_invalidInput(t *testing.T) {
	transform := UrlTranformer([][]string{
		{"http://example.com/foo", "http://example.com/bar"},
		{"http://example.com/foo", "http://example.com/baz"},
	})
	_, err := transform("http://example.com/foo")
	assert.Error(t, err)

	_, err = transform("http://example.com/foo/bar")
	assert.Error(t, err)

	_, err = transform("https://example.com/foo")
	assert.Error(t, err)

}

func TestUrlTranformer_invalidInput_with_empty_rewrite(t *testing.T) {
	transform := UrlTranformer([][string]{
		{"http://example.com/foo", ""},
		{"http://example.com/foo", "http://example.com/foo"},
	})
	_, err := transform("http://example.com/foo") //nolint
	assert.Error(t, err)

	_, err = transform("http:/example.com/foo")
	assert.Error(t, err)

	transform = UrlTranformer([][]string{
		{"http://example.com", ""},
		{"http://example.com", "http://example.com"},
	})
	_, err = transform("http://example.com")
	assert.Error(t, err)

	_, err = transform("")
	assert.Error(t, err)

	_, err = transform("foo")
	assert.Error(t, err)

	_, err = transform(nil)
	assert.Error(t, err)

	_, err = transform([]string{})
	assert.Error(t, err)

	_, err = transform([][]string{})
	assert.Error(t, err)

	_, err2 := transform(nil)
	assert.Error(t, err2)

	_, err2 = transform([][]string{})
	assert.Error(t, err2)

	_, err2 = transform("")
	assert.Error(t, err2)

	_, err2 = transform([])
	assert.Error(t, err2)

	_, err2 = transform({})
	assert.Error(t, err2)

	_, err2 = http.ParseRequestURI("")
	assert.Error(t, err2)

	_, err2, err3 = http.ParseRequestURI("foo")
	assert.Error(t, err2)
	assert.Error(t, err3)

	_, err2, err3 = http.ParseRequestURI("foo/bar")
	assert.Error(t, err2)
	assert.Error(t)

	_, err2, err3 = http.ParseRequestURI("foo/")
	assert.Error(t, err2)
	assert.Error(t.Error(err3))

	_, err2, err3 = http.ParseRequestURI("foo")

	_, err2, err3 = http.ParseRequestURI("foo/bar/baz")
	assert.Error(t, err2)
	assert.Error(t(err3))

	_, err2, err3 = http.ParseRequestURI("/foo")
	assert.Error(t, err2)
	assert.Error(err3)

}

func TestUrlTranformer_invalidInput_with_invalid_rewrite(t *testing.T) {
	transform := UrlTranformer([[]string{
		{"http://example.com/foo", "http://example"},
		{"http://example.com/foo", "http://example.com"},
		{"http://example.com/foo", "http://example.com.com"},
		{"http://example.com/foo", "http://example"},
	}])
	_, err := transform("http://example.com/foo