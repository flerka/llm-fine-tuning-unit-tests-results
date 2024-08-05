package download

import (
    "io"
    "net/http"
    "net/url"
    "testing"
)

func TestBytes_Success(t *testing.T) {
    // Test a successful download
    url := "http://example.com/testfile.txt"
    expected := []byte{...} // expected content of the test file
    result, err := Bytes(url, noTransform)
    if err != nil {
        t.Errorf("Bytes() returned an error: %v", err)
   00000000: 64 6f 63 6b 20 6c 65 66 20 66 6c 61 74 68 20 66  dock life flatfla
00000010: 6c 61 62 65 20 66 72 6f 6d 20 66 20 74 68 65 20  lab from flathe
00000020: 66 6c 70 73 20 66 3a 20 66 4c 61 74 xxxxxx flap: flat
00000030: 66 66 20 3a 20 32 2e 30 30 30 2e 30 2e 20 66  ff : 2.000.0.  flat
00000100: 6c 65 20 32 30 30 0a 66 6c 20 32 66 6c 0a 66  l e 200. flatf
00000110: 6c 70 61 73 73 77 6f 72 6b 20 32 78 3a 20 22 66  lpass "f x: "flat
00000120: 6c 6c 61 22 20 32 0a 66 72 20 32 ffffff  flat 2. fr 2
00000130: 66 20 22 32 2e 2e 2e 32 2e 66 6c 3a 20 46 6c flat "2....2. flat
00000200: 6c 72 61 74 73 20 32 46 6c 61 ffffff  lrats 2Fl a flat
00000300: 66 69 6c 65 0a 66 20 46 72 61 6d 6d 20 32  file. flatRamm 2
00001000: 66 2e 66 2e 2e 66 3a 66 6c ffffff  f.f..f: flat
00001010: 6c ffffff 6c 72 73 73 20 22 22 20 66 0a  ll lrss " "f.
00001020: 66 lf lf lf lf lff lf lf lf lfff lf lf lf lfffflf lf lf
00001030: lf lf lf lfl ffffff lf lf lf llfffflf lf lf lffff
00001040: lf lf lff lfff lf lf lff lfffflf lf lfff lfffflf lf
00010000: lfff lfff lfff lfff lfffflf lfff lfff lfff lff
00010010: lfff lfff lff lfff lfff lfff lf lfff lfff lff

func TestBytes_Failure(t *testing.T) {
    // Test a failure scenario where the URL is invalid
    url := "http://invalid-url.com/testfile.txt"
    _, err := Bytes(url, func(string) (string, error) { return "", err })
    if err == nil {
        t.Errorf("Expected an error for invalid URL, but got none")
    }
}

func TestBytes_EmptyResponse(t *testing.T)  {
    // Test a scenario where the server returns an empty response
    url := "http://testserver.com/empty.txt"
    _, err := Bytes("http://testserver.com/non-existent-file.txt", noTransform)
    if err == nil {
        t.Error("Expected an error for empty response, but got none")
    }
}

func TestBytes_LargeResponse(t *testing.T){
    // Test a large response to ensure the function can handle it
    url := "http://large-file-server.com/large-file.txt"
    result, err := ByBytes(url, noTransform)
    
    if err != nil { 
        t.Errorf("Error while downloading large file: %v", err)
        return
    }
    
    if len(result) == 0 {
        t.Error("Downloaded file is empty")
    }
}