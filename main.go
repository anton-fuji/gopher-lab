package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	minNum := math.MaxInt64
	maxNum := math.MinInt64
	sum := 0

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		if a > maxNum {
			maxNum = a
		}
		if a < minNum {
			minNum = a
		}
		sum += a
	}
	fmt.Printf("Max: %d Min: %d Sum: %d", maxNum, minNum, sum)
}
