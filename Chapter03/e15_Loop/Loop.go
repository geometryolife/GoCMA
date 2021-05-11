// Golang的循环体中仅提供for关键字
// 初始语句、条件表达式、结束语句都可不写
// 三者缺省为无限循环
/*
for init; condition; end {
	循环体代码
}
*/
package main

import "fmt"

func main() {
	// 无限循环
	for {
		fmt.Println("Hello Go!")
	}
}

// break跳出循环体，continue继续下一个循环
