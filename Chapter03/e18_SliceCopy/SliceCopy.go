// 1. 为了方便切片的数据快速复制到另一个切片中，Golang 提供
// 了内建的 copy 函数:
// copy(dst, src []T) int
package main

import "fmt"

func main() {
	sli := make([]int, 2, 5)
	sli[0] = 1
	sli[1] = 2

	// 切片名是其底层数组的首地址
	fmt.Printf("address of sli: %p\n", sli)

	// 切片的指针
	fmt.Printf("address of &sli: %p\n", &sli)

	// 切片引用底层数组的指针
	fmt.Printf("address of &sli[0]: %p\n", &sli[0])
	fmt.Printf("address of &sli[1]: %p\n", &sli[1])
	sli = append(sli, 3)
	fmt.Printf("address of &sli[2]: %p\n", &sli[2])

	fmt.Println(sli)
	fmt.Println()

	// 需要保证目标切片的长度不小于来源切片的长度，否则无法完整复制
	// 使用 make 必定会创建新的底层数组
	sli2 := make([]int, 3, 5)
	num := copy(sli2, sli)
	fmt.Printf("address of &sli[0]: %p, value of sli: %v\n", &sli[0], sli)
	fmt.Printf("address of &sli2[0]: %p, value of sli2: %v\n", &sli2[0], sli2)
	// 返回值是实际发生复制的元素个数
	fmt.Printf("length of sli2 is %v\n", num)

	sli3 := sli2[0:2:2]
	fmt.Printf("address of &sli3[0]: %p, value of sli3: %v\n", &sli3[0], sli3)
	sli3 = append(sli3, 10)
	fmt.Printf("address of &sli3[0]: %p, value of sli3: %v\n", &sli3[0], sli3)
}

/*
>>> Execution Result:
address of sli: 0xc000016240
address of &sli: 0xc00000c060
address of &sli[0]: 0xc000016240
address of &sli[1]: 0xc000016248
address of &sli[2]: 0xc000016250
[1 2 3]

address of &sli[0]: 0xc000016240, value of sli: [1 2 3]
address of &sli2[0]: 0xc000016270, value of sli2: [1 2 3]
length of sli2 is 3
address of &sli3[0]: 0xc000016270, value of sli3: [1 2]
address of &sli3[0]: 0xc000018180, value of sli3: [1 2 10]
*/
