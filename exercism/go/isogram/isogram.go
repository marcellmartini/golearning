package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	lw := strings.ToLower(word)

	for _, rune := range lw {
		if rune != ' ' && rune != '-' {
			if strings.Count(lw, string(rune)) > 1 {
				return false
			}
		}
	}

	return true
}
