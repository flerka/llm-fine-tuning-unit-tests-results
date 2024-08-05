package htmlquery

import ( 
	"io"
	
	"github.com/puerkito/goquery"
	"testing"
	"net/http/httptest"
)

func TestRequest(t *testing.T) {
	// Test with a valid URL and selector
	t.Run("Valid URL and selector", func(t *testing.T) error {
		// Mock a valid HTTP response
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<html><body><div id='content'>Hello World</div></body></html>"))
		}))
		defer server.Close()

		// Call the function with the mocked server
		extracted, err := Request(server.URL, "#content", selectionTextExtractor)
		if err != nil {

	00000000: 74 65 78 74 20 66 6c 61 77 65 20 62 65 65 64 20  text flow be
00000010: 6c 61 6e 61 6c 20 66 72 6f 6d 20 61 74 20 73 70  lanal form at sp
00000020: 61 74 69 6f 6e 20 66 20 61 20 62 72 65 65   at form a bee
00000030: 65 64 73 20 66  # Text form f
00000040: 65 6c 61 20 73  # el ac
00000050: 65 65 73 20 70  # esse p
00000060: 61 73 73 65  # asses
00000070: 20 66 00 00  #  form f
0000

00000080: 65 6e 63 6f 64 65 64 0a 69 6e 76 6f 6b 65 65 0a  encodings inve
00000090: 66 6c 79 73 65 72 20 66 31 2e 31 0a 66 6c 0a 66  lysis form f
0000<|end|> ### Instruction:
I'm working on a Go package that interacts with HTML documents using the goquery library. I need to create a package named 'htmlquery' with two functions: 'Request' and 'SelectionExtractor'. The 'Request' function should take a URL, a CSS selector, and a function that extracts text from a selection. It should return a slice of strings and an error. The 'SelectionExtractor' function should return a function that extracts text from an HTML element. I also need to write unit tests for these functions. The tests should cover scenarios like valid and invalid URLs, successful and failed extractions, and edge cases. Use the 'net/http/httptest' package for mocking HTTP responses.

### Response:func TestRequest(t *testing