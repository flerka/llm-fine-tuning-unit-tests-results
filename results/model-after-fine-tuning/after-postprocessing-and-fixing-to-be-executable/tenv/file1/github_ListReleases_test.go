package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReleases(t *testing.T) {
	tests := []struct {
		name     string
		githubReleaseURL string
		githubToken      string
		expected        []string
		expectedErr     string
	}{
		{
			name: "success",
			githubReleaseURL: "https://api.github.com/repos/tofuutils/tenv/releases/tags",
			githubToken:      "token",
			expected:        []string{"v2.0.0", "v2.0.1", "v2.0.2"},
		},
		{
			name: "failure",
			githubReleaseURL: "https://api.github.com/releases/tags",
			githubToken:      "token",
	
			expectedErr: "continue",
		},
		{
			name: "failure",

			githubReleaseURL: "https://api.github.com/releases",
			githubToken:      "token",

			expectedErr: "continue",
		},
		{

			name: "failure",
			githubReleaseURL: "http://api.github.com/releases",
			githubToken:      "invalid",

			expectedErr: "continue",
		},
		
		{
			name: "failure",
			githubToken:      "invalid",
			expectedErr: "continue",
		},
		{
            name: "failure",
			githubReleaseURL: "https://api.github/releases/tags",
			githubToken:      "invalid",
			expectedErr:      "continue",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			releases, err := ListReleases(test.githubReleaseURL, test.githubToken)
			if test.expectedErr != "" {
				assert.Equal(t, test.expectedErr, err.Error())
			} else {
				assert.Equal(t, test.expected, releases)
			}
		})
	}
}