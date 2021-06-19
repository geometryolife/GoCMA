// 1. Golang 的列表通过双向链表的方式实现，故通俗称其为链表。
// 2. 适用于存储需要经常进行元素插入和删除操作的元素集合。
// 3. 双向链表中每个元素都持有其前后元素的引用，在链表进行向
// 前或者向后遍历时会非常方便。
// 注意:
// 进行元素插入和删除时需要同时修改前后元素持有引用。
// Golang 的 list 是一个带头结点的链表，头结点中不存储数据。
// var name list.List
// or
// name := list.New()
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 初始化列表，返回列表对应的指针
	tmpList := list.New()

	for i := 1; i <= 10; i++ {
		tmpList.PushBack(i)
	}

	first := tmpList.PushFront(0)
	tmpList.Remove(first)

	for l := tmpList.Front(); l != nil; l = l.Next() {
		fmt.Print(l.Value, " ")
	}
}

/*
>>> Execution Result:
1 2 3 4 5 6 7 8 9 10
*/
