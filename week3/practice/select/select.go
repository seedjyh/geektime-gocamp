// select 的 case 里包含函数调用时，会先逐个（！）执行这些函数调用，全部完成后，才执行select。
// 可以把 select 看做一个特殊的「函数」，其执行之前，所有参数都要预先计算好。
// 所谓 select 的参数，就是select 的 case里出现的任何「函数调用」。
package main

import (
	"fmt"
	"time"
)

func query(q string) string {
	fmt.Println("start query", q)
	defer fmt.Println("finish query", q)
	time.Sleep(time.Second * 5)
	return "resp-of-" + q
}

func main() {
	ch := make(chan string)
	res := make(chan string, 1)
	go func(){
		select {
		case ch <- query("q1"):
			res <- "to-ch-1"
		case ch <- query("q2"):
			res <- "to-ch-2"
		default:
			res <- "default"
		}
	}()
	fmt.Println("main start wait")
	for {
		select {
		case x := <- ch:
			fmt.Println("main ch:", x)
		case x := <- res:
			fmt.Println("main res:", x)
		}
	}
}

