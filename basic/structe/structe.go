package structe

import (
	"fmt"

	"basic/tools"
)

// Examples of struct
func Examples() {
	tools.HeaderOutPut("structExample")

	type pessoa struct {
		nome             string
		sobrenome        string
		saboresFavoritos []string
	}

	eu := pessoa{
		nome:             "marcell",
		sobrenome:        "martini",
		saboresFavoritos: []string{"chocolate", "flocos"},
	}

	ela := pessoa{
		nome:      "vida",
		sobrenome: "viva",
		saboresFavoritos: []string{
			"flocos",
			"morango",
		},
	}

	fmt.Println(eu.nome, eu.sobrenome)
	for _, v := range eu.saboresFavoritos {
		fmt.Printf("\t%s\n", v)
	}

	fmt.Println(ela.nome, ela.sobrenome)
	for _, v := range ela.saboresFavoritos {
		fmt.Printf("\t%s\n", v)
	}

	fmt.Println()
	fmt.Println()

	m := map[string]pessoa{
		eu.sobrenome:  eu,
		ela.sobrenome: ela,
	}

	for _, v := range m {
		fmt.Println(v.nome, v.sobrenome)
		for _, sf := range v.saboresFavoritos {
			fmt.Printf("\t%s\n", sf)
		}
	}

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	fmt.Println()

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.color)
	fmt.Println(s.color)

	as := struct {
		nome      string
		amigos    map[string]int
		bebidaFav []string
	}{
		nome: "marcell",
		amigos: map[string]int{
			"A": 543,
			"B": 123,
			"C": 654,
		},
		bebidaFav: []string{
			"whisky",
			"Cerveja",
		},
	}
	fmt.Println()
	fmt.Println(as)
	fmt.Println(as.nome)
	fmt.Println(as.amigos)
	fmt.Println(as.bebidaFav)

	// ranger sobre MAP
	for key, val := range as.amigos {
		fmt.Println(key, val)
	}

	// RANGER sobre SLICE
	for index, val := range as.bebidaFav {
		fmt.Println(index, val)
	}
}
