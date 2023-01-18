package raindrops

import (
	"strconv"
)

type convert struct {
	factors int
	sound   string
}

var converts = []convert{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) (s string) {
	for k := range converts {
		if number%converts[k].factors == 0 {
			s += converts[k].sound
		}
	}

	if s == "" {
		return strconv.Itoa(number)
	}

	return
}
