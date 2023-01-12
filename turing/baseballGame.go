package turring

import (
	"strconv"
)

func callPoints(ops []string) int {
	var res int = 0
	record := []int{}

	for _, v := range ops {
		switch v {
		case "C":
			record = record[:len(record)-1]
		case "D":
			record = append(record, record[len(record)-1]*2)
		case "+":
			record = append(record, record[len(record)-1]+record[len(record)-2])
		default:
			if score, err := strconv.Atoi(v); err == nil {
				record = append(record, score)
			}
		}
	}

	for _, v := range record {
		res += v
	}

	return res
}
