// 1. Golang 的接口设计是非侵入式的。
// 2. 对于接口编写者来说，他无需关心接口是被什么类型实现的；对于接口
// 实现者来说，他仅需知道实现的接口具备什么样的方法，而无需指定具体
// 实现哪一个接口。
// 3. 这种低耦合度的接口和实现之间的关联关系给了 Go 语言很大的灵活性。
// 4. Golang 中，要实现一个接口，需要满足以下两个条件:
// (1) 实现中添加的方法签名和接口签名完全一致，包括名称、参数列表、返
// 回参数等。
// (2) 接口的所有方法均被实现。

// 使用一个结构体同时实现 Cat 和 Dog 接口
package main

import "fmt"

// 定义 Cat 接口
type Cat interface {
	// 捉老鼠
	CatchMouse()
}

// 定义 Dog 接口
type Dog interface {
	// 吠叫
	Bark()
}

// 定义一个新物种，它既能像猫一样抓老鼠，也可以像狗一样吠叫，我们
// 将其定义为 CatDog 结构体，它将同时实现 Cat 和 Dog 两个接口
type CatDog struct {
	Name string
}

// 实现 Cat 接口
func (catDog *CatDog) CatchMouse() {
	fmt.Printf("%v caught the mouse and ate it!\n", catDog.Name)
}

// 实现 Dog 接口
func (catDog *CatDog) Bark() {
	fmt.Printf("%v barked loudly!\n", catDog.Name)
}

func main() {
	catDog := &CatDog{
		"Lucy",
	}

	// 声明一个 Cat 接口，并将 catDog 指针类型赋值给 cat
	var cat Cat
	cat = catDog
	cat.CatchMouse()

	// 声明一个 Dog 接口，并将 catDog 指针类型赋值给 dog
	var dog Dog
	dog = catDog
	dog.Bark()
}

/*
>>> Execution Result:
Lucy caught the mouse and ate it!
Lucy barked loudly!
*/

// 很显然，Lucy 正常发挥了它的作用，既发挥了猫抓老鼠的作用，也能像狗一样吠叫。
