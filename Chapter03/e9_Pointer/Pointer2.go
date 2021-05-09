package main

import "fmt"

func main() {
	str := "Golang is Good!"
	strPtr := &str

	fmt.Printf("str type is %T, value is %v, address is %p\n", str, str, &str)
	fmt.Printf("strPtr type is %T, and value is %v\n", strPtr, strPtr)

	newStr := *strPtr // 获取指针对应变量的值
	fmt.Printf("newStr type is %T, value is %v, and address is %p\n", newStr,
		newStr, &newStr)

	*strPtr = "Java is Good too!" // 通过指针对变量进行赋值
	fmt.Printf("newStr type is %T, value is %v, and address is %p\n", newStr,
		newStr, &newStr)
	fmt.Printf("str type is %T, value is %v, address is %p\n", str, str, &str)

	// 除了使用&对变量进行取址操作创建指针，还可以使用
	// new函数直接分配内存，并返回指向内存的指针，此时
	// 内存中的值会被初始化为类型的默认值。
	fmt.Println()
	str2 := new(string)
	*str2 = "Golang is my favorite language."
	fmt.Printf("str2 type is %T, and value is %v\n", str2, str2)
	fmt.Printf("*str2 type is %T, and value is %v\n", *str2, *str2)
}

// str type is string, value is Golang is Good!, address is 0xc000010240
// strPtr type is *string, and value is 0xc000010240
// newStr type is string, value is Golang is Good!, and address is 0xc000010270
// newStr type is string, value is Golang is Good!, and address is 0xc000010270
// str type is string, value is Java is Good too!, address is 0xc000010240

// str2 type is *string, and value is 0xc0000102e0
// *str2 type is string, and value is Golang is my favorite language.

// 取址操作(&)，引用
// 取值操作(*)，解引用
