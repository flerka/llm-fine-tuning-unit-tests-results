package download

import (
	"errors"
	"testing"
)

func TestApplyUrlTransformer_Success(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		return url + "/transformed", nil
	}
	baseURLs := []string{"http://example.com", "http://test.com"}
	expected := []string{"http://example.com/transformed", "http://test.com/transformed"}

	result, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], v)
		}
	}
}

func TestApplyUrlTransformer_EmptyBaseURLs(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		return url + "/transformed", nil
	}
	baseURLs := []string{}
	expected := []string{}

	result, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != len(expected) {
		t.Errorf("expected %d results, got %d", len(expected), len(result))
	}
}

func TestApplyUrlTransformer_TransformerError(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		return "", errors.New("transformer error")
	}
	baseURLs := []string{"http://example.com"}

	_, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	if err.Error() != "transformer error" {
		t.Errorf("expected transformer error, got %v", err)
	}
}

func TestApplyUrlTransformer_InvalidURL(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		_, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		return url, nil
	}
	baseURLs := []string{":://invalid-url"}

	_, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestApplyUrlTransformer_SingleBaseURL(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		return url + "/transformed", nil
	}
	baseURLs := []string{"http://example.com"}
	expected := []string{"http://example.com/transformed"}

	result, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 1 || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestApplyUrlTransformer_MultipleBaseURLs(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		return url + "/transformed", nil
	}
	baseURLs := []string{"http://example.com", "http://test.com", "http://sample.com"}
	expected := []string{"http://example.com/transformed", "http://test.com/transformed", "http://sample.com/transformed"}

	result, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], v)
		}
	}
}

func TestApplyUrlTransformer_EmptyTransformer(t *testing.T) {
	urlTransformer := func(url string) (string, error) {
		return "", nil
	}
	baseURLs := []string{"http://example.com"}
	expected := []string{""}

	result, err := ApplyUrlTranformer(urlTransformer, baseURLs...)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 1 || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}