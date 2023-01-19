package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	var countRune = map[rune]int{}

	for _, v := range word {
		if v != ' ' && v != '-' {
			lr := unicode.ToLower(v)
			if countRune[lr] += 1; countRune[lr] > 1 {
				return false
			}
		}
	}

	return true
}
