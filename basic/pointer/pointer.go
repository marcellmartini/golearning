package pointer

import (
	"fmt"

	"basic/tools"
)

type person struct {
	nome      string
	sobrenome string
	idade     int
}

// Examples of pointer
func Examples() {
	tools.HeaderOutPut("pointerExamples")

	a := 1
	// Endereço de a
	fmt.Println(&a)

	marcell := person{
		nome:      "Marcell",
		sobrenome: "Martini",
		idade:     41,
	}
	fmt.Println(marcell)

	changeMe(&marcell)
	fmt.Println(marcell)

}

func changeMe(p *person) {
	(*p).idade++
}
