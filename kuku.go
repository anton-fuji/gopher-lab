package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		// fmt.Printf("======= %d の段 ========\n", i)
		for j := 1; j <= 9; j++ {
			fmt.Printf("%dx%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}
}
