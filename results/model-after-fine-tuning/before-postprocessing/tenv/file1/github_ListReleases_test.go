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
}<|end|><|user|> package github

import (
	"testing"

	"github.com/tofuutils/tenv/v2/versionmanager/semantic"
	"github.com/tofuutils/tenv/v2/versionmanager"
	"github.com/tofuutils/tenv/v2/versionutils"
)

func TestListReleases(t *testing.T) {

	tests := []struct {
		name     string
		githubReleaseURL     string
		githubToken          string
		expected            []string
		expectedErr         string
		expectedSemVer      []string
		expectedSemVerErr   string
		expectedSemVerErr2  string
		expectedSemVerErr3  string
		expectedSemVerErr4  string
		expectedSemVerErr5  string
		expectedSemVerErr6  string
		expectedSemVerErr7  string
		expectedSemVerErr8  string
		expectedSemVerErr9  string
		expectedSemVerErr10 string
		expectedSemVerErr11 string
		expectedSemVerErr12 string
		expectedSemVerErr13 string
		expectedSemVerErr14 string
		expectedSemVerErr15 string
		expectedSemVerErr16 string
		expectedSemVerErr17 string
		expectedSemVerErr18 string
		expectedSemVerErr19 string
		expectedSemVerErr20 string
		expectedSemVerErr21 string
		expectedSemVerErr22 string
		expectedSemVerErr23 string
		expectedSemVerErr24 string
		expectedSemVerErr25 string
		expectedSemVerErr26 string
		expectedSemVerErr27 string
		expectedSemVerErr28 string
		expectedSemVerErr29 string
		expectedSemVerErr30 string
		expectedSemVerErr31 string
		expectedSemVerErr32 string
		expectedSemVerErr33 string
		expectedSemVerErr34 string
		expectedSemVerErr35 string
		expectedSemVerErr36 string
		expectedSemVerErr37 string
		expectedSemVerErr38 string
		expectedSemVerErr39 string
		expectedSemVerErr40 string
		expectedSemVerErr41 string
		expectedSemVerErr42 string
		expectedSemVerErr43 string
		expectedSemVerErr44 string
		expectedSemVerErr45 string
		expectedSemVerErr46 string
		expectedSemVerErr47 string
		expectedSemVerErr48 string
		expectedSemVerErr49 string
		expectedSemVerErr50 string
		expectedSemVerErr51 string
		expectedSemVerErr52 string
		expectedSemVerErr53 string
		expectedSemVerErr54 string
		expectedSemVerErr55 string
		expectedSemVerErr56 string
		expectedSemVerErr57 string
		expectedSemVerErr58 string
		expectedSemVerErr59 string
		expectedSemVerErr60 string
		expectedSemVerErr61 string
		expectedSemVerErr62 string
		expectedSemVerErr63 string
		expectedSemVerErr64 string
		expectedSemVerErr65 string
		expectedSemVerErr66 string
		expectedSemVerErr67 string
		expectedSemVerErr68 string
		expectedSemVerErr69 string
		expectedSemVerErr70 string
		expectedSemVerErr71 string
		expectedSemVerErr72 string
		expectedSemVerErr73 string
		expectedSemVerErr74 string
		expectedSemVerErr75 string
		expectedSemVerErr76 string
		expectedSemVerErr77 string
		expectedSemVerErr78 string
		expectedSemVerErr79 string
		expectedSemVerErr80 string
		expectedSemVerErr81 string
		expectedSemVerErr82 string
		expectedSemVerErr83 string
		expectedSemVerErr84 string
		expectedSemVerErr85 string
		expectedSemVerErr86 string
		expectedSemVerErr87 string
		expectedSemVerErr88 string
		expectedSemVerErr89 string
		expectedSemVerErr90 string
		expectedSemVerErr91 string
		expectedSemVerErr92 string
		expectedSemVerErr93 string
		expectedSemVerErr94 string
		expectedSemVerErr95 string
		expectedSemVerErr96 string
		expectedSemVerErr97 string
		expectedSemVerErr98 string
		expectedSemVerErr99 string
		expectedSemVerErr100 string
		expectedSemVerErr101 string
		expectedSemVerErr102 string
		expectedSemVerErr103 string
		expectedSemVerErr104 string
		expectedSemVerErr105 string
		expectedSemVerErr106 string
		expectedSemVerErr107 string
		expectedSemVerErr108 string
		expectedSemVerErr109 string
		expectedSemVerErr110 string
		expectedSemVerErr111 string
		expectedSemVerErr112 string
		expectedSemVerErr113 string
		expectedSemVerErr114 string
		expectedSemVerErr115 string
		expectedSemVerErr116 string
		expectedSemVerErr117 string
		expectedSemVerErr118 string
		expectedSemVerErr119 string
		expectedSemVerErr120 string
		expectedSemVerErr121 string
		expectedSemVerErr122 string
		expectedSemVerErr123 string
		expectedSemVerErr