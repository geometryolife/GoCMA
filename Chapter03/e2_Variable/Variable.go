package main

import "fmt"

func main() {
	var a int = 100
	var b = "100"
	// 短变量声明不能用在全局变量上
	// 使用短变量声明，左值变量中至少有一个是未声明过的，否则报错
	// 为了提高精度，Golang对浮点数默认推到类型为float64
	c := 0.17
	// 声明变量时，Golang自动把变量对应的内存区初始化
	var d string

	// 每个声明的变量都必须被使用，否则编译不通过
	fmt.Printf("a value is %v, type is %T\n", a, a)
	fmt.Printf("b value is %v, type is %T\n", b, b)
	fmt.Printf("c value is %v, type is %T\n", c, c)
	fmt.Printf("d value is %v, type is %T\n", d, d)
}

// Golang 提供类型推导特性

// a value is 100, type is int
// b value is 100, type is string
// c value is 0.17, type is float64
// d value is , type is string
