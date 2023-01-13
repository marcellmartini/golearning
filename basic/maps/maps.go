package maps

import (
	"fmt"

	"basic/tools"
)

// Examples of map
func Examples() {
	tools.HeaderOutPut("mapExample")

	name := "Silva"

	m := map[string]int{
		"Marcell": 1,
		"Martini": 2,
	}

	h8 := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	// fmt.Println(h8)

	h8[`Marcell`] = []string{`Skydive`, `Swoop`, `4way`}

	delete(h8, `no_dr`)

	for k, names := range h8 {
		fmt.Println(k)
		for _, name := range names {
			fmt.Printf("\t%v\n", name)
		}
	}

	// comma ok idiom
	v, ok := m[name]
	fmt.Println(v)
	fmt.Println(ok)

	// comma ok idiom
	if v, ok := m[name]; ok {
		fmt.Println(v)
	}

	// add element to map
	m["Skydive"] = 3
	fmt.Println(m)

	// delete element from map
	delete(m, "Marcell")
	fmt.Println(m)
	delete(m, "Silva")

	if i, ok := m["Silva"]; ok {
		fmt.Println(i)
	}
}
