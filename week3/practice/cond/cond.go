package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	mu sync.Mutex
	items []string
	cond sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond.L = &q.mu
	return q
}

func (q *Queue) Push(v string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, v)
	q.cond.Signal() // 允许调用者持有锁，但不强制
}

func (q *Queue) Pop() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	q := NewQueue()
	go func() {
		time.Sleep(time.Second)
		q.Push("hello")
		q.Push("world")
	}()
	fmt.Println("start")
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
