// 最后的输出可能是 0、1、2，大多数情况是 0，这取决于当时调度器的调度情况
package main

import (
	"fmt"
)

func setVTo1(v *int) {
	*v = 1
}

func setVTo2(v *int) {
	*v = 2
}

func main() {
	v := new(int)
	go setVTo1(v)
	// 阻塞主 goroutine 1 秒，可以让 setVTo1 最后执行，输出 1
	// time.Sleep(time.Second)
	go setVTo2(v)
	fmt.Println(*v)
}

/*
>>> Execution Result:
0
*/
