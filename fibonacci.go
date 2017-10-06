package fibonacci

// FibonacciSlice returns n first numbers from fibonacci series
func FibonacciSlice(n uint64) []uint64 {
	var a, b uint64 = 0, 1
	result := make([]uint64, n)

	for i := uint64(0); i < n; i++ {
		a, b = b, a+b
		result[i] = a
	}
	return result
}

func fibonacciSliceReucrsiveHelper(left, a, b uint64) []uint64 {
	if left == 0 {
		return []uint64{}
	}
	return append([]uint64{b}, fibonacciSliceReucrsiveHelper(left-1, b, a+b)...)
}

// FibonacciSliceRecursive returns n first numbers from fibonacci series
// This variant will use recursion instead of a loop to construct the series
func FibonacciSliceRecursive(n uint64) []uint64 {
	return fibonacciSliceReucrsiveHelper(n, 0, 1)
}
