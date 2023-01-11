package d3

import "testing"

// type value struct {
// 	string string
// 	shift  int32
// 	expect string
// }

func TestCaesarCipher(t *testing.T) {

	value := struct {
		string string
		shift  int32
		expect string
	}{
		"middle-Outz",
		2,
		"okffng-Qwvb",
	}

	got := caesarCipher("middle-Outz", 2)

	if got != value.expect {
		t.Errorf("\nFor: %#v\nExpected: %v\nGot: %v", value.string, value.expect, got)
	}
}
