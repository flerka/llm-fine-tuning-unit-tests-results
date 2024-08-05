package internal

import (
	"net/http"
	"reflect"
	"testing"
)

func TestVariantHeader_EmptyHeaders(t *testing.T) {
	r := &http.Request{Header: http.Header{}}
	v := Variant{r: r, headerNames: []string{}}

	got := v.VariantHeader()
	want := http.Header{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}

func TestVariantHeader_SingleHeader(t *testing.T) {
	r := &http.Request{Header: http.Header{"X-Test": []string{"value"}}}
	v := Variant{r: r, headerNames: []string{"X-Test"}}

	got := v.VariantHeader()
	want := http.Header{"X-Test": []string{"value"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}

func TestVariantHeader_MultipleHeaders(t *testing.T) {
	r := &http.Request{Header: http.Header{
		"X-Test-1": []string{"value1"},
		"X-Test-2": []string{"value2"},
	}}
	v := Variant{r: r, headerNames: []string{"X-Test-1", "X-Test-2"}}

	got := v.VariantHeader()
	want := http.Header{
		"X-Test-1": []string{"value1"},
		"X-Test-2": []string{"value2"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}

func TestVariantHeader_HeaderNotPresent(t *testing.T) {
	r := &http.Request{Header: http.Header{"X-Test": []string{"value"}}}
	v := Variant{r: r, headerNames: []string{"X-Not-Present"}}

	got := v.VariantHeader()
	want := http.Header{"X-Not-Present": []string{""}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}

func TestVariantHeader_MixedHeaders(t *testing.T) {
	r := &http.Request{Header: http.Header{
		"X-Test-1":    []string{"value1"},
		"X-Test-2":    []string{"value2"},
		"X-Unrelated": []string{"ignore"},
	}}
	v := Variant{r: r, headerNames: []string{"X-Test-1", "X-Test-2"}}

	got := v.VariantHeader()
	want := http.Header{
		"X-Test-1": []string{"value1"},
		"X-Test-2": []string{"value2"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}

func TestVariantHeader_EmptyRequestHeaders(t *testing.T) {
	r := &http.Request{Header: http.Header{}}
	v := Variant{r: r, headerNames: []string{"X-Test"}}

	got := v.VariantHeader()
	want := http.Header{"X-Test": []string{""}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}

func TestVariantHeader_CaseInsensitive(t *testing.T) {
	r := &http.Request{Header: http.Header{"X-TEST": []string{"value"}}}
	v := Variant{r: r, headerNames: []string{"x-test"}}

	got := v.VariantHeader()
	want := http.Header{"x-test": []string{"value"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VariantHeader() = %v, want %v", got, want)
	}
}