package main

import "fmt"

func descerible(i interface{}) {
	fmt.Printf("value=%v type=%T\n", i, i)
}

func assert_main1() {
	var i interface{} = "hello"
	s := i.(string)
	descerible(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
		//fmt.Println("this is int")
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("i dont know type %T\n", i)
	}

}

func assert_main2() {
	//do(4)
	//do("nihao中国")
	ints := make([]int, 1, 1)
	do(ints)
}

func main() {
	//assert_main1()
	assert_main2()
}
