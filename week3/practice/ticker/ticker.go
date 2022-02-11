package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	tickerDone := make(chan struct{})
	time.AfterFunc(time.Second * 5, func(){
		ticker.Stop()
		close(tickerDone)
	})
	FOR:
	for {
		select {
		case x := <- ticker.C:
			fmt.Println("tick:", x)
		case <- tickerDone:
			fmt.Println("ticker stopped")
			break FOR
		}
	}
	fmt.Println("done")
}
