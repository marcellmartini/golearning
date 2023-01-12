package function

import (
	"fmt"
	"math"

	"basic/tools"
)

type person struct {
	nome      string
	sobrenome string
	idade     int
}

func (p person) falar() {
	fmt.Printf("Meu nome é %v %v, e tenho %v anos de idade.\n", p.nome, p.sobrenome, p.idade)
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (d square) area() float64 {
	return d.length * d.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	a := s.area()
	fmt.Println(a)
}

// Examples of func
func Examples() {
	fmt.Println(foo())
	fmt.Println(bar())

	xi := []int{1, 2, 3, 4, 5}
	i := foo2(xi...)
	fmt.Println(i)

	xi2 := []int{1, 2, 3, 4, 5}
	i2 := bar2(xi2)
	fmt.Println(i2)

	marcell := person{
		nome:      "Marcell",
		sobrenome: "Martini",
		idade:     41,
	}
	marcell.falar()

	c := circle{
		radius: 12.345,
	}

	s := square{
		length: 15,
	}

	info(c)
	info(s)

	func() {
		fmt.Println("Função anônima")
	}()

	f := func() {
		fmt.Println("função em uma variável")
	}

	f()

	count := maisUM()
	count(3)

	x := somaSliceBaseadoTamanho(separaSlicePARImparBaseadoTamanho, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)

	x = somaSliceBaseadoTamanho(separaSlicePARImparBaseadoTamanho, []int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(x)

	printString("Marcell")
	defer printString("Martini")

	fmt.Println("---- Ultima linha antes do '}' ----")
}

func separaSlicePARImparBaseadoTamanho(xi []int) []int {
	xp := []int{}
	xxi := []int{}

	for i, v := range xi {
		if i%2 == 0 {
			xp = append(xp, v)
		} else {
			xxi = append(xxi, v)
		}
	}
	if len(xi)%2 == 0 {
		return xp
	}
	return xxi
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v > lim {
		fmt.Printf("%g >= %g\n", v, lim)
	} else {
		return v
	}
	// can't use v here, though
	return lim
}

func somaSliceBaseadoTamanho(g func(xi []int) []int, xi []int) int {
	xii := g(xi)
	sum := 0
	for _, v := range xii {
		sum += v
	}
	return sum
}

func printString(s string) {
	fmt.Println(s)
}

func maisUM() func(c int) {
	return func(c int) {
		c++
		fmt.Println(c)
	}
}

func bar2(i []int) int {
	tools.HeaderOutPut("bar2 func")
	sum := 0
	for _, v := range i {
		sum += v
	}

	return sum
}

func foo2(i ...int) int {
	tools.HeaderOutPut("foo2 func")
	sum := 0
	for _, v := range i {
		sum += v
	}

	return sum
}

func foo() int {
	tools.HeaderOutPut("foo func")
	return 2 + 2
}

func bar() (int, string) {
	tools.HeaderOutPut("bar func")
	return 2 + 2, "Marcell Martini"
}
