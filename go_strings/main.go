package main

import (
	"fmt"
	"strings"
	"unicode"
)

func SplitFunctionSentence() {
	str := "hello&$ world"
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("%q", strings.FieldsFunc(str, f))

}

// 字符串的搜索与匹配
func main() {
	fmt.Println("____________")
	//strings.Contains可以检测字符串是否包含某个子串
	//contains := strings.Contains("abcd", "ab")
	//fmt.Println(contains)
	//可以检测字符串是否包含某个字符；
	//containsRune := strings.ContainsRune("abcd", 'c')
	//fmt.Println(containsRune)
	//any := strings.ContainsAny("abcd?", "?")
	//fmt.Println(any)
	//index := strings.Index("abdcd", "cd")
	//fmt.Println(index)
	//strings.IndexByte()
	//strings.IndexRune()
	//strings.IndexAny() 字符集合搜索，匹配chars中的任何一个字符
	//strings.LastIndex()
	//strings.LastIndexByte()
	//strings.LastIndexAny()
	//split := strings.Split("a,b,c", ",")
	//fmt.Println(split)
	//splitAfter := strings.SplitAfter("a,b,c", ",")
	//fmt.Println(splitAfter)
	//去除多个非字幕的字符
	//SplitFunctionSentence()
	//strings.Trim() 删除字符串首尾的多余字符
	//strings.TrimLeft()
	//strings.TrimRight()
	//strings.TrimSpace() 删除字符串首尾的空格
	//strings.Join()
	//strings.Compare()
	//strings.Count()
	//strings.Replace()
}
