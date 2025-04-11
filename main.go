package main

import (
	"fmt"
)

func main() {
	var a, b int
	var op string
	fmt.Scan(&a, &op, &b)

	if op == "?" {
		fmt.Println("変なの入れるな!")
	}
	switch {
	case op == "+":
		fmt.Println(a + b)
	case op == "-":
		fmt.Println(a - b)
	case op == "*":
		fmt.Println(a * b)
	case op == "/":
		fmt.Println(a / b)
	}
}
