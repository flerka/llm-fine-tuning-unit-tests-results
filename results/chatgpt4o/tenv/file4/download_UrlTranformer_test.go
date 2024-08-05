package download

import (
	"net/url"
	"testing"
)

func TestUrlTranformer_NoRewriteRule(t *testing.T) {
	transformer := UrlTranformer([]string{})
	urlValue := "http://example.com/path"

	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != urlValue {
		t.Errorf("expected %v, got %v", urlValue, result)
	}
}

func TestUrlTranformer_EmptyRewriteRule(t *testing.T) {
	transformer := UrlTranformer([]string{"", ""})
	urlValue := "http://example.com/path"

	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != urlValue {
		t.Errorf("expected %v, got %v", urlValue, result)
	}
}

func TestUrlTranformer_NoMatch(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := "http://example.com/path"

	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != urlValue {
		t.Errorf("expected %v, got %v", urlValue, result)
	}
}

func TestUrlTranformer_Match(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := "http://old.com/path"

	expected := "http://new.com/path"
	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestUrlTranformer_ShortInput(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := "http://ol"

	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != urlValue {
		t.Errorf("expected %v, got %v", urlValue, result)
	}
}

func TestUrlTranformer_SameBaseURL(t *testing.T) {
	transformer := UrlTranformer([]string{"http://same.com", "http://same.com"})
	urlValue := "http://same.com/path"

	expected := "http://same.com/path"
	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestUrlTranformer_InvalidURL(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := ":invalid-url"

	result, err := transformer(urlValue)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if result != ":invalid-url" {
		t.Errorf("expected %v, got %v", ":invalid-url", result)
	}
}

func TestUrlTranformer_TrailingSlash(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := "http://old.com/path/"

	expected := "http://new.com/path/"
	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestUrlTranformer_QueryParameters(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := "http://old.com/path?query=1"

	expected := "http://new.com/path?query=1"
	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestUrlTranformer_EmptyPath(t *testing.T) {
	transformer := UrlTranformer([]string{"http://old.com", "http://new.com"})
	urlValue := "http://old.com/"

	expected := "http://new.com/"
	result, err := transformer(urlValue)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}