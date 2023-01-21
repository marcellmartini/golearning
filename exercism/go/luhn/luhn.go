package luhn

import (
	"strings"
)

func Valid(id string) bool {
	var sum, pos int
	id = strings.ReplaceAll(id, " ", "")

	for i := len(id) - 1; i >= 0; i, pos = i-1, pos+1 {
		if id[i] < '0' || id[i] > '9' {
			return false
		}

		tempInt := int(id[i] - '0')

		if pos%2 != 0 {
			if tempInt*2 > 9 {
				tempInt = tempInt*2 - 9
			} else {
				tempInt = tempInt * 2
			}
		}
		sum += tempInt
	}

	return pos > 1 && (sum%10) == 0
}
