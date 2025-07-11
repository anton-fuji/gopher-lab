package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	minNum := math.MaxInt64
	maxNum := math.MinInt64
	sum := 0

	for i := 0; i < n; i++ {
		if s[i] > maxNum {
			maxNum = s[i]
		}
		if s[i] < minNum {
			minNum = s[i]
		}
		sum += s[i]
	}

	fmt.Println(maxNum, minNum, sum)

}
