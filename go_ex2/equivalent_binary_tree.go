package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		fmt.Println(v)
		t = insert(t, v)
	}
	//fmt.Printf("value is %v", t)
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}
func (t *Tree) String() string {
	if t == nil {
		return "nil"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
//((((0 (1)) 2 (3)) 4 ((5) 6 ((7) 8))) 9)

func main() {
	tree := New(1)
	fmt.Println(tree)
}
