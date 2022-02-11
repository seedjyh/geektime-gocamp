// defer f(g()) 需要在遇到defer行就预先执行g()。但f是在外部函数退出前才执行。
package main

import (
	"fmt"
	"time"
)

func query() string {
	fmt.Println("enter query")
	time.Sleep(time.Second * 5)
	fmt.Println("leave query")
	return "query result"
}

func main() {
	fmt.Println("enter main")
	fmt.Println("main.defer before")
	defer fmt.Println(query())
	fmt.Println("main.defer after")
	fmt.Println("leave main")
}
