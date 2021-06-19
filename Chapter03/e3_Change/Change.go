package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	var tmp int
	c := 3
	d := 4

	fmt.Printf("Original a = %v\n", a)
	fmt.Printf("Original b = %v\n", b)

	// 过去的编程语言常常借助第三方临时变量来实现变量交换
	tmp = a
	a = b
	b = tmp

	fmt.Printf("New a = %v\n", a)
	fmt.Printf("New b = %v\n", b)

	fmt.Printf("Original c = %v\n", c)
	fmt.Printf("Original d = %v\n", d)

	// Golang 提供多重赋值特性
	// 多重赋值过程中，变量的左值和右值按照从左往右的顺序赋值
	c, d = d, c

	fmt.Printf("New c = %v\n", c)
	fmt.Printf("New d = %v\n", d)
}

/*
>>> Execution Result:
Original a = 1
Original b = 2
New a = 2
New b = 1
Original c = 3
Original d = 4
New c = 4
New d = 3
*/
