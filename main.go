package main

import (
	"fmt"
)

func uniques[T comparable](s []T) []T {
	count := map[T]struct{}{}
	result := []T{}

	for _, v := range s {
		if _, ok := count[v]; !ok {
			count[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
func main() {
	s := []string{"apple", "banana", "oarange", "apple", "peach"}
	i := []int{1, 2, 4, 2, 4, 6, 2, 8, 3}

	fmt.Println(uniques(s))
	fmt.Println(uniques(i))
}
