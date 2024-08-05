package github

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetDownloadURL_Success(t *testing.T) {
	// Setup mock server
	mockResponse := `{
		"assets_url": "http://example.com/assets"
	}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Mock API response for assets
	mockAssetsResponse := `[
		{
			"name": "test-asset",
			"browser_download_url": "http://example.com/download/test-asset"
		}
	]`
	assetsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockAssetsResponse))
	}))
	defer assetsServer.Close()

	// Replace URL join path function to return mock server URL
	urlJoinPath = func(base string, elem ...string) (string, error) {
		return mockServer.URL, nil
	}
	defer func() { urlJoinPath = url.JoinPath }()

	// Test case
	tag := "v1.0.0"
	searchedAssetNames := []string{"test-asset"}
	githubReleaseURL := "http://example.com"
	githubToken := "dummy_token"

	assetURLs, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, func(msg string) {})

	assert.NoError(t, err)
	assert.Equal(t, []string{"http://example.com/download/test-asset"}, assetURLs)
}

func TestAssetDownloadURL_InvalidURL(t *testing.T) {
	// Replace URL join path function to return an error
	urlJoinPath = func(base string, elem ...string) (string, error) {
		return "", errors.New("invalid URL")
	}
	defer func() { urlJoinPath = url.JoinPath }()

	// Test case
	tag := "v1.0.0"
	searchedAssetNames := []string{"test-asset"}
	githubReleaseURL := "http://example.com"
	githubToken := "dummy_token"

	assetURLs, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, func(msg string) {})

	assert.Error(t, err)
	assert.Nil(t, assetURLs)
}

func TestAssetDownloadURL_NoAssetsURL(t *testing.T) {
	// Setup mock server
	mockResponse := `{}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Replace URL join path function to return mock server URL
	urlJoinPath = func(base string, elem ...string) (string, error) {
		return mockServer.URL, nil
	}
	defer func() { urlJoinPath = url.JoinPath }()

	// Test case
	tag := "v1.0.0"
	searchedAssetNames := []string{"test-asset"}
	githubReleaseURL := "http://example.com"
	githubToken := "dummy_token"

	assetURLs, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, func(msg string) {})

	assert.Error(t, err)
	assert.Nil(t, assetURLs)
}

func TestAssetDownloadURL_NoAssetsFound(t *testing.T) {
	// Setup mock server
	mockResponse := `{
		"assets_url": "http://example.com/assets"
	}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Mock API response for assets
	mockAssetsResponse := `[]`
	assetsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockAssetsResponse))
	}))
	defer assetsServer.Close()

	// Replace URL join path function to return mock server URL
	urlJoinPath = func(base string, elem ...string) (string, error) {
		return mockServer.URL, nil
	}
	defer func() { urlJoinPath = url.JoinPath }()

	// Test case
	tag := "v1.0.0"
	searchedAssetNames := []string{"test-asset"}
	githubReleaseURL := "http://example.com"
	githubToken := "dummy_token"

	assetURLs, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, func(msg string) {})

	assert.Error(t, err)
	assert.Nil(t, assetURLs)
}

func TestAssetDownloadURL_APIError(t *testing.T) {
	// Setup mock server to return error
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	// Replace URL join path function to return mock server URL
	urlJoinPath = func(base string, elem ...string) (string, error) {
		return mockServer.URL, nil
	}
	defer func() { urlJoinPath = url.JoinPath }()

	// Test case
	tag := "v1.0.0"
	searchedAssetNames := []string{"test-asset"}
	githubReleaseURL := "http://example.com"
	githubToken := "dummy_token"

	assetURLs, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, func(msg string) {})

	assert.Error(t, err)
	assert.Nil(t, assetURLs)
}