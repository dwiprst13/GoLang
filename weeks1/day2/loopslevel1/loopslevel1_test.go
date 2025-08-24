package main

import (
	"reflect"
	"testing"
)

func TestPrintNumbers(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{5, []int{1, 2, 3, 4, 5}},
		{10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		result := PrintNumbers(test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("PrintNumbers(%d) = %v; want %v", test.n, result, test.expected)
		}
	}
}