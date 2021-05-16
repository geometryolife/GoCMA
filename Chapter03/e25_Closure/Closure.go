// 闭包是携带状态的函数，可简单将其理解为"对象"，因为它同时具备状态和行为。
// 闭包是将函数内部和函数外部连接起来的桥梁，通过闭包可以读取函数内部的变量。
// 可以使用闭包封装变量的私有状态，让它们常驻于内存当中。
// 闭包能够引用其作用域上部的变量并进行修改，被捕获到闭包中的变量将随着闭包
// 的生命周期一直存在，函数本身是不存储信息的，但是闭包中的变量使闭包本身具
// 备了存储信息的能力。

// 用闭包的特性实现一个简单的计数器
package main

import "fmt"

// 函数createCounter()返回了一个闭包，该闭包中封装了计数值initial，
// 从外部代码无法直接访问该变量，不同的闭包之间变量不会干扰，c1和c2
// 两个计数器均独立进行计数。
func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	// 引用initial，创建一个闭包
	return func() int {
		initial++
		// 返回当前计数
		return initial
	}
}

func main() {
	// 计数器1
	c1 := createCounter(1)

	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	// 计数器2
	c2 := createCounter(10)

	fmt.Println(c2()) // 11
	fmt.Println(c1()) // 4
}

// 2
// 3
// 11
// 4
