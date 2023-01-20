package diffsquares

import "math"

func SquareOfSum(n int) int {
	var count int

	for i := 0; i < n; i++ {
		count += i + 1
	}
	return int(math.Pow(float64(count), 2))
}

func SumOfSquares(n int) int {
	var count float64

	for i := 0; i < n; i++ {
		count += math.Pow(float64(i+1), 2)
	}
	return int(count)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
