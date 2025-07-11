package main

import (
	"fmt"
)

func uniqueString(s []string) []string {
	count := map[string]bool{}
	result := []string{}

	for _, v := range s {
		if !count[v] {
			count[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	s := []string{"apple", "hello", "世界", "banana", "世界"}
	ans := uniqueString(s)
	fmt.Println(ans)
}
