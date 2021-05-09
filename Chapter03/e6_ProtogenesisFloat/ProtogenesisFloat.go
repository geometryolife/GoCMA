package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", math.E)     // 按照默认宽度和精度输出
	fmt.Printf("%.2f\n", math.E)   // 按照默认宽度和2位精度输出
	fmt.Printf("%10.2f\n", math.E) // 指定10位宽度和2位精度输出，右对齐
}

// 2.718282
// 2.72
//       2.72

// float32 & float64 转换时注意精度的损失

// 额外:
// Golang的布尔型不可以与整型强转，也无法参与数值运算
