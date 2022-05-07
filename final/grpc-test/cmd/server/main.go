package main

import (
	"fmt"
	"grpc-test/app/server"
)

func main() {
	port := 11221
	fmt.Println("This is server! Listen at port:", port)
	defer fmt.Println("Server exit!")
	if err := server.Run(port); err != nil {
		fmt.Println("Server exit! err:", err)
	}
}
