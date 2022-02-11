package main

import (
	"fmt"
	"geektime-gocamp/week3/homework/app"
	"geektime-gocamp/week3/homework/service"
)

func main() {
	a := app.New(
		app.Name("week-3-homework"),
		app.Version("1.0.0"),
		service.New(":8083", "Hello"),
		service.New(":8084", "World"),
	)
	if err := a.Run(); err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("finished")
}
