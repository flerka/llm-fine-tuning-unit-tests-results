package download_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyUrlTranformer_Success(t *testing.T) {
	transformer := func(url string) (string, error) {
		return "transformed_" + url, nil
	}
	urls := []string{"http://example.com", "http://example.org"}

	transformedURLs, err := ApplyUrlTranformer(transformer, urls...)

	assert.NoError(t, err)
	assert.Equal(t, []string{"transformed_http://example.com", "transformed_http://example.org"}, transformedURLs)
}

func TestApplyUrlTranformer_TransformerError(t *testing.T) {
	transformer := func(url string) (string, error) {
		return "", errors.New("transform error")
	}
	urls := []string{"http://example.com"}

	transformedURLs, err := ApplyUrlTranformer(transformer, urls...)

	assert.Error(t, err)
	assert.Nil(t, transformedURLs)
}

func TestApplyUrlTranformer_EmptyInput(t *testing.T) {
	transformer := func(url string) (string, error) {
		return "transformed_" + url, nil
	}

	transformedURLs, err := ApplyUrlTranformer(transformer)

	assert.NoError(t, err)
	assert.Empty(t, transformedURLs)
}

func TestBytes_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))
	defer server.Close()

	display := func(msg string) {}
	data, err := Bytes(server.URL, display)

	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), data)
}

func TestBytes_HttpError(t *testing.T) {
	display := func(msg string) {}
	data, err := Bytes("http://invalid.url", display)

	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestBytes_InvalidUrl(t *testing.T) {
	display := func(msg string) {}
	data, err := Bytes(":", display)

	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestUrlTranformer_Success(t *testing.T) {
	rule := []string{"http://example.com", "http://newexample.com"}
	transformer := UrlTranformer(rule)

	transformedURL, err := transformer("http://example.com/path")

	assert.NoError(t, err)
	assert.Equal(t, "http://newexample.com/path", transformedURL)
}

func TestUrlTranformer_NoTransformNeeded(t *testing.T) {
	rule := []string{"http://example.com", "http://newexample.com"}
	transformer := UrlTranformer(rule)

	transformedURL, err := transformer("http://other.com/path")

	assert.NoError(t, err)
	assert.Equal(t, "http://other.com/path", transformedURL)
}

func TestUrlTranformer_InvalidRules(t *testing.T) {
	rule := []string{}
	transformer := UrlTranformer(rule)

	transformedURL, err := transformer("http://example.com/path")

	assert.NoError(t, err)
	assert.Equal(t, "http://example.com/path", transformedURL)
}

func TestUrlTranformer_EmptyPreviousBaseUrl(t *testing.T) {
	rule := []string{"", "http://newexample.com"}
	transformer := UrlTranformer(rule)

	transformedURL, err := transformer("http://example.com/path")

	assert.NoError(t, err)
	assert.Equal(t, "http://example.com/path", transformedURL)
}

func TestUrlTranformer_EmptyBaseUrl(t *testing.T) {
	rule := []string{"http://example.com", ""}
	transformer := UrlTranformer(rule)

	transformedURL, err := transformer("http://example.com/path")

	assert.NoError(t, err)
	assert.Equal(t, "http://example.com/path", transformedURL)
}