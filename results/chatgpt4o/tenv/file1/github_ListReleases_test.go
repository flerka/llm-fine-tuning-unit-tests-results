package github

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock the apiGetRequest function
func mockAPIGetRequest(callURL string, authorizationHeader string) (any, error) {
	if callURL == "https://api.github.com/repos/user/repo/releases?page=1" {
		return []any{
			map[string]any{"tag_name": "v1.0.0"},
			map[string]any{"tag_name": "v1.1.0"},
		}, nil
	} else if callURL == "https://api.github.com/repos/user/repo/releases?page=2" {
		return []any{}, nil
	}
	return nil, errors.New("not found")
}

func TestListReleases_Success(t *testing.T) {
	apiGetRequest = mockAPIGetRequest // Override the apiGetRequest function

	releases, err := ListReleases("https://api.github.com/repos/user/repo/releases", "dummyToken")

	assert.NoError(t, err)
	assert.Equal(t, []string{"v1.0.0", "v1.1.0"}, releases)
}

func TestListReleases_InvalidURL(t *testing.T) {
	apiGetRequest = func(callURL string, authorizationHeader string) (any, error) {
		return nil, errors.New("invalid URL")
	}

	_, err := ListReleases(":", "dummyToken")
	assert.Error(t, err)
}

func TestListReleases_EmptyReleases(t *testing.T) {
	apiGetRequest = func(callURL string, authorizationHeader string) (any, error) {
		return []any{}, nil
	}

	releases, err := ListReleases("https://api.github.com/repos/user/repo/releases", "dummyToken")

	assert.NoError(t, err)
	assert.Empty(t, releases)
}

func TestListReleases_NoToken(t *testing.T) {
	apiGetRequest = mockAPIGetRequest

	releases, err := ListReleases("https://api.github.com/repos/user/repo/releases", "")

	assert.NoError(t, err)
	assert.Equal(t, []string{"v1.0.0", "v1.1.0"}, releases)
}

func TestListReleases_APIRequestError(t *testing.T) {
	apiGetRequest = func(callURL string, authorizationHeader string) (any, error) {
		return nil, errors.New("API request failed")
	}

	_, err := ListReleases("https://api.github.com/repos/user/repo/releases", "dummyToken")
	assert.Error(t, err)
	assert.Equal(t, "API request failed", err.Error())
}

func TestListReleases_InvalidResponseFormat(t *testing.T) {
	apiGetRequest = func(callURL string, authorizationHeader string) (any, error) {
		return "invalid response", nil
	}

	_, err := ListReleases("https://api.github.com/repos/user/repo/releases", "dummyToken")
	assert.Error(t, err)
	assert.Equal(t, apimsg.ErrReturn, err)
}

func TestListReleases_NoTagName(t *testing.T) {
	apiGetRequest = func(callURL string, authorizationHeader string) (any, error) {
		return []any{
			map[string]any{"no_tag_name": "v1.0.0"},
		}, nil
	}

	_, err := ListReleases("https://api.github.com/repos/user/repo/releases", "dummyToken")
	assert.Error(t, err)
	assert.Equal(t, apimsg.ErrReturn, err)
}

func TestListReleases_MultiplePages(t *testing.T) {
	apiGetRequest = func(callURL string, authorizationHeader string) (any, error) {
		if callURL == "https://api.github.com/repos/user/repo/releases?page=1" {
			return []any{
				map[string]any{"tag_name": "v1.0.0"},
			}, nil
		} else if callURL == "https://api.github.com/repos/user/repo/releases?page=2" {
			return []any{
				map[string]any{"tag_name": "v1.1.0"},
			}, nil
		}
		return nil, errors.New("not found")
	}

	releases, err := ListReleases("https://api.github.com/repos/user/repo/releases", "dummyToken")
	assert.NoError(t, err)
	assert.Equal(t, []string{"v1.0.0", "v1.1.0"}, releases)
}