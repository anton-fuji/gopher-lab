package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	sc := bufio.NewReader(os.Stdin)
	fmt.Fscan(sc, &s)

	fmt.Println(s)

}
