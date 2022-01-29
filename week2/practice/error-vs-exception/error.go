package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var myError = errors.New("sentinel error")

func Super(a, b int) error {
	return errors.Wrap(Medium(a, b), "here is super")
	// return errors.WithMessage(Bottom(a, b), "here is super")
}

func Medium(a, b int) error {
	// return errors.Wrap(Bottom(a, b), "here is medium")
	return errors.WithMessage(Bottom(a, b), "here is medium")
}

func Bottom(a, b int) error {
	// return errors.Errorf("found error: a=%d, b=%d", a, b)
	return errors.Wrap(myError, fmt.Sprintf("a=%d, b=%d", a, b))
}

func main() {
	e := Super(3, 4)
	fmt.Println("e:", e)
	fmt.Println("cause:", errors.Cause(e))
	fmt.Println("==", errors.Cause(e) == myError)
	fmt.Println("Is", errors.Is(e, myError))
	fmt.Printf("stack%%s: %s\n", e)
	fmt.Printf("stack%%v: %v\n", e)
	fmt.Printf("stack%%+v: %+v\n", e)
	// fmt.Println("end")
}

