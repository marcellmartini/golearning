package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var count int

	if len(a) != len(b) {
		return -1, errors.New("lengths of sequences are differents")
	}

	for k, v := range a {
		if v != rune(b[k]) {
			count++
		}
	}

	return count, nil
}
