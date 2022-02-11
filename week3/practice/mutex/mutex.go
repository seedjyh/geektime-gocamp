package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	mu := sync.Mutex{}
	g1 := 0
	g2 := 0
	// g1
	go func () {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				g1++
				time.Sleep(time.Microsecond * 100)
				mu.Unlock()
			}
		}
	}()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 100)
		mu.Lock()
		g2++
		mu.Unlock()
	}
	fmt.Println("g1:", g1)
	fmt.Println("g2:", g2)
	close(done)
}
