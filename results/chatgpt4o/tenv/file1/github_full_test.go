package github

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock display function
func mockDisplay(msg string) {
	// Do nothing
}

// Mock versionfinder
var mockVersionfinder = versionfinder.Find

// Mock API responses
func mockAPIGetRequest(url string, authorizationHeader string) (any, error) {
	if url == "https://api.github.com/repos/owner/repo/tags/v1.0.0" {
		return map[string]any{
			"assets_url": "https://api.github.com/repos/owner/repo/releases/1/assets",
		}, nil
	} else if url == "https://api.github.com/repos/owner/repo/releases/1/assets?page=1" {
		return []any{
			map[string]any{
				"name":                 "asset1.zip",
				"browser_download_url": "https://github.com/owner/repo/releases/download/v1.0.0/asset1.zip",
			},
		}, nil
	} else if url == "https://api.github.com/repos/owner/repo/releases?page=1" {
		return []any{
			map[string]any{
				"tag_name": "v1.0.0",
			},
		}, nil
	}
	return nil, errors.New("not found")
}

func TestAssetDownloadURL_Success(t *testing.T) {
	originalApiGetRequest := apiGetRequest
	defer func() { apiGetRequest = originalApiGetRequest }()
	apiGetRequest = mockAPIGetRequest

	tag := "v1.0.0"
	searchedAssetNames := []string{"asset1.zip"}
	githubReleaseURL := "https://api.github.com/repos/owner/repo"
	githubToken := "mocktoken"

	urls, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, mockDisplay)
	assert.NoError(t, err)
	assert.Equal(t, []string{"https://github.com/owner/repo/releases/download/v1.0.0/asset1.zip"}, urls)
}

func TestAssetDownloadURL_AssetNotFound(t *testing.T) {
	originalApiGetRequest := apiGetRequest
	defer func() { apiGetRequest = originalApiGetRequest }()
	apiGetRequest = mockAPIGetRequest

	tag := "v1.0.0"
	searchedAssetNames := []string{"nonexistent.zip"}
	githubReleaseURL := "https://api.github.com/repos/owner/repo"
	githubToken := "mocktoken"

	urls, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, mockDisplay)
	assert.Error(t, err)
	assert.Nil(t, urls)
}

func TestAssetDownloadURL_InvalidURL(t *testing.T) {
	originalApiGetRequest := apiGetRequest
	defer func() { apiGetRequest = originalApiGetRequest }()
	apiGetRequest = mockAPIGetRequest

	tag := "v1.0.0"
	searchedAssetNames := []string{"asset1.zip"}
	githubReleaseURL := ":/invalid_url"
	githubToken := "mocktoken"

	urls, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, mockDisplay)
	assert.Error(t, err)
	assert.Nil(t, urls)
}

func TestListReleases_Success(t *testing.T) {
	originalApiGetRequest := apiGetRequest
	defer func() { apiGetRequest = originalApiGetRequest }()
	apiGetRequest = mockAPIGetRequest

	githubReleaseURL := "https://api.github.com/repos/owner/repo"
	githubToken := "mocktoken"

	releases, err := ListReleases(githubReleaseURL, githubToken)
	assert.NoError(t, err)
	assert.Equal(t, []string{"v1.0.0"}, releases)
}

func TestListReleases_NoReleases(t *testing.T) {
	originalApiGetRequest := apiGetRequest
	defer func() { apiGetRequest = originalApiGetRequest }()
	apiGetRequest = func(url string, authorizationHeader string) (any, error) {
		return []any{}, nil
	}

	githubReleaseURL := "https://api.github.com/repos/owner/repo"
	githubToken := "mocktoken"

	releases, err := ListReleases(githubReleaseURL, githubToken)
	assert.NoError(t, err)
	assert.Empty(t, releases)
}

func TestListReleases_ApiError(t *testing.T) {
	originalApiGetRequest := apiGetRequest
	defer func() { apiGetRequest = originalApiGetRequest }()
	apiGetRequest = func(url string, authorizationHeader string) (any, error) {
		return nil, errors.New("api error")
	}

	githubReleaseURL := "https://api.github.com/repos/owner/repo"
	githubToken := "mocktoken"

	releases, err := ListReleases(githubReleaseURL, githubToken)
	assert.Error(t, err)
	assert.Nil(t, releases)
}