package main

import "fmt"

var pow2 = []int{1, 2, 3, 4, 5, 6, 7, 8}

func main_range1() {
	for i, v := range pow2 {
		fmt.Printf("the index %d ,the value id %+v\n", i, v)
	}
}

func main() {
	main_range1()
}
