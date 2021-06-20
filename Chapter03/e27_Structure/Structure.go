// 1. 结构体作为一种复合类型，由多个字段组成，每个字段都具备自己
// 的类型和值，结构体和字段可以理解为实体和实体对应的属性。
// 2. 在 Golang 中，不仅结构体可以拥有方法，每一种自定义类型也可
// 以拥有方法。
// 3. 配合使用 type 和 struct 关键字，可以自定义结构体。
// 4. Golang 的 type 关键字可以将各种基本类型定义为自定义类型，
// 结构体中可以复合多种基本类型和结构体，更便于使用。
/*
type structName struct {
	value1 valueType1
	value2 valueType2
	...
}
*/
// 5. 结构体的名称在同一个包内不能重复。
// 6. 在使用结构体之前需要对其进行实例化和初始化。实例化是为我
// 们创建的结构体在内存中分配具体的内存进行存储；而初始化则是为
// 刚刚实例化好的结构体的字段赋予初始值，用于特例化该结构体。
// 7. 每个结构体实例之间的内存区域是相互独立的。
package main

import "fmt"

type Person struct {
	Name  string // 姓名
	Birth string // 生日
	ID    int64  // 身份证号
}

func main() {
	// 实例化
	// 声明实例化
	var p1 Person
	p1.Name = "Joe1"
	p1.Birth = "1997-01-01"

	// new 函数实例化
	p2 := new(Person)
	p2.Name = "Joe2"
	p2.Birth = "1997-01-02"

	// 取址实例化
	p3 := &Person{}
	p3.Name = "Joe3"
	p3.Birth = "1997-01-03"

	// 初始化(类似 JSON 的键值对的形式填充结构体的字段)
	p4 := Person{
		Name:  "Joe4",
		Birth: "1997-01-04",
	}

	// 初始化(填充顺序与字段顺序一致，可省略键)
	p5 := &Person{
		"Joe5",
		"1997-01-05",
		5,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)
}

// 解释:
// 在上述例子中，我们初始化 p5 时并没有将键值一一对应填充给 Person
// 结构体，仅把对应的值按照结构体内定义的顺序填充给了 Person 结构体，
// 而省略了其中的键。初始化 p5 使用的是取址实例化，返回的 p5类型为
// Person 结构体的指针类型，即 *Person。

/*
>>> Execution Result:
{Joe1 1997-01-01 0}
&{Joe2 1997-01-02 0}
&{Joe3 1997-01-03 0}
{Joe4 1997-01-04 0}
&{Joe5 1997-01-05 5}
*/
