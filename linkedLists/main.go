package main

import (
	"fmt"
	"linkedLists/singly"
)

func main() {
	sl := singly.New()

	sl.InsertEnd(1)
	sl.InsertEnd(2)
	sl.InsertEnd(3)
	sl.InsertEnd(4)

	fmt.Print(sl.String())
	fmt.Println("Len:", sl.Len())

	fmt.Println("-----------------")
	sl = singly.New()

	sl.InsertFront(1)
	sl.InsertFront(2)
	sl.InsertFront(3)
	sl.InsertFront(4)

	fmt.Print(sl.String())
	fmt.Println("Len:", sl.Len())
}
