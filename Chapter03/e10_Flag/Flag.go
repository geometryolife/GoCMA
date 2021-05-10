// Golang的flag包中，命令行参数一般以指针的形式返回
package main

import (
	"flag"
	"fmt"
)

// 使用flag从命令行中读取参数
func main() {
	// 定义一个类型为string，名称为surname的命令行参数
	// 参数依次是命令行参数的名称，默认值，提示
	surname := flag.String("surname", "王", "您的姓")

	// 定义一个类型为string，名称为personalName的命令行参数
	// 除了返回指针类型结果，还可以直接传入变量地址获取参数值
	var personalName string
	flag.StringVar(&personalName, "personalName", "小二", "您的名")

	// 定义一个类型为int，名称为id的命令行参数
	id := flag.Int("id", 0, "您的ID")

	// 解析命令行参数
	flag.Parse()
	fmt.Printf("I am %v %v, and my ID is %v\n", *surname, personalName, *id)
}

// I am 王 小二, and my ID is 0

// Golang中的flag支持多种样式的命令行参数

// go run . -surname="Chen" -personalName="Joe" -id=100
// I am Chen Joe, and my ID is 100
// go run Flag.go --surname="Chen" --personalName="Joe" --id=101
// I am Chen Joe, and my ID is 101
// go run Flag.go -surname "Chen" -personalName "Joe" -id 102
// I am Chen Joe, and my ID is 102
// go run Flag.go --surname "Chen" --personalName "Joe" --id 103
// I am Chen Joe, and my ID is 103
