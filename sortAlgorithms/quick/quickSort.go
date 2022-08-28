// ,You,can,edit,this,code!
// ,Click,here,and,start,typing.
package quick

import (
	"fmt"
	"sortAlgorithms/utils"
)

func Sort() {
	arr := utils.GetArr2()

	utils.PrintHeader("Quick Sort")

	fmt.Println("Initial:\n", arr)

	quickSort(&arr, 0, len(arr)-1)

	fmt.Println("Result:\n", arr)
}

func quickSort(arr *[]int32, low, high int) {

	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)

		quickSort(arr, pi+1, high)
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
