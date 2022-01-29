package main

import "fmt"

type Book struct {
	Title string
	Price float32
}
type Shelf struct {
	Pos          int
	Book1, Book2 Book
}

func main() {
	var s1, s2 Shelf
	s1.Book1 = Book{"a", 1.0}
	s1.Book2 = Book{"b", 2.0}
	s1.Pos = 3
	s2.Book1 = Book{"a", 1.0}
	s2.Book2 = Book{"b", 2.0}
	s2.Pos = 3

	fmt.Println(s1 == s2)
}
