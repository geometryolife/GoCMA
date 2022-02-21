// 协程 goroutine 在 Go 中属于轻量级线程，main 函数运行在 goroutine 上
// 大多数情况下，test 函数不会执行，主 goroutine 在调度 go test() 执行
// 前就可能结束了，只要主 go goroutine (main 函数) 结束，就意味着整个程
// 序已经运行结束了。
package main

import (
	"fmt"
)

func test() {
	fmt.Println("I am work in a single goroutine.")
}

func main() {
	go test()
}
