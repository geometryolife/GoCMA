package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("I am work in a single goroutine.")
}

func main() {
	go test()
	// 主 goroutine 休眠 1 秒
	time.Sleep(time.Second)
}

/*
>>> Execution Result:
I am work in a single goroutine.
*/
