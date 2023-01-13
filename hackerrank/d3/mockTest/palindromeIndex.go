package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'palindromeIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func palindromeIndex(s string) int32 {
	if len(s) < 2 {
		fmt.Println(len(s))
		return -1
	}

	var l int32 = 0
	r := int32(len(s) - 1)

	if isPal(s, l, r) {
		fmt.Println("isPal")
		return -1
	}

	for l < r {
		if isPal(s, l+1, r) {
			fmt.Println("l")
			return l
		}
		if isPal(s, l, r-1) {
			fmt.Println("r")
			return r
		}
		// if s[l] == s[r] {
		// 	fmt.Println("l + 1")
		// 	return -1
		// }
		l++
		r--
	}

	return 0
}

func isPal(s string, l, r int32) bool {
	if len(s) == 0 || l > r {
		return false
	}

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := palindromeIndex(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
