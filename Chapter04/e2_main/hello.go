package main

import "fmt"

// 同一个包下的代码随意访问
func sayHello() {
	fmt.Println("Hello World")
}
