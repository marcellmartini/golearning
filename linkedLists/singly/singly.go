package singly

import (
	"fmt"
)

const (
	end   = "end"
	front = "front"
)

type Node struct {
	data int32
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type List struct {
	head *Node
	len  int
}

func New() *List {
	return new(List)
}

func (l *List) String() string {
	if l.head == nil {
		return fmt.Sprintln("No nodes in list")
	}

	ptr := l.head
	var output string

	for ptr != nil {
		output += fmt.Sprintln("Node:", ptr.data)
		ptr = ptr.Next()
	}

	return output
}

func (l *List) Len() int {
	return l.len
}

func (l *List) insert(value int32, location string) {
	n := new(Node)
	n.data = value

	switch {
	case l.head == nil:
		l.head = n
		l.len++
		return

	case location == end:
		ptr := l.head

		for ptr.Next() != nil {
			ptr = ptr.Next()
		}

		ptr.next = n

	case location == front:
		n.next = l.head
		l.head = n
	}

	l.len++
}

func (l *List) InsertEnd(value int32) {
	l.insert(value, end)
}

func (l *List) InsertFront(value int32) {
	l.insert(value, front)
}
