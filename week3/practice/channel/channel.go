package main

import "fmt"

func main() {
	ch := make(chan error, 2)
	fmt.Println("cap(ch)=", cap(ch))
	ch <- nil
	fmt.Println("push 1st item...")
	fmt.Println("cap(ch)=", cap(ch))
	ch <- nil
	fmt.Println("push 2nd item...")
	fmt.Println("cap(ch)=", cap(ch))
	// BUG
	ch <- nil
	fmt.Println("push 3rd item...")
	fmt.Println("cap(ch)=", cap(ch))
}
