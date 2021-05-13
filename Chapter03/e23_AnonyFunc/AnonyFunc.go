// 匿名函数是一种没有函数名的函数，即定义即使用
// 匿名函数没有函数名，只有函数体，它只有在被调用的时候才会被
// 初始化。匿名函数一般被当作一种类型赋值给函数类型的的变量，
// 经常被用作回调函数。
/*
func(params) (return params) {
	function body
}
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// 匿名函数声明后可直接调用
	// 在声明匿名函数后，在其后加上调用的参数列表，即可对匿名函数进行调用

	// 即在匿名函数体后使用实参列表
	func(name string) {
		fmt.Println("My name is ", name)
	}("王小二")

	// 可以将匿名函数赋值给函数类型的变量，用于多次调用或者求值
	// 使用匿名函数实现一个简单的报时器，并赋值给currentTime,
	// 每次调用currentTime都能知道当前系统的最新时间
	currentTime := func() {
		fmt.Println(time.Now())
	}
	// 调用匿名函数
	currentTime()
}

// My name is  王小二
// 2021-05-13 12:20:13.188804385 +0800 CST m=+0.000085640
