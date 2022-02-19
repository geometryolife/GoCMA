// 获取一个结构体内的字段数量
// NumField() int
// 根据 index 获取结构体内的成员字段类型对象
// Field(i int) StructField
// 根据字段名获取结构体内的成员字段类型对象
// FieldByName(name string) (StructField, bool)
//
// type StructField struct {
//     Name string // 成员字段的名称
//     Type Type // 成员字段的类型
//     Tag StructTag // 标签
//     Offset uintptr // 字节偏移
//     Index []int // 成员字段的 index
//     Anonymous bool // 成员字段是否公开
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
	typeOfHero := reflect.TypeOf(Hero{})

	// 通过 NumField 方法获取结构体字段的数量
	for i := 0; i < typeOfHero.NumField(); i++ {
		fmt.Printf("field's name is %s, type is %s, kind is %s\n",
			typeOfHero.Field(i).Name,
			typeOfHero.Field(i).Type,
			typeOfHero.Field(i).Type.Kind())
	}
	// === Output ===
	// field's name is Name, type is string, kind is string
	// field's name is Age, type is int, kind is int
	// field's name is Speed, type is int, kind is int

	// 获取名称为 Name 的成员字段类型对象
	nameField, _ := typeOfHero.FieldByName("Name")
	fmt.Printf("field's name is %s, type is %s, kind is %s\n",
		nameField.Name, nameField.Type, nameField.Type.Kind())
	// === Output ===
	// field's name is Name, type is string, kind is string
}
