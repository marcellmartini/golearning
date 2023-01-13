package lonelyinteger

/*
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func lonelyinteger(a []int32) int32 {
	var result int32 = 0
	xInt := [101]int32{}

	for _, num := range a {
		xInt[num]++
	}

	for key, value := range xInt {
		if value == 1 {
			result = int32(key)
		}
	}

	return result
}

func Case1() int32 {
	return lonelyinteger([]int32{1, 1, 2})
}
