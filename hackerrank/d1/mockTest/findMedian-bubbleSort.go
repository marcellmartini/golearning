package mocktest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'findMedianB' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func findMedianB(arr []int32) int32 {
	result := bubleSort(arr)

	return result[(len(arr)-1)/2]

}

func bubleSort(arr []int32) []int32 {
	check := len(arr) - 1
	changed := 0

	for i := 0; i < len(arr); i++ {
		changed = 0
		for j := 0; j < check-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changed = 1
			}
		}

		if changed == 0 {
			return arr
		}
	}

	return nil
}

func mainB() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := findMedianB(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
