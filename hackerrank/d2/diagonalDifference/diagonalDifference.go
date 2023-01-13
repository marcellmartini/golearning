package diagonaldifference

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var d1 int32
	var d2 int32

	for i := 0; i < len(arr[0]); i++ {
		d1 += arr[i][i]
		d2 += arr[i][len(arr[0])-1-i]
	}

	result := d1 - d2
	if result < 0 {
		return -result
	}

	return result
}

func Case2() {
}
