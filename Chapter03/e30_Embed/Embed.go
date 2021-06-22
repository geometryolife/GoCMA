// 1. 在结构体定义时，Golang 允许声明没有字段名的字段，这种形式的字段
// 被称为类型内嵌或匿名字段，此时字段名就是字段类型本身，由于结构体要
// 求字段名称必须唯一，因此同一类型的类型内嵌在同一结构体只能存在一个:
/*
type temp struct {
	string
	int
}
*/
// 2. 如果内嵌的类型为结构体，就可以直接访问内嵌结构体中的所有成员:
package main

import "fmt"

type Wheel struct {
	shape string
}

// 将 Wheel 结构体内嵌到 Car，就可以通过 Car 的引
// 用直接访问到 Wheel 结构体中的成员属性
type Car struct {
	Wheel
	Name string
}

func main() {
	// 通过取址实例化来初始化 car
	car := &Car{
		Wheel{
			"圆形的",
		},
		"福特",
	}
	fmt.Println(car.Name, car.shape, " ") // 福特 圆形的
	fmt.Println(car)
}

/*
>>> Execution Result:
福特 圆形的
&{{圆形的} 福特}
*/
