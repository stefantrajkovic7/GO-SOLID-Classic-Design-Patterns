package main

import "fmt"

type Color int
type Size int

const (
	red Color = iota
	green
	blue
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name string
	color Color
	size Size
}

type Filter struct {
	// ...some settings
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products{
		if v.color == color{
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{ apple, tree, house }
	fmt.Printf("Green products (old:\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
}