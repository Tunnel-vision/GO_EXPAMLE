package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	//io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
	r := strings.NewReader("hello world")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
