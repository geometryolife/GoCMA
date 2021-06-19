// 1. 函数是一段封装好、可重复使用、针对单一功能的代码片段。
// 2. 函数的作用:
// 模块化，提高代码的可重用性
// 3. Golang 中函数声明包括函数名、形参列表、返回值。
/*
func name(params) (return params) {
	function body
}
*/
// 4. 在同一个包内，函数名不可重名。
// 5. 一个函数如果希望被包外代码访问，函数名的首字母需要大写。
// 6. 参数列表中每个参数由参数变量名和参数类型组成，它们作为函数的
// 局部变量被使用。在参数列表中多个参数之间通过逗号分隔。如果相邻的
// 参数的类型是相同的，则可以省略类型。
/*
func cal(a, b int) int {
	return a + b;
}
*/
// 7. Golang 不仅支持多返回值，还支持对返回值进行命名。
// 8. 命名返回值的作用:
// 可以直接在函数体内对返回值进行赋值。
// 注意:
// 在使用命名返回值的函数中，在函数结束前我们需要
// 显式使用 return 语句进行返回。
// 命名返回值和非命名返回值不能混合使用，只能二选
// 一，否则会出现编译错误。
// Tips:
// Golang 中函数的传递方式都是值传递，在实际开发中
// 为了减少复制时产生的性能损耗，可以在参数中使用
// 指针或者引用来减少内存复制的操作。
package main

import "fmt"

func main() {
	q, r := div(10, 3)
	fmt.Printf("quotient: %v, remainder: %v\n", q, r)
}

func div(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

/*
>>> Execution Result:
quotient: 3, remainder: 1
*/
