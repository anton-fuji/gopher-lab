package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		fmt.Println("======", i, "======")
		for j := 1; j <= 9; j++ {
			fmt.Println(i, "*", j, "= ", i*j)
		}
	}
}
