// 分别以 byte 和 rune 的方式遍历字符串
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	f := "Golang编程"
	// 中文字符在 UTF-8 编码中占3个字节
	// 统计字符串的字节长度，byte 类型为"uint8"，代表一个 ASCII 字符
	fmt.Printf("byte lenth of f is %v\n", len(f))
	// utf8.RuneCountInString() 方法统计字符串的 Unicode 字符数量
	fmt.Printf("rune lenth of f is %v\n", utf8.RuneCountInString(f))

	fmt.Println()

	// 按字节遍历字符串
	for _, b := range []byte(f) {
		fmt.Printf("%c", b)
	}

	fmt.Println()

	// 按字符遍历字符串
	for _, r := range f {
		fmt.Printf("%c", r)
	}
}

/*
>>> Execution Result:
byte lenth of f is 12
rune lenth of f is 8

Golangçç¨
Golang编程
*/

// 解释:
// 在进行字节遍历时，Unicode 编码的中文字符会被截断，导致中文字
// 符输出乱码，byte 与 rune 的底层类型分别为 uint8 和 int32
// rune 能够处理一切的字符，而 byte 仅仅局限于处理 ASCII 字符
