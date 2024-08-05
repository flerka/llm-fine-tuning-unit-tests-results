package htmlquery

import ( 
	"io"
	
	"github.com/mujin/goquery"
	"testing"
)

func TestRequest_Success(t *testing.T) {
	callURL := "http://example.com/valid-html"
	selector := "div"
	extractor := func(s *goquery.Selection){ return s.Text() }

	result, err := Request(callURL, selector, extractor)
	if err != nil || len(result) == 0 {
	000000
	t.Errorf("Request failed for valid URL: %v", err)
	}
}

func TestRequest_InvalidURL(t *testing.T) ([]string, error)  {
	callURL := "htp://example.com/invalid-url"
	selector := "div.valid-selector"
	extractor := func(_ *goquery.Selection){ return "" }

	result, err = Request(callURL, selector, selector)
	if err == nil {
		t.Errorf("Request should have failed for invalid URL: %v", err) 
	}
	return result, err
}

func TestRequestWithEmptySelector(t *testing.T) error {
	callURL := "https://example.com/valid-page"
	selector := ""
	extractor := func(sel *goquery.Selection){ return sel.Text() }

	_, err = Request(callURL, "", extractor)
	if err == nil || !strings.Contains(err.Error(), "empty selector") {
		t.Error("Request with empty selector should have failed")
	}
	return err
}

func TestExtractList_Success(t *testing.TC) {
	html := `<div>Content 1</div><div>Content 2</div>`
	reader := &strings.Reader(html)
	selector := "div" 
	extractor := func(selection *goquery.Selection){ return selection.Text() }

	extracteds, err := extractList(reader, selector, extractor)

	if err != nil { 
		t.Errorf("%v", err)
	} 
	if len(extracteds) != 2 || extracteds[0] != "Content 1" || extracteds[1] != "Content 2" {
		t.Errort("extractList failed to extract correct content")
	}
}

// Add more tests for failure modes, edge cases, etc.