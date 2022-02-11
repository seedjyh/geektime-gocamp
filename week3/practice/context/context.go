package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1 * time.Second))

	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println(ctx.Err())
	// nil 表示 context还没结束；
	// "context deadline exceeded" 表示先发生deadline超时
	// "context canceled" 表示先发生cancel()调用
}
