package slice

import (
	"fmt"

	"basic/tools"
)

// Examples of slices
func Examples() {
	tools.HeaderOutPut("sliceExample")

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	z := append(x[0:2], x[4:]...)
	fmt.Println(z)

	ex := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}

	fmt.Println(ex[:5])
	fmt.Println(ex[5:])
	fmt.Println(ex[2:7])
	fmt.Println(ex[1:6])

	m := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	m = append(m[:3], m[6:]...)
	fmt.Println(m)

	soa := make([]string, 50, 50)
	fmt.Println(soa)
	fmt.Println(len(soa))
	fmt.Println(cap(soa))

	soa = append(soa, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	for i := 0; i < len(soa); i++ {
		fmt.Println(i, soa[i])
	}

	h71 := []string{"James", "Bond", "Shaken, not stirred"}
	h72 := []string{"Miss", "Moneypenny", "Hello, James."}
	fmt.Println(h71)
	fmt.Println(h72)

	h7 := [][]string{h71, h72}
	for i, xs := range h7 {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, val)

		}
	}
	fmt.Println(h7)

}
