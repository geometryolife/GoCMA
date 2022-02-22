// 【例4-2】协程使用 channel 发送和接收数据
package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	// 使用 for 循环从 channel 中读取数据
	for val := range ch {
		// 读取到结束符号
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func main() {
	// 创建一个无缓冲的 channel
	ch := make(chan string)
	go printInput(ch)

	// 从命令行读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("End the game!")
			break
		}
	}

	// 程序最后关闭 ch
	defer close(ch)
}

/*
>>> Execution Result:
Hello
Input is Hello
Hi
Input is Hi
EOF
End the game!
*/
