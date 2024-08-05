package internal

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSetResponseHeader_SingleHeader(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "Accept-Encoding")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_MultipleHeaders(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "Accept-Encoding, User-Agent")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding", "User-Agent"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_DuplicateHeaders(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "Accept-Encoding, Accept-Encoding")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_EmptyHeader(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "")

	v.SetResponseHeader(header)

	expected := []string{}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_MixedCaseHeaders(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "accept-encoding, USER-AGENT")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding", "User-Agent"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_LeadingTrailingSpaces(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", " Accept-Encoding , User-Agent ")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding", "User-Agent"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_CommaSeparatedWithSpaces(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "Accept-Encoding , User-Agent")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding", "User-Agent"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_EmptyHeaderKey(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", ",")

	v.SetResponseHeader(header)

	expected := []string{}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_EmptySpacesBetweenCommas(t *testing.T) {
	v := &Variant{}
	header := http.Header{}
	header.Set("Vary", "Accept-Encoding,  , User-Agent")

	v.SetResponseHeader(header)

	expected := []string{"Accept-Encoding", "User-Agent"}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}

func TestSetResponseHeader_NoVaryHeader(t *testing.T) {
	v := &Variant{}
	header := http.Header{}

	v.SetResponseHeader(header)

	expected := []string{}
	if !reflect.DeepEqual(v.headerNames, expected) {
		t.Errorf("expected %v, got %v", expected, v.headerNames)
	}
}