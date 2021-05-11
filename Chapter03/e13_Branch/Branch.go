// Golang的分支控制与其他语言相似，但更简略
// 表达式两边可省略()
/*
if expression1 {
	branch1
} else if expression2 {
	branch2
} else {
	branch3
}
*/
package main

import "fmt"

func main() {
	age := 18

	if age >= 0 && age <= 6 {
		fmt.Printf("You don't need to pay.")
	} else if age <= 18 {
		fmt.Printf("You need to pay 10$.")
	} else {
		fmt.Printf("You need to pay 15$.")
	}
}

// You need to pay 10$.
