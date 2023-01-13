package main

import (
	"fmt"

	"basic/testing/examples/quote"

	"basic/testing/examples/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
