package singly

import "fmt"

type Node struct {
	data int32
	next *Node
}

func (e *Node) Next() *Node {
	if e.next != nil {
		return e.next
	}

	return nil
}

type List struct {
	head *Node
	len  int
}

func New() *List {
	return new(List)
}

func (l *List) Print() {
	if l.head == nil {
		fmt.Println("No nodes in list")
	}

	ptr := l.head

	for ptr != nil {
		fmt.Println("Node: ", ptr.data)
		ptr = ptr.Next()
	}
}

func (l *List) Len() {
	fmt.Println(l.len)
}

func (l *List) Insert(value int32, location string) {
	n := new(Node)
	n.data = value

	if l.head == nil {
		l.head = n
		l.len++

		return
	}

	switch {
	case location == "end":
		ptr := l.head

		for ptr.Next() != nil {
			ptr = ptr.Next()
		}

		ptr.next = n
		l.len++

	case location == "front":
		n.next = l.head
		l.head = n
	}
}

func (l *List) InsertEnd(value int32) {
	l.Insert(value, "end")
}

func (l *List) InsertFront(value int32) {
	l.Insert(value, "front")
}
