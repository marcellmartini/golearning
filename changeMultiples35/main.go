/*
Write a program that prints the numbers from 1 to 100 and for multiples of '3'
print “M3” instead of the number and for the multiples of '5' print “M5”.

For multiples of both 3 and 5 print both “M3” and “M5” in one line.

Example of program output:
1
2
M3
4
M5
M3
7
8
M3
M5
11
M3
13
14
M3M5
16
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for i := 1; i <= 16; i++ {
		m3m5(os.Stdout, i)
	}

}

func m3m5(output io.Writer, number int) {
	switch {
	case number%3 == 0 && number%5 == 0:
		fmt.Fprintln(output, "M3M5")

	case number%3 == 0:
		fmt.Fprintln(output, "M3")

	case number%5 == 0:
		fmt.Fprintln(output, "M5")

	default:
		fmt.Fprintln(output, number)
	}
}
