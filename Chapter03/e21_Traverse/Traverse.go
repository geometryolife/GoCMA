// 对给出的数组 nums、切片 slis 和字典 tmpMap 分别进行遍历
package main

import "fmt"

func main() {
	// 数组的遍历
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range nums {
		// k 为下标，v 为对应的值
		fmt.Println(k, v)
	}

	fmt.Println()

	// 切片的遍历
	slis := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range slis {
		// k 为下标，v 为对应的值
		fmt.Println(k, v)
	}

	fmt.Println()

	// 字典的遍历
	tmpMap := map[int]string{
		0: "小明",
		1: "小红",
		2: "小张",
	}

	for k, v := range tmpMap {
		// k 为键，v 为对应的值
		fmt.Println(k, v)
	}

	fmt.Println()

	// 借助匿名变量仅遍历值
	for _, v := range nums {
		// k 为下标，v 为对应的值
		fmt.Println(v)
	}

	fmt.Println()

	// 仅遍历键时，可以直接省略值的赋值
	for k := range nums {
		fmt.Println(k)
	}
}

/*
>>> Execution Result:
0 1
1 2
2 3
3 4
4 5
5 6
6 7
7 8

0 1
1 2
2 3
3 4
4 5
5 6
6 7
7 8

0 小明
1 小红
2 小张

1
2
3
4
5
6
7
8

0
1
2
3
4
5
6
7
*/
