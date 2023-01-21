package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var op2 int
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	j := 0
	for i := len(id) - 1; i >= 0; i-- {
		tempInt, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		if j%2 != 0 {
			if tempInt*2 > 9 {
				tempInt = tempInt*2 - 9
			} else {
				tempInt = tempInt * 2
			}
		}
		op2 += tempInt
		j++
	}
	return (op2 % 10) == 0
}
