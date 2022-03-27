package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "你好"
	fmt.Println("len(s) =>", len(s))
	fmt.Println("utf8.RuneCountInString(s) => ", utf8.RuneCountInString(s))
	a := make([]int, 0, 8)
	fmt.Println("len(a) => ", len(a))
	fmt.Println("cap(a) => ", cap(a))
}
