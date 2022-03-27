package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)
import "golang.org/x/sync/errgroup"

type task struct {
	id        int
	cost      time.Duration
	returnErr error
}

func (t *task) Do() error {
	fmt.Println("start task id ", t.id)
	time.Sleep(t.cost)
	return t.returnErr
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	tasks := []*task{
		{1, time.Second * 2, nil},
		{2, time.Second * 4, errors.New("err2")},
		{3, time.Second * 6, errors.New("err3")},
	}
	for _, t := range tasks {
		g.Go(t.Do)
	}
	fmt.Println("start <-Done at:", time.Now())
	fmt.Println("wait ctx:", <-ctx.Done())
	fmt.Println("finish <-Done at:", time.Now())
	fmt.Println("start Wait at:", time.Now())
	fmt.Println("wait return:", g.Wait())
	fmt.Println("finish wait at", time.Now())
}
