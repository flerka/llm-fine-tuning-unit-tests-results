package htmlquery

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestSelectionExtractor_TextExtraction(t *testing.T) {
	extractor := SelectionExtractor("#text")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div>Sample Text</div>`))
	text := extractor(sel.Selection)

	expected := "Sample Text"
	if text != expected {
		t.Errorf("expected %q but got %q", expected, text)
	}
}

func TestSelectionExtractor_AttributeExtraction(t *testing.T) {
	extractor := SelectionExtractor("class")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div class="test-class">Sample Text</div>`))
	attr := extractor(sel.Selection)

	expected := "test-class"
	if attr != expected {
		t.Errorf("expected %q but got %q", expected, attr)
	}
}

func TestSelectionExtractor_EmptyAttribute(t *testing.T) {
	extractor := SelectionExtractor("id")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div>Sample Text</div>`))
	attr := extractor(sel.Selection)

	expected := ""
	if attr != expected {
		t.Errorf("expected %q but got %q", expected, attr)
	}
}

func TestSelectionExtractor_WhitespaceTrimming(t *testing.T) {
	extractor := SelectionExtractor("#text")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div>  Sample Text  </div>`))
	text := extractor(sel.Selection)

	expected := "Sample Text"
	if text != expected {
		t.Errorf("expected %q but got %q", expected, text)
	}
}

func TestSelectionExtractor_NonExistentAttribute(t *testing.T) {
	extractor := SelectionExtractor("nonexistent")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div class="test-class">Sample Text</div>`))
	attr := extractor(sel.Selection)

	expected := ""
	if attr != expected {
		t.Errorf("expected %q but got %q", expected, attr)
	}
}

func TestSelectionExtractor_HTMLContent(t *testing.T) {
	extractor := SelectionExtractor("#text")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div><span>Sample</span> Text</div>`))
	text := extractor(sel.Selection)

	expected := "Sample Text"
	if text != expected {
		t.Errorf("expected %q but got %q", expected, text)
	}
}

func TestSelectionExtractor_SpecialCharacters(t *testing.T) {
	extractor := SelectionExtractor("#text")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div>Special &amp; Characters</div>`))
	text := extractor(sel.Selection)

	expected := "Special & Characters"
	if text != expected {
		t.Errorf("expected %q but got %q", expected, text)
	}
}

func TestSelectionExtractor_EmptyTextNode(t *testing.T) {
	extractor := SelectionExtractor("#text")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div></div>`))
	text := extractor(sel.Selection)

	expected := ""
	if text != expected {
		t.Errorf("expected %q but got %q", expected, text)
	}
}

func TestSelectionExtractor_MultipleAttributes(t *testing.T) {
	extractor := SelectionExtractor("class")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div class="test-class another-class">Sample Text</div>`))
	attr := extractor(sel.Selection)

	expected := "test-class another-class"
	if attr != expected {
		t.Errorf("expected %q but got %q", expected, attr)
	}
}

func TestSelectionExtractor_HTMLTagsInText(t *testing.T) {
	extractor := SelectionExtractor("#text")

	sel := goquery.NewDocumentFromNode(goquery.MustParseHtmlNode(`<div>Text with <b>HTML</b> tags</div>`))
	text := extractor(sel.Selection)

	expected := "Text with HTML tags"
	if text != expected {
		t.Errorf("expected %q but got %q", expected, text)
	}
}