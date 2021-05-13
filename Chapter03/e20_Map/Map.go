// Golang中的字典(map)为映射关系容器，内部通过散列表实现
// name := make(map[keyType]valueType)
// map需要使用make函数进行初始化
package main

import "fmt"

func main() {
	classMates1 := make(map[int]string)

	// 添加映射关系
	classMates1[0] = "小明"
	classMates1[1] = "小红"
	classMates1[2] = "小张"

	// 根据key获取value
	fmt.Printf("id %v is %v\n", 1, classMates1[1])

	// 在声明时初始化数据
	// 以类JSON格式添加键值对映射关系
	classMates2 := map[int]string{
		0: "小明",
		1: "小红",
		2: "小张",
	}

	// 通过map的键查询对应的值
	fmt.Printf("id %v is %v\n", 2, classMates2[2])
	// 如果键不存在，返回值类型的默认值
	fmt.Printf("id %v is %v\n", 3, classMates2[3])

	// 查询某个键是否存在于map中
	mate, ok := classMates2[1]
	fmt.Printf("%v, %v\n", mate, ok)
	mate, ok = classMates2[3]
	fmt.Printf("%v, %v\n", mate, ok)
}

// id 1 is 小红
// id 2 is 小张
// id 3 is
// 小红, true
// , false
