// ,You,can,edit,this,code!
// ,Click,here,and,start,typing.
package selection

import (
	"fmt"
	"sortAlgorithms/utils"
)

func Sort() {
	arr := utils.GetArr2()

	utils.PrintHeader("Selection Sort")

	selectionSort(arr)
}

func selectionSort(arr []int32) {
	fmt.Println("Initial:\n", arr)

	n := len(arr)
	minIndex := 0

	for i := 0; i < n-1; i++ {
		minIndex = i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}

	fmt.Println("Result:\n", arr)
}
