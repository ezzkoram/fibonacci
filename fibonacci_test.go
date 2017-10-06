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

func TestFibonacciSliceRecursive(t *testing.T) {
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
		actual := FibonacciSliceRecursive(tt.n)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("fibonacciSlice(%d): expected %v, actual %v", tt.n, tt.expected, actual)
		}
	}
}

func benchmarkFibonacciSliceN(n uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciSlice(n)
	}
}
func BenchmarkFibonacciSlice1(b *testing.B)  { benchmarkFibonacciSliceN(1, b) }
func BenchmarkFibonacciSlice5(b *testing.B)  { benchmarkFibonacciSliceN(5, b) }
func BenchmarkFibonacciSlice15(b *testing.B) { benchmarkFibonacciSliceN(15, b) }

func benchmarkFibonacciSliceRecursiveN(n uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciSliceRecursive(n)
	}
}
func BenchmarkFibonacciSliceRecursive1(b *testing.B)  { benchmarkFibonacciSliceRecursiveN(1, b) }
func BenchmarkFibonacciSliceRecursive5(b *testing.B)  { benchmarkFibonacciSliceRecursiveN(5, b) }
func BenchmarkFibonacciSliceRecursive15(b *testing.B) { benchmarkFibonacciSliceRecursiveN(15, b) }
