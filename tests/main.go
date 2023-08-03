package main

import (
	"strconv"
)

func main() {
	addN := func(m int) func(n int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addN(5)
	result := addFive(6)

	addSeven := addN(7)
	result2 := addSeven(6)

	println(result)
	//5 + 6 must print 11
	println(result2)
	//7 + 6 must print 13

	println("Fibo: " + strconv.Itoa(Fibonnaci(5)))
}

// Fibonnaci return the fibo value of n value
func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}
