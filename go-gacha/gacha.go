package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	var kind int
	fmt.Println("1: 単発ガチャ 2: 11連ガチャ")
	fmt.Print(">> ")

LOOP:
	for {
		_, err := fmt.Scan(&kind)
		if err != nil {
			fmt.Println("入力エラーです:", err)
			continue
		}
		switch kind {
		case 1:
			n = 1
			break LOOP
		case 2:
			n = 11
			break LOOP
		default:
			fmt.Println("もう一度入力してください")
		}
	}

	for i := 0; i < n; i++ {
		num := rand.Intn(100)
		fmt.Printf("%d回目\n", i+1)

		switch {
		case num < 80:
			fmt.Println("N")
		case num < 95:
			fmt.Println("R")
		case num < 99:
			fmt.Println("SR")
		default:
			fmt.Println("SSR")

		}
	}
	fmt.Println("\nガチャ終了!")
}
