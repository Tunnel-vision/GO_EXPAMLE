package main

import "testing"

type MultipleFieldStructer struct {
	a int
	b string
	c float32
	d float64
	e int32
	f bool
	g uint64
	h *string
	i uint16
}

var x MultipleFieldStructer
////go:noinline
func emptyInterface(i interface{}) {
	s := i.(MultipleFieldStructer)
	x = s
}
////go:noinline
func typed(s MultipleFieldStructer) {
	x = s
}

func BenchmarkWithType(b *testing.B) {
	s := MultipleFieldStructer{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		typed(s)
	}
}

func BenchmarkWithemptyInterface(b *testing.B) {
	s := MultipleFieldStructer{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterface(s)
	}
}
