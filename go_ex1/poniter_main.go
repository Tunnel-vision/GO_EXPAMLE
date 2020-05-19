package main

import "fmt"

// 指针保存了值的内存地址
func main1111() {
	i, j := 42, 2701
	p := &i         // 指向指针 & 操作符会生成一个指向其操作数的指针。
	fmt.Println(*p) // 通过指针读取i的值
	*p = 21         //通过指针设置i的值

	fmt.Println(i)

	p = &j       // 指向j
	*p = *p / 37 // 通过指针 对j进行除法运算

	fmt.Println(j)
}

func main() {
	main1111()
}

// *T 是指向 T 类型的值的指针
// & 操作符会生成一个指向其操作的数的指针
// * 操作符表示指针指向的底层值
//fmt.Println(*p) // 通过指针 p 读取 i
//*p = 21         // 通过指针 p 设置 i
// 这也就是通常所说的“间接引用”或“重定向”。