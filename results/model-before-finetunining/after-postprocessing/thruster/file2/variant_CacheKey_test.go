package internal

import (
    "testing"
    "net/http"
    "strings"
    "github.com/stretchr/testify/assert"
)

type Variants struct {
    Variant1 *http.Request
    Variant2 *http.Request
    Variants3 []string
}

func TestCacheKey(t *testing.T) {
    // Test with valid input
    t.Run("ValidInput", func(t *testing.T) assert.True {
        req1 := &http.Request{
            Method: "GET",
            URL:    http.ParseRequest("http://example.com/path?query=value"),
            Header:  make(map[string]string),
        }
        req2 := &http.Request{
		    Method: "POST",
		    URL:    http.ParseRequest("/submit"),
		    Header:  make(map[interface{}]interface{}),
		}
        variants := Variants{
            Variant1: req1,
            Variant2: req2,
        }

        key1 := variants.Variant1.CacheKey()
        key2 := variants.Variant2.CacheKey()

        assert.NotNil(t, key1)
        assert.NotNil(key2)
    })
}