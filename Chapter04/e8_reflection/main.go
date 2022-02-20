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
	name := "小明"
	valueOfName := reflect.ValueOf(name) // 获取 name 的 Value 对象（值）
	fmt.Printf("name variable's object is %v\n", valueOfName)
	fmt.Println(valueOfName.Interface())
	// === Output ===
	// name variable's object is 小明
	// 小明

	// 如果取值的类型与取值的方式不匹配，那么程序就会 panic
	// valueOfName2 := reflect.ValueOf(name)
	// fmt.Println(valueOfName2.Bytes())
	// === Output ===
	// panic: reflect: call of reflect.Value.Bytes on string Value

	typeOfHero := reflect.TypeOf(Hero{})
	// reflect.New 方法根据变量的 Type 创建一个相同类型的变量
	heroValue := reflect.New(typeOfHero)
	fmt.Printf("Hero's type is %s, kind is %s\n", heroValue.Type(), heroValue.Kind())
	// === Output ===
	// Hero's type is *main.Hero, kind is ptr

	valueOfName3 := reflect.ValueOf(&name) // 获取 name 的 Value 指针
	fmt.Printf("name's pointer is %v\n", valueOfName3)
	// === Output ===
	// name's pointer is 0xc000010200

	// Elem 方法解引用获取 name 的 Value 类型值，Set 方法对值进行修改
	valueOfName3.Elem().Set(reflect.ValueOf("小红"))
	fmt.Println(name)
	// === Output ===
	// 小红

	// CanAddr 方法判断 reflect.Value 类型变量是否可以寻址
	valueOfName4 := reflect.ValueOf(name)
	fmt.Printf("name can be address: %t\n", valueOfName4.CanAddr()) // false
	valueOfName4 = reflect.ValueOf(&name)
	fmt.Printf("&name can be address: %t\n", valueOfName4.CanAddr()) // false
	valueOfName4 = valueOfName4.Elem()
	fmt.Printf("&name's Elem can be address: %t\n", valueOfName4.CanAddr()) // true
	// === Output ===
	// name can be address: false
	// &name can be address: false
	// &name's Elem can be address: true

	hero := &Hero{
		Name: "小白",
	}

	valueOfHero := reflect.ValueOf(hero).Elem()
	valueOfName5 := valueOfHero.FieldByName("Name")

	// 判断字段的 Value 是否可以设定变量值
	if valueOfName5.CanSet() {
		valueOfName5.Set(reflect.ValueOf("小张"))
	}

	fmt.Printf("hero name is %s\n", hero.Name)
	// === Output ===
	// hero name is 小张

	// 【例4-1】使用反射调用接口方法
	var person Person = &Hero{
		Name:  "小红",
		Speed: 100,
	}

	valueOfPerson := reflect.ValueOf(person)

	// 获取 SayHello 方法
	sayHelloMethod := valueOfPerson.MethodByName("SayHello")
	// 构建调用参数并通过反射调用 Call 方法
	sayHelloMethod.Call([]reflect.Value{reflect.ValueOf("小张")})
	// === Output ===
	// Hello 小张 , I am  小红

	// 获取 Run 方法
	runMethod := valueOfPerson.MethodByName("Run")
	// 通过 Call 调用方法并获取结果
	result := runMethod.Call([]reflect.Value{})
	fmt.Printf("result of run method is %s\n", result[0])
	// === Output ===
	// I am running at speed 100
	// result of run method is Running

	// 【使用 Type 获取】
	var person2 Person = &Hero{
		Name: "小红",
	}

	// 获取接口 Person 的类型对象
	typeOfPerson := reflect.TypeOf(person2)
	// 打印 Person 的方法类型和名称
	sayHelloMethod2, _ := typeOfPerson.MethodByName("Run")
	// fmt.Println(sayHelloMethod2)
	// 将 person 接收器放在参数的第一位
	result2 := sayHelloMethod2.Func.Call([]reflect.Value{reflect.ValueOf(person)})
	fmt.Printf("result2 of run method is %s.\n", result2[0])
	// === Output ===
	// I am running at speed 100
	// result2 of run method is Running.

	// 【一般方法】
	methodOfHello := reflect.ValueOf(hello)
	methodOfHello.Call([]reflect.Value{})
	// === Output ===
	// Hello World!
}

func hello() {
	fmt.Printf("Hello World!")
}
