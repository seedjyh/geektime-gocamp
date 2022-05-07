package main

import (
	"fmt"
	"grpc-test/app/client"
)

func main() {
	port := 11221
	serverAddr := fmt.Sprintf("127.0.0.1:%d", port)
	fmt.Println("This is client! Connect to address:", serverAddr)
	defer fmt.Println("Client exit!")
	if err := client.Run(serverAddr); err != nil {
		fmt.Println("Client run failed! err:", err)
	}
}
