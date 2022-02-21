// 使用 go 关键字可以很简单地启动一个新的协程并发执行函数
// 任务，提高程序的并行能力。
package main

import (
	"fmt"
	"time"
)

func main() {
	// go 语句既可以启动命名函数，也可以启动匿名函数
	go func(name string) {
		fmt.Println("Hello " + name)
	}("Joe")

	// 主 goroutine 阻塞 1s
	time.Sleep(time.Second)
}

/*
>>> Execution Result:
Hello Joe
*/
