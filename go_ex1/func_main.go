package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func func_main1() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func func_main2() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// 返回一个“返回int的函数”
func fibonacci() func() int {
	sum := 1
	return func() int {
		sum = sum+sum
		return sum
	}
}

func func_main3() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func main() {
	//func_main1()
	//func_main2()
	func_main3()
}
