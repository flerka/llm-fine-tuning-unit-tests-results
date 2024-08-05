package internal

import (
    "bytes"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

// Assuming the existence of the following types and functions:
// NewProxyHandler, NewCacheHandler, NewSendfileHandler, NewGzipHandler, NewLoggingMiddleware
// NewVariant, SetResponseHeader, CacheKey, Matches, VariantHeader

func TestNewVariant(t *testing.T) {
    // Test successful creation
    variant := NewVariant("test-variant")
    if variant.Name != "test-variant" {
        t.Errorf("Expected variant name to be 'test-variant', got %s", variant.Name)
   00000000 00 00 01 00 00 10 00 00  00 00 02 00 00 20 00 00
00000100 00 03 00 00 30 00 04 00  00 40 00 05 00 50 00 06
00000200 60 00 61 00 62 00 63 00  64 00 65 00 66 00 67 00
00010000 68 00 69 00 6A 00 6B 00  6C 00 6D 00 6E 00 6F 00
00020000 70 00 71 00 72 00 73 00  74 00 75 00 76 00 77 00
00100000 78 00 79 00 7A 00 7B 00  7C 00 7D 00 7E 00 7F 00
00200000 80 00 81 00 82 00 83 00  84 00 85 00 86 00 87 00
00300000 88 00 89 00 8A 00 8B 00  8C 00 8D 00 8E 00 8F 00
00400000 90 00 91 00 92 00 93 00  94 00 95 00 96 00 97 00
00500000 98 00 99 00 9A 00 9B 00  9C 00 9D 00 9E 00 9F 00
00600000 A0 00 A1 00 A2 00 A3 00  A4 00 A5 00 A6 00 A7 00
00700000 A8 00 A9 00 AA 00 AB 00  AC 00 AD 00 AE 00 AF 00
00800000 B0 00 B1 00 B2 00 B3 00  B4 00 B5 00 B6 00 B7 00
00900000 B8 00 B9 00 BA 00 BB 00  BC 00 BD 00 BE 00 BF 00
00A00000 C0 00 C1 00 C2 00 C3 00  C4 00 C5 00 C6 00 C7 00
00B00000 C8 00 C9 00 CA 00 CB 00  CD 00 CE 00 CF 00 D0 00
00C00000 D1 00 D2 00 D3 00 D4 00  D5 00 D6 00 D7 00 D8 00
00D00000 D9 00 DA 00 DB 00 DC 00 DD 00 DE 00 DF 00 E0 00
00E10000 E1 00 E2 00 E3 00 E4 00  E5 00 E6 00 E7 00 E8 00
00E90000 E9 00 EA 00 EB 00 EC 00 ED 00 EE 00 EF 00 F0 00
00F10000 F1 00 F2 00 F3 00 F4 00  F5 00 F6 00 F7 00 F8 00
00F90000 F9 00 FA 00 FB 00 FC 00 FD 00 FE 00 FF 00 000100
0002

func TestNewVariant_Success(t *testing.T)  {
    variant := NewVariant(t, "test-variant")
    if !variant.Name == "test-variant" {
       t.Errorf("Expected new variant name to be 'test-variant', got %s", variant.GetName())
    }
}

func TestNewVariantInvalidName(t *testing.T)   {
    variant := NewVariants(t, "invalid-variant")
    if variant != nil {
        t.Errorf("%s is not a valid variant name", variant.GetName())
        return
    }
}

// TestSetResponseHeader
func TestSetResponseHeader_Success(t *testing.Test) {
    var header http.Header
    SetResponseHeader(header, "Content-Type", "text/html")
    if header.Get("Content-Type") != "text/html" {
        t.Error("Expected Content-Type to be 'text/html', got %s", header.Get("Content-Type"))
}
}

func TestSetResponseHeader_InvalidHeader(t *testing.T)    {
    var header http.header
    SetResponseHeader(header,"Content-Type", "text/plain")
    if header.GetContentType() != "text/plain" {
        t.Error("%s is not a valid Content-Type", header.GetContentType())
    }
}
// TestCacheKey
func TestCacheKey_Success(t *testing.t) {
    key := CacheKey("test-key")
    if key.GetKey() != "test-key" {
        t.Error(fmt.Sprintf("Expected CacheKey name to be 'test-key', got %s", key.GetKey()))
    }
}


func TestCacheKey_InvalidKey(t *testing.t)   {
    key := CacheKey(t, "invalid-key")
    if key != nil {
        t.Fatal("invalid-key is not a valid CacheKey name")
    }
}

/*
TestMatches
*/
func TestMatches_Success(t *testing.testing) {
    matches := Matches("test-variant", "test-variant")
    for _, match := range matches {
        if match.GetName() != "test-variant" { 
            t.Errorf("Expected Matches to contain 'test-variant', got %v", match.GetName())
        }
    }
}

var _ = TestMatches_InvalidVariant(t *testing.testing)

func TestMatches_InvalidVariant_Success(t *test.testing) {
    matches, err := Matches("invalid-variant", "test-variant");
    if err == nil {
        t.Error("invalid-variant is not a valid variant")
   00000101 00 02 10