package d3

import (
	"fmt"
	"unicode"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

const (
	qntLetters = 26
)

func caesarCipher(s string, k int32) string {
	output := make([]byte, len(s))

	for key, v := range s {
		if unicode.IsLetter(rune(s[key])) {
			ascii := 97

			if IsUpper(rune(s[key])) {
				ascii = 65
			}

			output[key] = byte(((int32(v) + k - int32(ascii)) % qntLetters) + int32(ascii))
		} else {
			output[key] = s[key]
		}
	}

	return string(output)
}

func IsUpper(s rune) bool {
	if !unicode.IsUpper(s) && unicode.IsLetter(s) {
		return false
	}

	return true
}

func main() {
	fmt.Println(caesarCipher("middle-Outz", 2))

	caesarCipher("middle-Outz", 2)

}
