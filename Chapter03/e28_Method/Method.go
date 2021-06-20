// 1. 在 Golang 中，方法是有特定接收器的函数，接收器可以是任意类型而
// 不仅仅是结构体，即 Golang 的任何类型都可以拥有自己的方法。
// 2. 与普通函数相比，方法的定义中多了一个接收器的设定，每一个方法只
// 能有一个接收器，接收器的概念就类似于面向对象语言中的 this 或者
// self，我们可以直接在方法中使用接收器的相关属性。
// 3. 接收器有两种类别: 指针类型 与 非指针(值)类型，它们在使用时会产
// 生不同的效果，会被使用在不同的应用场景中。
// 4. 指针类型的接收器传递的是类型的指针，通过该指针可以在方法中操作
// 接收器内的成员属性，操作结果在方法结束后依然存在于接收器中，因为
// 指针操作的是接收器的内存区域；而非指针类型的接收器传递的是方法调
// 用时接收器的一份值拷贝，对该接收器的成员属性进行操作并不会影响到
// 原接收器。
// 5. 当接收器占用内存较大或者需要对原接收器的成员属性进行修改时，建
// 议使用指针类型接收器；如果接收器占用内存较小，且方法对其仅需要只
// 读功能，可以采用非指针接收器。

// 为 Person 结构体添加修改姓名和输出个人信息的两个方法
package main

import "fmt"

type Person struct {
	Name  string // 姓名
	Birth string // 生日
	ID    int64  // 身份证号
}

// 指针类型，修改个人信息
func (person *Person) changeName(name string) {
	person.Name = name
}

// 非指针类型，打印个人信息
func (person Person) printMess() {
	fmt.Printf("My name is %v, and my birthday is %v, and my id is %v\n",
		person.Name, person.Birth, person.ID)

	// 尝试修改个人信息，但是对原接收器并没有影响
	person.ID = 3
}

func main() {
	p1 := Person{
		Name:  "Joe",
		Birth: "1997-01-01",
		ID:    1,
	}

	p1.printMess()

	p1.changeName("Tom")
	p1.printMess()
}

/*
>>> Execution Result:
My name is Joe, and my birthday is 1997-01-01, and my id is 1
My name is Tom, and my birthday is 1997-01-01, and my id is 1
*/
