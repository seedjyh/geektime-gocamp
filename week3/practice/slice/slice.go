package main

import "fmt"

func main() {
	raw := []int{1,2,3,4,5}
	s := raw[1:3:5]
	for _, x := range s {
		fmt.Printf("%d ", x) // 2 3
	}
	fmt.Println("")
	for _, x := range s[:len(s) + 1] {
		fmt.Printf("%d ", x) // 2 3 4
	}
	fmt.Println("")
	for _, x := range s[:len(s) + 2] {
		fmt.Printf("%d ", x) // 2 3 4 5
	}
	fmt.Println("")
}
