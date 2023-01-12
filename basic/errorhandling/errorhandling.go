package errorhandling

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	"basic/tools"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func exercise1() {
	tools.HeaderOutPut("exercise 1")

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))

	fmt.Printf("------\n")
}

func exercise2() {
	tools.HeaderOutPut("exercise 2")

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))

	fmt.Printf("------\n")
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Json marshal error: %s", err)
	}

	return bs, nil
}

type customErr struct {
	s string
}

func (ce customErr) Error() string {
	return ce.s
}

func exercise3() {
	tools.HeaderOutPut("exercise 3")

	ce := customErr{
		s: "Marcell custom error",
	}

	foo(ce)

	fmt.Printf("------\n")
}

func foo(e error) {
	fmt.Printf("%T\n", e)
	fmt.Println(e)
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func exercise4() {
	tools.HeaderOutPut("exercise 4")

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("------\n")
}

// Sqrt square a float and use a sqrtError
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			lat:  `3°04'18.9"S`,
			long: `60°01'04.5"W`,
			err:  fmt.Errorf("Negative lat erro"),
		}
	}
	return math.Sqrt(f), nil
}

// Examples group error handling examples
func Examples() {
	tools.HeaderOutPut("concurrencyExample")

	exercise1()
	exercise2()
	exercise3()
	exercise4()
}
