package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getPrime(n int) []int {
	prime := make([]int, 0, n/2)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			prime = append(prime, i)
		}
	}
	return prime
}

func main() {
	n := 100
	fmt.Printf("1~%d の素数判定: %v\n", n, getPrime(n))
}
