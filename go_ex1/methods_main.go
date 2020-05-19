package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 方法就是一类带特殊的 接收者 参数的函数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods_main1() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods_main2() {
	f := MyFloat1(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func main() {
	//methods_main1()
	//methods_main2()
}
