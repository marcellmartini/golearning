package mocktest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'findMedianQ' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func findMedianQ(arr []int32) int32 {
	quickSort(&arr, 0, len(arr)-1)

	return arr[(len(arr)-1)/2]

}

func quickSort(arr *[]int32, low, high int) {

	if low < high {
		pivot := partition(arr, low, high)

		quickSort(arr, low, pivot-1)

		quickSort(arr, pivot+1, high)
	}
}

func partition(arr *[]int32, low, high int) int {

	pivot := (*arr)[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]

	return i + 1
}

func mainQ() {
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

	result := findMedianQ(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
