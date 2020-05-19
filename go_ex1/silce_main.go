package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main_silce1() {
	primes := [6]int{2, 3, 4, 5, 6}
	var s = primes[1:6]
	fmt.Println(s)
	fmt.Println(primes)
	fmt.Println("primes type:", reflect.TypeOf(primes))
	fmt.Println("s type:", reflect.TypeOf(s))
}

func main_silce2() {
	names := [4]string{"john", "paul", "George", "Ringo"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)
	names[3] = "hh"
	fmt.Println(names)
	cc := []int{}
	fmt.Println(reflect.TypeOf(cc))
}

func main_silce3() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main_silce4() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

func main_silce5() {
	b := make([]int, 0, 5)
	fmt.Println(b, len(b), cap(b))
	b = b[:cap(b)]
	printSlice(b)
	fmt.Printf("b type: %T", b)
	b = b[1:]
	fmt.Println()
	printSlice(b)
}

func main_silce6() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// 两个玩家轮流打赏 X和O
	board[0][0] = "x"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main_silce7() {
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2, 3, 4, 5)
	printSlice(s)
}

func Pic(dx, dy int) [][]uint8 {

}

func main_silce8() {
	pic.Show(Pic)
}

func main() {
	//main_silce1()
	//main_silce3()
	//main_silce4()
	//main_silce5()
	//main_silce6()
	main_silce7()
}

// 扩展阅读 https://blog.go-zh.org/go-slices-usage-and-internals
