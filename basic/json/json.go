package json

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"basic/tools"
)

type user struct {
	First string
	Age   int
}

type person2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func example1() {

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "Marcell",
		Age:   41,
	}

	users := []user{u1, u2, u3}

	userMarchal, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
	fmt.Println(string(userMarchal))
}

func example2() {
	var people2 []person2

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println("-----")
	fmt.Println(s)

	err := json.Unmarshal([]byte(s), &people2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----")
	for _, person := range people2 {
		fmt.Println(person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t", saying)
		}
	}
	fmt.Println("-----")
}

func example3() {
	u1 := person2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person2{u1, u2, u3}
	fmt.Println("users:\n", users)
	fmt.Println("-----")

	fmt.Println("users Encoded:\n", users)
	if err := json.NewEncoder(os.Stdout).Encode(users); err != nil {
		fmt.Println("error")
	}
	fmt.Println("-----")
}

func example4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("-----")
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-----")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

// ByAge : Interface de person2 para sort.Sort()
type ByAge []person2

// ByLastName : Interface de person2 para sort.Sort()
type ByLastName []person2

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p ByLastName) Len() int           { return len(p) }
func (p ByLastName) Less(i, j int) bool { return p[i].Last < p[j].Last }
func (p ByLastName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func example5() {
	u1 := person2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person2{
		First: "M",
		Last:  "Hmmmm",
		Age:   42,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person2{u3, u2, u1}
	fmt.Println("-----")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, saying := range u.Sayings {
			fmt.Println("\t", saying)
		}
	}

	fmt.Println("-----")
	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, saying := range u.Sayings {
			fmt.Println("\t", saying)
		}
	}

	fmt.Println("-----")
	sort.Sort(ByLastName(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, saying := range u.Sayings {
			fmt.Println("\t", saying)
		}
	}
	fmt.Println("-----")
}

// Examples of json
func Examples() {
	tools.HeaderOutPut("jsonExample")

	example1()
	example2()
	example3()
	example4()
	example5()
}
