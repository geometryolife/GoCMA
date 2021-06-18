package main

import (
	// 导入路径与 go.mod 文件相关，导入包的路径如果
	// 与包名不同，则会进行重命名
	compute "GoCMA/Chapter04/e3_compute"
	"fmt"
)

func main() {
	params := &compute.IntParams{
		P1: 1,
		P2: 2,
	}
	fmt.Println(params.Add())
}

/*
>>> Execution Result:
3
*/
