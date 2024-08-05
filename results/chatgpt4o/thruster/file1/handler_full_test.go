package internal

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/klauspost/compress/gzhttp"
)

type mockCache struct{}

func (c *mockCache) Get(key string) ([]byte, bool) { return nil, false }
func (c *mockCache) Set(key string, value []byte)  {}

func TestNewHandler_ValidInput(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    &mockCache{},
		maxCacheableResponseBody: 1024,
		maxRequestBody:           2048,
		targetUrl:                targetURL,
		xSendfileEnabled:         true,
	}

	handler := NewHandler(options)
	req := httptest.NewRequest("GET", "http://example.com", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}
}

func TestNewHandler_NilTargetURL(t *testing.T) {
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    &mockCache{},
		maxCacheableResponseBody: 1024,
		maxRequestBody:           2048,
		targetUrl:                nil,
		xSendfileEnabled:         true,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic due to nil targetUrl")
		}
	}()

	NewHandler(options)
}

func TestNewHandler_NoCache(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    nil,
		maxCacheableResponseBody: 1024,
		maxRequestBody:           2048,
		targetUrl:                targetURL,
		xSendfileEnabled:         true,
	}

	handler := NewHandler(options)
	req := httptest.NewRequest("GET", "http://example.com", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}
}

func TestNewHandler_MaxRequestBodyExceeded(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    &mockCache{},
		maxCacheableResponseBody: 1024,
		maxRequestBody:           1,
		targetUrl:                targetURL,
		xSendfileEnabled:         true,
	}

	handler := NewHandler(options)
	req := httptest.NewRequest("POST", "http://example.com", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("expected status RequestEntityTooLarge, got %v", rr.Code)
	}
}

func TestNewHandler_XSendfileDisabled(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    &mockCache{},
		maxCacheableResponseBody: 1024,
		maxRequestBody:           2048,
		targetUrl:                targetURL,
		xSendfileEnabled:         false,
	}

	handler := NewHandler(options)
	req := httptest.NewRequest("GET", "http://example.com", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}
}

func TestNewHandler_GzipHandler(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    &mockCache{},
		maxCacheableResponseBody: 1024,
		maxRequestBody:           2048,
		targetUrl:                targetURL,
		xSendfileEnabled:         true,
	}

	handler := NewHandler(options)
	handler = gzhttp.GzipHandler(handler)

	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}

	if rr.Header().Get("Content-Encoding") != "gzip" {
		t.Errorf("expected Content-Encoding gzip, got %v", rr.Header().Get("Content-Encoding"))
	}
}

func TestNewHandler_BadGatewayPage(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	options := HandlerOptions{
		badGatewayPage:           "/bad-gateway",
		cache:                    &mockCache{},
		maxCacheableResponseBody: 1024,
		maxRequestBody:           2048,
		targetUrl:                targetURL,
		xSendfileEnabled:         true,
	}

	handler := NewHandler(options)
	req := httptest.NewRequest("GET", "http://example.com/bad-gateway", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}

	if rr.Body.String() != "/bad-gateway" {
		t.Errorf("expected bad gateway page content, got %v", rr.Body.String())
	}
}