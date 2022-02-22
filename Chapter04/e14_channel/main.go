// 【例4-3】使用带缓冲区的 channel
package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	// 线程休息 100s 再从 channel 读取数据
	time.Sleep(time.Second * 100)
	<-ch
}

func main() {
	// 创建一个长度为 2 的 channel
	ch := make(chan int, 2)
	go consume(ch)

	ch <- 0
	ch <- 1
	// ch <- 2
	// === Output ===
	// I am free!
	// fatal error: all goroutines are asleep - deadlock!
	// 通道被占满了还继续往通道放就会导致死锁！
	// 疑问？为什么第四个才报死锁，而不是第三个就报告？

	// 发送数据不被阻塞
	fmt.Println("I am free!")
	ch <- 2
	fmt.Println("I can not go there within 100s!")

	time.Sleep(time.Second)
}

/*
>>> Execution Result:
I am free!
I can not go there within 100s!
*/
