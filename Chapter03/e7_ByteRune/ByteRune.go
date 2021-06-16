// 分别以byte和rune的方式遍历字符串
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	f := "Golang编程"
	// 中文字符在UTF-8编码中占3个字节
	// 统计字符串的字节长度，byte类型为"uint8"，代表一个ASCII字符
	fmt.Printf("byte lenth of f is %v\n", len(f))
	// utf8.RuneCountInString()方法统计字符串的Unicode字符数量
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

// byte lenth of f is 12
// rune lenth of f is 8

// Golangçç¨
// Golang编程

// 解释:
// 在进行字节遍历时，Unicode编码的中文字符会被截断，导致中文字符输出乱码
// byte & rune 的底层类型分别为uint8和int32
// rune能够处理一切的字符，而byte仅仅局限于处理ASCII字符
