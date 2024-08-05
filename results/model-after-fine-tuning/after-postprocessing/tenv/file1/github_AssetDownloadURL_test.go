package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetDownloadURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		tag      string
		searchedAssetNames []string
		githubReleaseURL string
		githubToken      string
		display          func(string)
		want            []string
		wantErr         bool
	}{
		{
			name:     "success",
			tag:      "v1.0.0",
			searchedAssetNames: []string{"asset1", "asset2"},
			githubReleaseURL: "https://api.github.com/repos/tenv/tenv/releases/tags/v1.0.0",
			githubToken:      "token",
			display:          func(string) {},
			want: []string{"https://github.com/tenv/tenv/releases/download/v1.0.0/asset1", "https://github.com/tenv/tenv/releases/download/v1/asset2"},
		},
		{
			name:     "success with empty token",
			tag:      "v1.0.0",
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssetDownloadURL(tt.tag, tt.searchedAssetNames, tt.githubReleaseURL, tt.githubToken, tt.display)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetDownloadURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}