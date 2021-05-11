// 除了if关键字，Golang中还提供switch语句对大量的值
// 和表达式进行判断。switch中每个case都是独立代码块，
// 不需要通过break关键字跳出switch选择体。如果想继续
// 执行接下来的case判断，需要添加fallthrough关键字对
// 上下两个case进行连接。
// 除了支持数值常量，Golang的switch还能对字符串、表达
// 式等复杂情况进行处理。
package main

import "fmt"

func main() {
	// 根据人名分配工作
	name := "小红"
	// 每个case都是字符串样式，且无需通过break控制跳出
	switch name {
	case "小明":
		fmt.Println("扫地")
	case "小红":
		fmt.Println("擦黑板")
	case "小刚":
		fmt.Println("倒垃圾")
	default:
		fmt.Println("没人干活")
	}

	// 根据分数判断成绩程度
	score := 90
	// 在case中判断表达式的情况下switch后面不需要指定判断变量
	// 此形式和if-else相似
	switch {
	case score < 100 && score >= 90:
		fmt.Println("优秀")
	case score < 90 && score >= 80:
		fmt.Println("良好")
	case score < 80 && score >= 60:
		fmt.Println("及格")
	case score < 60 && score >= 0:
		fmt.Println("不及格")
	default:
		fmt.Println("分数错误")
	}
}

// 擦黑板
// 优秀
