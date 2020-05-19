package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	//v := Vertex1{3, 4}

	a = &f
	//a = &v
	fmt.Println(a.Abs())
}
//import "fmt"
//
//func main() {
//	var i interface{}
//	describe(i)
//
//	i = 42
//	describe(i)
//
//	i = "hello"
//	describe(i)
//}
//
//func describe(i interface{}) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}



//package main
//
//import (
//"fmt"
//"math"
//)
//
//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//
//func (t *T) M() {
//	fmt.Println(t.S)
//}
//
//type F float64
//
//func (f F) M() {
//	fmt.Println(f)
//}
//
//func main() {
//	var i I
//
//	i = &T{"Hello"}
//	describe(i)
//	i.M()
//
//	i = F(math.Pi)
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
