package main

import (
	"fmt"
	"math"
	"time"
)

func main1() {
	for i := 0; i < 100; i++ {
		fmt.Printf("the %d\n", i)
	}
}

func main2() {
	sum := 1
	for sum < 100 {
		sum += 1
		fmt.Println(sum)
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main3() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func main4(z, x float64) {
	//var z, x float64
	for {
		z -= (z*z - x/(2*z))
		time.Sleep(time.Duration(1)*time.Second)
		fmt.Println(z)
	}
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	//main1()
	//main2()
	//main3()
	//fmt.Println(
	//	pow1(3, 2, 10),
	//)
	main4(1.0,3.0)
}
