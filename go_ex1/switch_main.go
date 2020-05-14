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

func main() {
	main22()
}
