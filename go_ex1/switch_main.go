package main

import (
	"fmt"
	"runtime"
	"time"
)

func main11() {
	fmt.Println("GO runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")

	case "linux":
		fmt.Println("Linux")
	default:
		//plan9 windows:
		fmt.Printf("%s \n", os)

	}
}

func main22() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning!!")
	case t.Hour() > 12:
		fmt.Println("good afternoon!")
	default:
		fmt.Println("good evening!")
	}
	fmt.Println(t.Hour())
}

func main33(a int) {
	switch {
	case a < 10:
		fmt.Println(a)
		fallthrough
	case a < 13:
		fmt.Println(a)
		fallthrough
	default:
		fmt.Println("null")
	}
}

func main() {
	//main22()
	main33(5)
}
