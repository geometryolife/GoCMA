// 整型: 有符号整型 & 无符号整型(指定长度)
// int8, int16, int32, int64
// uint8, uint16, uint32, uint64
// Golang提供平台自匹配长度的int型和uint型
package main

import (
	"fmt"
	"math"
)

func main() {
	var e uint16 = math.MaxUint8 + 1
	fmt.Printf("e value(uint16) is %v\n", e)

	var d = uint8(e)
	fmt.Printf("d value(uint8) is %v\n", d)
}

// e value(uint16) is 256
// d value(uint8) is 0

// 解释:
// 高长度类型向低长度类型转换，值会被截断，仅保留低位值，造成转换错误。
// 256在"uint16"底层存储方式为"00000001 00000000"，转换为"uint8"之后，
// 只截取后8位，故d为0。
