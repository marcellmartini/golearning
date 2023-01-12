package errorhandling

import (
	"fmt"
	"testing"

	"basic/tools"
)

// TestSqrt test the sqrt function
func TestSqrt(t *testing.T) {
	tools.HeaderOutPut("exercise 5")

	got, _ := sqrt(4)
	if got != 2 {
		t.Errorf("sqrt(4) = %f; want 2", got)
	}

	fmt.Printf("------\n")
}
