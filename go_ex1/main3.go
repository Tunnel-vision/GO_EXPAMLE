package main

import (
"fmt"
"math"
)

func main() {
	var x, y int = 3, 5
	var f float64  = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	z  := uint(f)
	fmt.Println(x, y, z)
}