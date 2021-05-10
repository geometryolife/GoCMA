// Golang中提供类型别名的语法特性
// 类型别名本质上与原类型同属于一个类型，相当于原类型的别称
// type name = T
// 类型定义将会创建一种新的类型，新的类型具备原类型的特性
// type name T
package main

import "fmt"

type aliasInt = int // 定义一个类型别名
type myInt int      // 定义一个新的类型

func main() {
	var alias aliasInt
	fmt.Printf("alias value is %v, type is %T\n", alias, alias)

	var myint myInt
	fmt.Printf("myint value is %v, type is %T\n", myint, myint)
}

// alias value is 0, type is int
// myint value is 0, type is main.myInt
