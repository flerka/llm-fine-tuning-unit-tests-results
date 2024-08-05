package internal

import (
    "testing"
    "net/http"
)

// TestSetResponseHeader_ValidInput tests the SetResponseHeader function with valid input.
func TestSetResponseHeader_ValidInput(t *testing.T) {
    req := &http.Request{
        Method: "GET",
        // ... other request fields
    }

    header := http.Header{
        "Vary": []string{"Accept-Language", "Accept-Encoding"},
    }

    var variant = Variant{r: req, headerNames: []string{}}
    variant.SetResponseHeader(header)

    if !fnv.Equal(header.Get("Vary"), "Accept-Language,Accept-Encoding") {
        t.Errorf("Expected Vary header to be 'Accept-Language,Accept-Encoding', got '%v'", header.Get("Vary"))
   00000000 00 00 01 00 00 3c 00 3c 20 2f 73 75 62 22 3e 0a  |...<< /sub"">..|
00000020 00 3c a3 61 63 63 65 6e 74 61 6c 20 69 6e 20 54  |.<accent in T|
00000200 65 6d 6f 75 6e 69 6f 6e 20 66 6f 72 20 69 73 20  |emoune from is |
00000220 6e 6f 74 65 6e 69
