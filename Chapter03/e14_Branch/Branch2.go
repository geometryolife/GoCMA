// 1. 除了 if 关键字，Golang 中还提供 switch 语句对大量的值
// 和表达式进行判断。switch 中每个 case 都是独立代码块，不
// 需要通过 break 关键字跳出 switch 选择体。如果想继续执行
// 接下来的 case 判断，需要添加 fallthrough 关键字对上下两
// 个 case 进行连接。
// 2. 除了支持数值常量，Golang 的 switch 还能对字符串、表达
// 式等复杂情况进行处理。
package main

import "fmt"

func main() {
	// 根据人名分配工作
	name := "小红"
	// 每个 case 都是字符串样式，且无需通过 break 控制跳出
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
	// 在 case 中判断表达式的情况下 switch 后面不需要指定判断变量
	// 此形式和 if-else 相似
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

/*
>>> Execution Result:
擦黑板
优秀
*/
