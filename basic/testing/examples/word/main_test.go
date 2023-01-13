package word

import (
	"fmt"
	"testing"

	"basic/testing/examples/quote"
)

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("Marcell"))
	// Output:
	// 1
}

func TestUseCount(t *testing.T) {
	values := []struct {
		value  string
		expect map[string]int
	}{
		{
			value: "Marcell Martini",
			expect: map[string]int{
				"Marcell": 1,
				"Martini": 1,
			},
		},
		{
			value: "one two two three three three",
			expect: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
	}

	for _, v := range values {
		got := UseCount(v.value)

		for k := range got {
			if got[k] != v.expect[k] {
				t.Errorf("\nFor: %#v\nExpected: %v\nGot: %v", k, v.expect[k], got[k])
			}
		}
	}
}

func TestCount(t *testing.T) {
	got := Count("Marcell")
	if got != 1 {
		t.Errorf("\nFor: %#v\nExpected: %v\nGot: %v", "Marcell", 1, got)
	}
}
