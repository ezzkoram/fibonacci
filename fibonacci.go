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
