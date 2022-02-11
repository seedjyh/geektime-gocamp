package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpu=", runtime.NumCPU())
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0), "->0")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(2), "->2")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(4), "->4")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(3), "->3")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(10), "->10")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(1), "->1")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0), "->0")
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0), "->0")
}
