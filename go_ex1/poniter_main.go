package main

import "fmt"

func main1111() {
	i, j := 42, 2701
	p := &i // 指向指针
	fmt.Println(*p) // 通过指针读取i的值
	*p = 21 //通过指针设置i的值

	fmt.Println(i)

	p = &j // 指向j
	*p = *p/37 // 通过指针 对j进行除法运算

	fmt.Println(j)
}

func main() {
	main1111()
}
