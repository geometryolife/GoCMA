package main

import (
	compute "GoCMA/Chapter04/e3_compute"
	"fmt"
)

func main() {
	params := &compute.IntParams{
		P1: 1,
		P2: 2,
	}
	fmt.Println(params.Add())
}
