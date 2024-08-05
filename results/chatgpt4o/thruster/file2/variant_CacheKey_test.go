package internal

import (
	"hash/fnv"
	"net/http"
	"testing"
)

func TestCacheKey_Success(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept"},
	}

	expectedHash := fnv.New64()
	expectedHash.Write([]byte(req.Method))
	expectedHash.Write([]byte(req.URL.Path))
	expectedHash.Write([]byte(req.URL.Query().Encode()))
	expectedHash.Write([]byte("Accept=" + req.Header.Get("Accept")))

	expectedCacheKey := CacheKey(expectedHash.Sum64())

	if v.CacheKey() != expectedCacheKey {
		t.Errorf("expected %v, got %v", expectedCacheKey, v.CacheKey())
	}
}

func TestCacheKey_DifferentMethods(t *testing.T) {
	req1, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req2, _ := http.NewRequest("POST", "http://example.com/path", nil)

	v1 := Variant{r: req1, headerNames: []string{}}
	v2 := Variant{r: req2, headerNames: []string{}}

	if v1.CacheKey() == v2.CacheKey() {
		t.Errorf("expected different cache keys for different methods")
	}
}

func TestCacheKey_DifferentPaths(t *testing.T) {
	req1, _ := http.NewRequest("GET", "http://example.com/path1", nil)
	req2, _ := http.NewRequest("GET", "http://example.com/path2", nil)

	v1 := Variant{r: req1, headerNames: []string{}}
	v2 := Variant{r: req2, headerNames: []string{}}

	if v1.CacheKey() == v2.CacheKey() {
		t.Errorf("expected different cache keys for different paths")
	}
}

func TestCacheKey_DifferentQueries(t *testing.T) {
	req1, _ := http.NewRequest("GET", "http://example.com/path?query=1", nil)
	req2, _ := http.NewRequest("GET", "http://example.com/path?query=2", nil)

	v1 := Variant{r: req1, headerNames: []string{}}
	v2 := Variant{r: req2, headerNames: []string{}}

	if v1.CacheKey() == v2.CacheKey() {
		t.Errorf("expected different cache keys for different query parameters")
	}
}

func TestCacheKey_DifferentHeaders(t *testing.T) {
	req1, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req1.Header.Set("Accept", "application/json")
	req2, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req2.Header.Set("Accept", "application/xml")

	v1 := Variant{r: req1, headerNames: []string{"Accept"}}
	v2 := Variant{r: req2, headerNames: []string{"Accept"}}

	if v1.CacheKey() == v2.CacheKey() {
		t.Errorf("expected different cache keys for different header values")
	}
}

func TestCacheKey_EmptyHeaderNames(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)

	v := Variant{
		r:           req,
		headerNames: []string{},
	}

	expectedHash := fnv.New64()
	expectedHash.Write([]byte(req.Method))
	expectedHash.Write([]byte(req.URL.Path))
	expectedHash.Write([]byte(req.URL.Query().Encode()))

	expectedCacheKey := CacheKey(expectedHash.Sum64())

	if v.CacheKey() != expectedCacheKey {
		t.Errorf("expected %v, got %v", expectedCacheKey, v.CacheKey())
	}
}

func TestCacheKey_MultipleHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-http-client")

	v := Variant{
		r:           req,
		headerNames: []string{"Accept", "User-Agent"},
	}

	expectedHash := fnv.New64()
	expectedHash.Write([]byte(req.Method))
	expectedHash.Write([]byte(req.URL.Path))
	expectedHash.Write([]byte(req.URL.Query().Encode()))
	expectedHash.Write([]byte("Accept=" + req.Header.Get("Accept")))
	expectedHash.Write([]byte("User-Agent=" + req.Header.Get("User-Agent")))

	expectedCacheKey := CacheKey(expectedHash.Sum64())

	if v.CacheKey() != expectedCacheKey {
		t.Errorf("expected %v, got %v", expectedCacheKey, v.CacheKey())
	}
}

func TestCacheKey_NoHeaderValues(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/path", nil)

	v := Variant{
		r:           req,
		headerNames: []string{"Non-Existent-Header"},
	}

	expectedHash := fnv.New64()
	expectedHash.Write([]byte(req.Method))
	expectedHash.Write([]byte(req.URL.Path))
	expectedHash.Write([]byte(req.URL.Query().Encode()))
	expectedHash.Write([]byte("Non-Existent-Header="))

	expectedCacheKey := CacheKey(expectedHash.Sum64())

	if v.CacheKey() != expectedCacheKey {
		t.Errorf("expected %v, got %v", expectedCacheKey, v.CacheKey())
	}
}