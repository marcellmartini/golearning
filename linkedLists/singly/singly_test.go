package singly

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		xInsert  []int32
		location string
		want     List
	}{
		{
			name:     "Inset at the end of the list",
			xInsert:  []int32{1, 2},
			location: end,
			want: List{
				&Node{
					1,
					&Node{
						2,
						nil,
					},
				},
				2,
			},
		},
		{
			name:     "Inset in the front of the list",
			xInsert:  []int32{1, 2},
			location: front,
			want: List{
				&Node{
					2,
					&Node{
						1,
						nil,
					},
				},
				2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			for _, i := range tt.xInsert {
				l.insert(i, tt.location)
			}

			if !reflect.DeepEqual(*l, tt.want) {
				t.Errorf("Insert(%v, %v) got ", tt.xInsert, tt.location)
				fmt.Print(l.String())
				fmt.Println("Len:", l.Len())
				fmt.Println("want: ")
				fmt.Print(tt.want.String())
				fmt.Println("Len:", tt.want.Len())
			}
		})
	}
}

func TestInsertEnd(t *testing.T) {
	test := struct {
		xInsert  []int32
		location string
		want     List
	}{
		xInsert:  []int32{1, 2, 3},
		location: end,
		want: List{
			&Node{
				1,
				&Node{
					2,
					&Node{
						3,
						nil,
					},
				},
			},
			3,
		},
	}

	l := New()
	for _, i := range test.xInsert {
		l.InsertEnd(i)
	}

	if !reflect.DeepEqual(*l, test.want) {
		t.Errorf("InsertEnd(%v) got ", test.xInsert)
		fmt.Println(l.String())
		fmt.Println("want: ")
		fmt.Println(test.want.String())
	}
}

func TestInsertFront(t *testing.T) {
	test := struct {
		xInsert  []int32
		location string
		want     List
	}{
		xInsert:  []int32{1, 2, 3},
		location: end,
		want: List{
			&Node{
				3,
				&Node{
					2,
					&Node{
						1,
						nil,
					},
				},
			},
			3,
		},
	}

	l := New()
	for _, i := range test.xInsert {
		l.InsertFront(i)
	}

	if !reflect.DeepEqual(*l, test.want) {
		t.Errorf("InsertFront(%v) = ", test.xInsert)
		fmt.Println(l.String())
		fmt.Println("want: ")
		fmt.Println(test.want.String())
	}
}

func TestFLen(t *testing.T) {
	test := struct {
		list List
		want int
	}{
		list: List{
			&Node{
				1,
				nil,
			},
			1,
		},
		want: 1,
	}

	if got := test.list.Len(); got != test.want {
		t.Errorf("list.Len() got %v want: %v", got, test.want)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		list List
		want string
	}{
		{
			name: "Not empty list",
			list: List{
				&Node{
					1,
					nil,
				},
				1,
			},
			want: "Node: 1\n",
		},
		{
			name: "Empty list",
			list: List{},
			want: "No nodes in list\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.String(); got != tt.want {
				t.Errorf("list.String() got %v want: %v", got, tt.want)
			}
		})
	}
}
