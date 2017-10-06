package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacciSlice(t *testing.T) {
	tests := []struct {
		n        uint64
		expected []uint64
	}{
		{0, []uint64{}},
		{1, []uint64{1}},
		{2, []uint64{1, 1}},
		{4, []uint64{1, 1, 2, 3}},
		{7, []uint64{1, 1, 2, 3, 5, 8, 13}},
		{9, []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, tt := range tests {
		actual := FibonacciSlice(tt.n)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("fibonacciSlice(%d): expected %v, actual %v", tt.n, tt.expected, actual)
		}
	}
}
