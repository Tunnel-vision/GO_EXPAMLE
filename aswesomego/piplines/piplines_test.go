package piplines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName1(t *testing.T) {
	//c := gen(2, 3, 4, 5, 6)
	//out := sq(c)
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//fmt.Println(<-out)

	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
			fmt.Println("----->>>", n)
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		fmt.Println("------------")
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func TestName2(t *testing.T) {
	num := make([]int, 0)
	for i := 0; i < 100; i++ {
		num = append(num, i)
	}
	c := gen(num...)
	c1 := sq(c)
	c2 := sq(c)
	go func() {
		for n := range c2 {
			fmt.Println("0000000>", n)
		}
	}()

	go func() {
		for n := range c1 {
			//time.Sleep(100 * time.Millisecond)
			fmt.Println("------->", n)
		}
	}()
	time.Sleep(3 * time.Second)
	//for n := range merge(c1, c2) {
	//	fmt.Println(n)
	//}
}

func TestName3(t *testing.T) {
	out := make(chan int, 100)
	go func() {
		for i := 1; i < 100; i++ {
			out <- i
		}
		close(out)
	}()
	for i := range out {
		fmt.Println(i)
	}
}

func TestName4(t *testing.T) {
	c := gen(2, 3)
	c1 := sq(c)
	c2 := sq(c)
	out := merge(c1, c2)
	fmt.Println(<-out)
	return

}
