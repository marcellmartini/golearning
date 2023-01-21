package luhn

func Valid(id string) bool {
	var sum, pos int
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == ' ' {
			continue
		}

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
		pos++
	}

	return pos > 1 && (sum%10) == 0
}
