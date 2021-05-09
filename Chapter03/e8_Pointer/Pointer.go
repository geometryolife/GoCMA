package main

import "fmt"

func main() {
	// 声明一个string类型
	str := "Golang is Good!"
	// 获取str的指针
	strPtr := &str
	// 获取指针strPtr的指针
	strPtrPtr := &strPtr
	// 以此类推，多重指针
	strPtrPtrPtr := &strPtrPtr

	fmt.Printf("str type is %T, and value is %v\n", str, str)
	fmt.Printf("strPtr type is %T, and value is %v\n", strPtr, strPtr)
	fmt.Printf("strPtrPtr type is %T, and value is %v\n", strPtrPtr, strPtrPtr)
	fmt.Printf("strPtrPtr type is %T, and value is %v\n", strPtrPtrPtr,
		strPtrPtrPtr)
}

// str type is string, and value is Golang is Good!
// strPtr type is *string, and value is 0xc000010240
// strPtrPtr type is **string, and value is 0xc00000e030
// strPtrPtr type is ***string, and value is 0xc00000e038
