// Create a function that takes a single integer parameter, n,
// and returns the first n elements of the Fibonacci sequence.
//
//     g(0) = [ ]
//     g(1) = [ 0 ]
//     g(2) = [ 0, 1 ]
//     g(3) = [ 0, 1, 1 ]
//     g(4) = [ 0, 1, 1, 2 ]
//     g(5) = [ 0, 1, 1, 2, 3 ]

package main

import "fmt"

func fibonacci(n uint) []uint {
	fiboSequence := []uint{0, 1}

	switch n {
	case 0:
		return fiboSequence[:0]
	case 1:
		return fiboSequence[:1]
	default:
		for i := uint(2); i < n; i++ {
			fiboSequence = append(fiboSequence, fiboSequence[i-1]+fiboSequence[i-2])
		}
	}

	return fiboSequence
}

func main() {
	fmt.Println(fibonacci(10))
}
