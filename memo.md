```golang
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	d := a / b
	r := a % b
	f := float64(a) / float64(b)

	fmt.Printf("%d %d %.5f\n", d, r, f)
}
```