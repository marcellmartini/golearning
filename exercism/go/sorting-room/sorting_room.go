package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	number := 0

	switch v := fnb.(type) {
	case FancyNumber:
		number, _ = strconv.Atoi(v.Value())
	}

	return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %v.0", ExtractFancyNumber(fnb))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	describeReturn := "Return to sender"

	switch v := i.(type) {
	case int:
		describeReturn = DescribeNumber(float64(v))
	case float64:
		describeReturn = DescribeNumber(v)
	case NumberBox:
		describeReturn = DescribeNumberBox(v)
	case FancyNumberBox:
		describeReturn = DescribeFancyNumberBox(v)
	}

	return describeReturn
}
