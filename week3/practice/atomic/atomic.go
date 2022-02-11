// 试用一下 atomic.Value ，对比和 sync.Mutex 的性能区别
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	a []int
}

func dataRaceMethod() {
	cfg := Config{}
	go func() {
		for i := 0; ; i++ {
			cfg.a = []int{
				i, i + 1, i + 2, i + 3, i + 4, i + 5,
			}
		}
	}()
	time.Sleep(time.Millisecond)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%+v\n", cfg)
		}()
	}
	wg.Wait()
}

func mutexMethod() {
	cfg := new(Config)
	rwm := sync.RWMutex{}
	go func() {
		for i := 0; ; i++ {
			rwm.Lock()
			cfg.a = []int{
				i, i + 1, i + 2, i + 3, i + 4, i + 5,
			}
			rwm.Unlock()
		}
	}()
	time.Sleep(time.Millisecond)
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwm.RLock()
			defer rwm.RUnlock()
			fmt.Printf("%+v\n", cfg)
		}()
	}
	wg.Wait()
}

func atomicMethod() {
	v := atomic.Value{}
	go func() {
		for i := 0; ; i++ {
			cfg := new(Config)
			cfg.a = []int{
				i, i + 1, i + 2, i + 3, i + 4, i + 5,
			}
			v.Store(cfg)
		}
	}()
	time.Sleep(time.Millisecond)
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cfg := v.Load()
			fmt.Printf("%+v\n", cfg)
		}()
	}
	wg.Wait()
}

func main() {
	dataRaceMethod()
}
