package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func base_main() {
	s := []int{7, 1, 3, 1, -10}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func base2_main() {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func base3_main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i ++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func main() {
	//base_main()
	//base2_main()
	base3_main()

}
