package main

import (
	"fmt"
	"time"
)

func main() {
	stop := false
	time.AfterFunc(time.Second, func() { stop = true })
	for !stop {
		time.Sleep(time.Millisecond)
	}
	fmt.Println("done")
}
