// 根据 index 查找方法
// Method(int) Method
// 根据方法名查找方法
// MethodByName(string) (Method, bool)
// 获取类型中公开的方法数量
// NumMethod() int
//
// type Method struct {
//     Name string // 方法名
//     Type Type // 方法类型
//     Func Value // 反射对象，可用于调用方法
//     Index int // 方法的 index
// }
package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	SayHello(name string)
	Run() string
}

type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello "+name, ", I am ", hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed", hero.Speed)
	return "Running"
}

func main() {
	// 声明一个 Person 接口，并用 Hero 作为接收器
	var person Person = &Hero{}

	// 获取接口 Person 的类型对象
	typeOfPerson := reflect.TypeOf(person)

	// 打印 Person 的方法类型和名称
	for i := 0; i < typeOfPerson.NumMethod(); i++ {
		fmt.Printf("method is %s, type is %s, kind is %s.\n",
			typeOfPerson.Method(i).Name,
			typeOfPerson.Method(i).Type,
			typeOfPerson.Method(i).Type.Kind())
	}
	// === Output ===
	// method is Run, type is func(*main.Hero) string, kind is func.
	// method is SayHello, type is func(*main.Hero, string), kind is func.

	method, _ := typeOfPerson.MethodByName("Run")
	fmt.Printf("method is %s, type is %s, kind is %s.\n",
		method.Name, method.Type, method.Type.Kind())
	// === Output ===
	// method is Run, type is func(*main.Hero) string, kind is func.
}
