package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibo(n int) []int {
	fibos := make([]int, n)
	switch n {
	case 0:
		return []int{}
	case 1:
		return []int{0}
	}

	fibos[0] = 0
	fibos[1] = 1

	for i := 2; i < n; i++ {
		fibos[i] = fibos[i-1] + fibos[i-2]
	}
	return fibos
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(sc, &n)

	ans := fibo(n)
	fmt.Println(ans)
}
