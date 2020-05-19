package main

import "fmt"

func mainx1() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func mainx2()  {
	var b []int
	fmt.Println(b)
}

func main() {
	//mainx1()
	mainx2()
}
