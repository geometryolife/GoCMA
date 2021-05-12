// 数组是一段存储固定类型固定长度的连续内存空间
// var name [size]T
// 数组必须指定大小，可以是一个常量或者表达式，但必
// 须在静态编译时就确定其大小，不能动态指定。
// "..."让编译器根据{}内的成员数量来确定数组大小
package main

import "fmt"

func main() {
	var classMates1 [3]string
	classMates1[0] = "小明"
	classMates1[1] = "小红"
	classMates1[2] = "小李" // 通过下标为数组成员赋值
	fmt.Println(classMates1)
	// 通过下标访问数组成员
	fmt.Println("The No.1 student is " + classMates1[0])

	// 使用初始化列表初始化数组
	classMates2 := [...]string{"小明", "小红", "小李"}
	fmt.Println(classMates2)

	// 使用指针操作数组
	classMates3 := new([3]string)
	classMates3[0] = "小明"
	classMates3[1] = "小红"
	classMates3[2] = "小李"
	fmt.Println(*classMates3)

	// 与C语言不同，数组名不是数组首地址的别名
	fmt.Printf("\nType of classMates1 is %T\n", classMates1)
	// 在Golang中数组名代表的是整个数组的值，而非地址别名
	// fmt.Printf("Address of classMates1 is %p\n", classMates1)

	// 指向数组的指针
	ptr := &classMates1
	fmt.Printf("%v\n", *ptr)
	fmt.Printf("classMates1[1] is %v\n", ptr[1])
}

// [小明 小红 小李]
// The No.1 student is 小明
// [小明 小红 小李]
// [小明 小红 小李]

// Type of classMates1 is [3]string
// [小明 小红 小李]
// classMates1[1] is 小红
