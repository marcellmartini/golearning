// You can edit this code!
// Click here and start typing.
package bubble

import (
	"fmt"
	"sortAlgorithms/utils"
)

func Sort() {
	arr := utils.GetArr2()

	utils.PrintHeader("Bubble Sort")

	bubleSort(arr)
}

func bubleSort(arr []int32) {
	check := len(arr) - 1
	changed := 0

	fmt.Println("Initial:\n", arr)

	for i := 0; i < len(arr); i++ {
		changed = 0

		for j := 0; j < check-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changed = 1
			}
		}

		if changed == 0 {
			fmt.Println("Result:\n", arr)
			return
		}
	}
}
