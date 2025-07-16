package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p Product) discount(percentage int) Product {
	disPrice := p.Price * (100 - percentage) / 100
	return Product{
		Name:  p.Name,
		Price: disPrice,
	}
}

func main() {
	p1 := Product{Name: "Goハンドブック", Price: 3400}
	p2 := p1.discount(30)
	fmt.Println(p1.Name, p1.Price)
	fmt.Println(p2.Name, p2.Price)
}
