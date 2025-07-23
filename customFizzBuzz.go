package main

import (
	"fmt"
)

func customFizzBuzz(n int, rules map[int]string) string {
	output := ""
	for k, v := range rules {
		if n%k == 0 {
			output += v
		}
	}
	if output == "" {
		return fmt.Sprintf("%d", n)
	}
	return output
}

func main() {
	rules := map[int]string{3: "Fizz", 5: "Buzz", 7: "Qux", 11: "Pong"}
	for i := 1; i <= 100; i++ {
		fmt.Println(customFizzBuzz(i, rules))
	}
}
