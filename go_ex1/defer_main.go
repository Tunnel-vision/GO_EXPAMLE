package main

import "fmt"
var a int

func main111()  {
	a = 1
	// defer 语句会将函数推迟到外层函数返回之后执行。
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	// 推迟的函数调用会被压入一个栈中,当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	defer fmt.Println("world11",a)
	a +=1
	defer fmt.Println("world22",a)
	fmt.Println("hello")
	a +=1
}

func main222()  {
	fmt.Println("counting")
	for i :=0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main()  {
	//main111()
	main222()
}