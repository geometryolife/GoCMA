// 匿名函数比较常用的场景是用作回调函数
package main

import "fmt"

// 使用回调函数处理字符串
// 接受string和匿名函数的参数输入，然后使用匿名函数对string进行处理
func proc(input string, processor func(str string)) {
	// 调用匿名函数
	processor(input)
}
func main() {
	proc("王小二", func(str string) {
		for _, v := range str {
			fmt.Printf("%c\n", v)
		}
	})
}

// 上面代码中的匿名函数被作为回调函数用于对传递的字符串进行
// 处理，用户可以根据自己的需要传递不同的匿名函数实现对字符
// 串进行不同的处理操作。

// 王
// 小
// 二
