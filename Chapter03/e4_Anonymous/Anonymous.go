package main

import "fmt"

// 返回一个人的姓和名
func getName() (string, string) {
	return "Chen", "Joe"
}

func main() {
	// 使用匿名变量"_"替换不需要的变量名
	// 匿名变量不占用命名空间，不分配内存
	// 匿名变量不会因为多次声明而无法使用
	surname, _ := getName()      // 使用匿名变量
	_, personalName := getName() // 使用匿名变量

	fmt.Printf("My surname is %v and my personal name is %v", surname,
		personalName)
}
