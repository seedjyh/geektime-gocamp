package main

import (
	"fmt"
	"geektime-gocamp/week9/homework/tcp-demo/internal/app/server"
	"geektime-gocamp/week9/homework/tcp-demo/internal/pkg/app"
)

func main() {
	a := app.New(
		"server",
		"1.0",
		server.NewServer(":8123"),
	)
	if err := a.Run(); err != nil {
		fmt.Println("Exit with error:", err)
	} else {
		fmt.Println("Exit OK.")
	}
}
