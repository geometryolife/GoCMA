// 1.接口是调用方和实现方约定的一种合作协议
// 2.接口中定义了一系列将要被实现的函数，调用方通过接口了解
// 可使用的方法而无需了解具体实现，实现方通过接口对外提供
// 能使用的特性。
// 3.每个接口由一个或者多个方法组成
// 4.定义接口需要配合使用type和interface关键字
// 5.接口名的首字母和方法名的首字母皆大写时，表示该方法是公开的
// 6.Golang的接口具有嵌套特性，可以实现类似面向对象中的接口继承
// 特性，从而创造出新的接口。
// 7.Golang中所有类型都可以实现接口，函数作为Golang中的一种类型，也不例外
// 8.通过函数类型实现接口，可以将函数作为接口来使用，在运行时可以通过替换
// 具体的实现函数，实现类似多态的效果。

// 定义坦克的接口，具有行走、开跑功能
/*
type Tank interface {
	Walk() // 行走
	Fire() // 开炮
}
*/

// 定义飞机的接口，具有飞行功能
/*
type Plane interface {
	Fly() // 飞行
}
*/

// 定义新型武器，具有坦克和飞机的功能
/*
type PlaneTank interface {
	Tank
	Plane
}
*/

package main

import "fmt"

// 定义一个简单的接口Printer
// 功能是将输入的参数直接打印到命令行中
type Printer interface {
	// 打印方法
	Print(interface{})
}

// 由于函数不能直接实现接口，需要将函数定义为类型后，再使用
// 定义好的函数类型实现接口。

// 函数定义为类型
type FuncCaller func(p interface{})

// 实现Printer的Print方法
// 接口的实现由直接调用定义好的函数类型来完成
func (funcCaller FuncCaller) Print(p interface{}) {
	// 调用funcCaller函数本体
	funcCaller(p)
}

func main() {
	var printer Printer
	// 将匿名函数强转为FuncCaller赋值给printer
	printer = FuncCaller(func(p interface{}) {
		fmt.Println(p)
	})
	printer.Print("Golang is Good!")
}

// Golang is Good!
