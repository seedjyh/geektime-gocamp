package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ch1 := make(chan os.Signal, 10)
	ch2 := make(chan os.Signal, 10)
	signal.Notify(ch1)
	signal.Notify(ch2)
	signal.Notify(ch1)
	signal.Notify(ch2)
	timer3 := time.NewTimer(time.Second * 3)
	timer10 := time.NewTimer(time.Second * 10)
	for {
		select {
		case x := <- ch1:
			fmt.Println("ch1 received", x)
		case x := <- ch2:
			fmt.Println("ch2 received", x)
		case <- timer3.C:
			fmt.Println("stop interrupt")
			signal.Reset(os.Interrupt)
		case <- timer10.C:
			return
		}
	}
}
