package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	if i%2 == 0 {
		fmt.Printf("%d is Even", i)
	} else {
		fmt.Printf("%d is Odd", i)
	}
}
