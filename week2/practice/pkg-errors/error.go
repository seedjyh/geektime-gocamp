package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	e := errors.Wrap(nil, "abc")
	fmt.Println(e)
}
