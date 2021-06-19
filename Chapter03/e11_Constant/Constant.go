// 1. 每个声明的常量，不必被使用也能通过编译，区别于变量。
package main

func main() {
	// 声明常量时指定类型
	const str string = "Golang is Good!"
	// 使用类型推导特性，省略声明时的类型和同时声明多个常量
	const name = "Golang is Good!"
	const a, b, c = 1, 2, 3
	const (
		surname      = "王"
		personalName = "小二"
	)
}
