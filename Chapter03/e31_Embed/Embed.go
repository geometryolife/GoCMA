// 1. 使用结构体内嵌，可以很形象地模拟面向对象语言设计中组合的特性。
// 2. 结构体内嵌属于一种组合的特性，通过组合不同的基础结构体，可以
// 构建出具备不同基础特性的复杂结构体。
// 3. 通过 Golang 的内嵌结构体特性可以实现对象的组合特性，为结构体
// 添加各式各样的功能和特性，提高代码的可复用性和可扩展性。

// 内嵌不同结构体表现不同行为
// 鸭子们可以飞行，也可以游泳，但并不是所有的鸭子都会两种行为
package main

import "fmt"

// 游泳特性
type Swimming struct {
}

func (swim *Swimming) swim() {
	fmt.Println("swimming is my ability")
}

// 飞行特性
type Flying struct {
}

func (fly *Flying) fly() {
	fmt.Println("flying is my ability")
}

// 野鸭既可以飞行，也可以游泳；而家鸭只会游泳
// 在它们之间分别内嵌不同的结构体而体现不同的行为

// 野鸭，具备飞行和游泳特性
type WildDuck struct {
	Swimming
	Flying
}

// 家鸭，具备游泳特性
type DomesticDuck struct {
	Swimming
}

// 在 main 函数中查看它们是否具备相对应的行为特性
func main() {
	// 声明一只野鸭，可以飞，也可以游泳
	fmt.Println("野鸭:")
	wild := WildDuck{}
	wild.fly()
	wild.swim()

	// 声明一只家鸭，只会游泳
	fmt.Println("家鸭:")
	domestic := DomesticDuck{}
	domestic.swim()
}

/*
>>> Execution Result:
野鸭:
flying is my ability
swimming is my ability
家鸭:
swimming is my ability
*/
