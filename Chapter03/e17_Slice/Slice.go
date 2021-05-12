// 切片是对数组一个连续片段的引用，它是一个容量可变的序列
// 切片类似于一个动态数组，通过指针引用底层数组
// 结构:
// 底层数组指针、大小、容量
// array, len, cap
package main

import "fmt"

func main() {
	sourceSlice()
	dynamicSlice()
	newSlice()
	appendSlice()
}

// 从原生数组中生成切片并修改切片成员
// slice := source[begin:end]
func sourceSlice() {
	source := [...]int{1, 2, 3}
	sli := source[0:1]

	// 仅能访问长度内的值
	fmt.Printf("sli value is %v\n", sli)
	fmt.Printf("sli len is %v\n", len(sli))
	fmt.Printf("sli cap is %v\n", cap(sli))

	// 切片作为指向原有数组的引用，对切片修改就是对原有数组进行修改
	sli[0] = 4
	fmt.Printf("sli value is %v\n", sli)
	fmt.Printf("source value is %v\n", source)
}

// 动态创建切片
// make函数动态创建切片，在创建过程中指定切片的长度和容量
// make([]T, size, cap)
func dynamicSlice() {
	sli := make([]int, 2, 4)
	fmt.Println()
	fmt.Printf("sli value is %v\n", sli)
	fmt.Printf("sli len is %v\n", len(sli))
	fmt.Printf("sli cap is %v\n", cap(sli))
}

// 声明新切片
// 直接声明新的切片类似于声明数组，但是不需要指定其大小
// var name []T{}
func newSlice() {
	// 切片分配了内存，但底层数组并未分配内存
	ex := []int{}
	fmt.Println()
	fmt.Printf("ex value is %v\n", ex)
	fmt.Printf("ex len is %v\n", len(ex))
	fmt.Printf("ex cap is %v\n", cap(ex))
	fmt.Printf("address of ex is %p\n", ex)

	// 声明切片并初始化
	// 切片大小和容量均为3
	ex2 := []int{1, 2, 3}
	fmt.Printf("ex2 value is %v\n", ex2)
	fmt.Printf("ex2 len is %v\n", len(ex2))
	fmt.Printf("ex2 cap is %v\n", cap(ex2))

	// 切片名是切片首地址的别名
	fmt.Println()
	fmt.Printf("address of ex2 is %p\n", ex2)
	fmt.Printf("address of &ex2[0] is %p\n", &ex2[0])
	fmt.Printf("address of &ex2[1] is %p\n", &ex2[1])
	fmt.Printf("address of &ex2[2] is %p\n", &ex2[2])
}

// 向切片添加元素
// 切片的动态扩容
func appendSlice() {
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4}

	sli1 := arr1[0:2]   // 长度为2，容量为4
	sli2 := arr2[2:4]   // 长度为2，容量为2
	sli3 := arr3[0:2:2] // 长度为2，容量为2

	fmt.Println()
	fmt.Printf("sli1 pointer is %p, len is %v, cap is %v, value is %v\n",
		&sli1, len(sli1), cap(sli1), sli1)
	fmt.Printf("sli2 pointer is %p, len is %v, cap is %v, value is %v\n",
		&sli2, len(sli2), cap(sli2), sli2)
	fmt.Printf("sli3 pointer is %p, len is %v, cap is %v, value is %v\n",
		&sli3, len(sli3), cap(sli3), sli3)

	// 容量足够的sli1直接将append添加的新元素覆盖到原有数组arr1中
	newSli1 := append(sli1, 5)
	fmt.Printf("newSli1 pointer is %p, len is %v, cap is %v, value is %v\n",
		&newSli1, len(newSli1), cap(newSli1), newSli1)
	fmt.Printf("source arr1 become %v\n", arr1)

	// 容量不够的sli2进行了扩容操作，申请新的底层数组，不在原数组的基础上操作
	newSli2 := append(sli2, 5)
	fmt.Printf("newSli2 pointer is %p, len is %v, cap is %v, value is %v\n",
		&newSli2, len(newSli2), cap(newSli2), newSli2)
	fmt.Printf("source arr2 become %v\n", arr2)

	// 如果原有数组可以添加新的元素，即切片指向的数组后还有空间，但切
	// 片自身的容量已经饱和，此时进行append操作，同样会进行扩容，申请
	// 新的内存空间
	newSli3 := append(sli3, 5)
	fmt.Printf("newSli3 pointer is %p, len is %v, cap is %v, value is %v\n",
		&newSli3, len(newSli3), cap(newSli3), newSli3)
	fmt.Printf("source arr3 become %v\n", arr3)
}

// sli value is [1]
// sli len is 1
// sli cap is 3
// sli value is [4]
// source value is [4 2 3]

// sli value is [0 0]
// sli len is 2
// sli cap is 4

// ex value is []
// ex len is 0
// ex cap is 0
// address of ex is 0x57c400
// ex2 value is [1 2 3]
// ex2 len is 3
// ex2 cap is 3

// address of ex2 is 0xc000018108
// address of &ex2[0] is 0xc000018108
// address of &ex2[1] is 0xc000018110
// address of &ex2[2] is 0xc000018118

// sli1 pointer is 0xc00000c0d8, len is 2, cap is 4, value is [1 2]
// sli2 pointer is 0xc00000c0f0, len is 2, cap is 2, value is [3 4]
// sli3 pointer is 0xc00000c108, len is 2, cap is 2, value is [1 2]
// newSli1 pointer is 0xc00000c168, len is 3, cap is 4, value is [1 2 5]
// source arr1 become [1 2 5 4]
// newSli2 pointer is 0xc00000c198, len is 3, cap is 4, value is [3 4 5]
// source arr2 become [1 2 3 4]
// newSli3 pointer is 0xc00000c1c8, len is 3, cap is 4, value is [1 2 5]
// source arr3 become [1 2 3 4]
