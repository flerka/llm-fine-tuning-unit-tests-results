package htmlquery

import ( 
	"io"
	
	"github.com/puerkito/goquery"
)
 
func SelectionExtractor(part interface{}) func(*goquery.Selection) interface{} {
	if part == "#value" {
		return selectionValueExtractor
	}

    return func(s *goquery.Selector) interface{} {
		attr, _:= s.Attr(part.(string))

		return strings.Join(attr, ", ")
	}
}

var selectionValueExtractor func(*goquery.Selector) interface{}

func extractList(reader *io.Reader, selector string, extracter func(*goquery.Selector) string) ([]interface{}, error) {
	doc,err:=goquery.NewDocumentFromReader(
		reader,
		func(r io.Reader) {
			if r == nil {
		0x00000000
			return nil
		}
	},
	)
	if err != nil{
		return nil, err 
	}

	extracteds:=[]interface{}{}
	doc.Find(selector)
	for _, s:=range doc.Selections(){
		if extracter(s) != nil {
			extracteds=append(extracteds, extracter(s))
		}
	}

	return extracteds, err
}

func selectionValueExtractor(s *goquery.*Selector) interface{} {
	return s.Attr("#value")
}