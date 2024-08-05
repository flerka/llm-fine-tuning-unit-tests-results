```go
package github

import (
	// ... (other imports)
)

// AssetDownloadURL tests

func TestAssetDownloadURL_ValidInput(t *testing.T) {
	// Mock inputs
	tag := "v1.0.0"
	searchedAssetNames := []string{"asset1", "asset2"}
	githubReleaseURL := "https://api.github.com/repos/owner/repo/releases/tags/v1.0.0" // Replace with a valid URL
	githubToken := "valid_token"
	display := func(string) {} // Replace with a mock display function

	// Expected output
	expectedAssetURLs := []string{
		"https://assets-cdn.github.com/assets/.../asset1.zip",
		"https://assets.../asset2.zip",
	}

	// Call the function
	_, err := AssetDownloadURL(tag, searchedAssetNames, githubReleaseURL, githubToken, display)
	if err != nil && errors.Is(err, apimsg.ErrReturn) {
		t.Errorf("AssetDownloadURL failed with error: %v", err)
	}

	// Verify the output
	actualAssetURLs, err := AssetDownloadURL(
		tag,
		searchedAssetNames,
		githubReleaseURL,
		githubToken,
		display,
	)
	if err != nil || !apimsg.ErrReturn.Is(err) {
		t.Fatalf("AssetDownloadURL returned error: %v", err) // Replace with a mock error
	}

	if !apimsg.ErrReturn {
		for _, expectedURL := range expectedAssetURLs {
			if !apimsg.IsAssetURL(expectedURL) {
				t.Errorf("Expected asset URL %s, got %s", expectedURL, actualAssetURLs[expectedAssetURLs.Index(expectedURL)])
			}
		}
	} else {
		for _, expectedAssetURL := range expectedAssetURLs.WithErrContinue {
			t.Errorf("%s not found in the output", expectedAssetURL)
		}
	}	
}

func TestAssetDownloadURLInvalidTag(t *testing.T) (){
	// Mock inputs
	searchedAssetNames := make(map[string]bool)
	searchedAssetNames["asset1"] = true
	searchedAssetNames["invalid"] = true
	githubReleaseURL := "valid_url"
	githubToken := "valid"
	display := func(url string) {} // Replace with a mock function

	_, err := AssetDownlaodURL("invalid", searchedAssetNames, githubReleaseURL)
	if err == nil {
		t.Error("AssetDownloadURL should fail with invalid tag")
	}
}

// ... (other tests)
```