package main

import (
	"fmt"
)

func countString(s []string) map[string]int {
	counts := make(map[string]int, len(s))

	for _, v := range s {
		counts[v]++
	}
	return counts
}

func main() {
	s := []string{"apple", "hello", "世界", "banana", "世界"}
	ans := countString(s)
	for key, value := range ans {
		fmt.Printf("%s:%d\n", key, value)
	}
}
