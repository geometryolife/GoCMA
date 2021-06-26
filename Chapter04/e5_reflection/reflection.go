// 1. 反射是一项功能强大的工具，它给开发人员提供了在运行时对代码本身
// 进行访问和修改的能力。通过反射，我们可以拿到丰富的类型信息，比如
// 变量的字段名称、类型信息和结构体信息等，并使用这些类型信息做一些
// 灵活的工作。
// 2. Golang 的反射实现了反射的大多数功能，获取类型的信息需要配合使
// 用标准库中的词法、语法解析器和抽象语法树对源码进行扫描。
// 3. Golang 的反射主要通过 Type 和 Value 两个基本概念来表达。
// 4. Type 主要用于表示被反射变量的类型信息，而 Value 用于表示被反射
// 变量自身的实例信息。
// 5. 通过 reflect.TypeOf 方法，我们可以获取一个变量的类型信息
// reflect.Type。
// 6. 通过 reflect.Type 类型对象，我们可以访问到其对应类型的各项类型
// 信息。我们可以创建一个 Hero 结构体，通过 reflect.TypeOf 来查看其
// 对应的类型信息。
// 7. Golang 中，存在着 Type (类型)和 Kind (种类)的区别。
// 8. Type 是指变量所属的类型，包括系统的原生数据类型(如 int、string
// 等)和我们通过 type 关键字定义的类型，比如我们定义的 Hero 结构体，
// 这些类型的名称一般就是其类型本身。而 Kind 是指变量类型所归属的品
// 种，参考 reflect.Kind 中的定义，主要有以下类型:
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
/*
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)
*/
// 9. 一般我们通过 type 关键字定义的结构体都属于 Struct，而指针变量的
// 种类统一为 Ptr。
// 10. 对于指针类型的变量，可以使用 Type.Elem 获取到指针指向变量的真
// 实类型对象。
package main

import (
	"fmt"
	"reflect"
)

// 定义一个人的接口
type Person interface {
	// 与人说 hello
	SayHello(name string)
	// 跑步
	Run() string
}

// 定义 Hero 结构体来实现 Person 接口中的方法，同时 Hero 结构
// 体中还包含3个成员字段
type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello "+name, ", I am "+hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed", hero.Speed)
	return "Running"
}

func main() {
	// 获取实例的反射类型对象
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Printf("Hero's type is %s, kind is %s.\n", typeOfHero, typeOfHero.Kind())
	fmt.Printf("*Hero's type is %s, kind is %s.\n\n", reflect.TypeOf(&Hero{}),
		reflect.TypeOf(&Hero{}).Kind())

	typeOfPtrHero := reflect.TypeOf(&Hero{})
	fmt.Printf("*Hero's type is %s, kind is %s.\n", typeOfPtrHero,
		typeOfPtrHero.Kind())
	fmt.Printf("typeOfPtrHero elem to typeOfHero, Hero's type is %s, kind is %s.\n",
		typeOfHero, typeOfHero.Kind())
}

/*
>>> Execution Result:
Hero's type is main.Hero, kind is struct.
*Hero's type is *main.Hero, kind is ptr.

*Hero's type is *main.Hero, kind is ptr.
typeOfPtrHero elem to typeOfHero, Hero's type is main.Hero, kind is struct.
*/
