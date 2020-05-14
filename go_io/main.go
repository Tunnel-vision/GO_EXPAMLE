package main

import (
	"io"
	"os"
	"strings"
)

func fun1() {
	r := strings.NewReader("some io.reader stream to be read \n")
	lr := io.LimitReader(r, 4)
	io.Copy(os.Stdout, lr)
}

func main() {
	//io.Reader()
	//io.MultiReader()
	//io.LimitReader()
	//io.TeeReader()
	//fun1()
	//func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
	//func ReadFull(r Reader, buf []byte) (n int, err error)

}
