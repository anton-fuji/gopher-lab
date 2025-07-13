package main

import (
	"fmt"
	"math/big"
)

func fibo(n int) []*big.Int {
	switch n {
	case 0:
		return []*big.Int{}
	case 1:
		return []*big.Int{big.NewInt(0)}
	}

	if n < 0 {
		return nil
	}

	fibos := make([]*big.Int, n)
	fibos[0] = big.NewInt(0)
	fibos[1] = big.NewInt(1)
	for i := 2; i < n; i++ {
		fibos[i] = new(big.Int).Add(fibos[i-2], fibos[i-1])
	}
	return fibos
}

func main() {
	var n int
	fmt.Scan(&n)

	for i, v := range fibo(n) {
		fmt.Printf("fibo[%d] = %s\n", i, v)
	}
}
