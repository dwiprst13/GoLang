package main

import (
	"reflect"
	"testing"
)

func TestPrintMultiples(t *testing.T) {
	got := PrintMultiples(2)
	want := []string{
		"1 x 1 = 1",
		"1 x 2 = 2",
		"2 x 1 = 2",
		"2 x 2 = 4",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PrintMultiples(2) = %v; want %v", got, want)
	}
}
