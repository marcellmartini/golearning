package cases

/*
 * Complete the 'countingSort' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func countingSort(arr []int32) []int32 {
	xFrequency := make([]int32, 100)

	for _, value := range arr {
		xFrequency[value]++
	}

	return xFrequency
}

func Case3() {}
