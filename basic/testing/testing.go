package testing

import (
	"fmt"

	"basic/tools"

	"basic/testing/dog"
)

type canine struct {
	name string
	age  int
}

func example1() {
	tools.HeaderOutPut("Example 1")

	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))

}

// Examples call examples functions
func Examples() {
	tools.HeaderOutPut("Test Examples")

	example1()
}
