package main

type Color int

const (
	red Color = iota
	green
	orange
	black
	silver
)

type Size int

const (
	small Size = iota
	medium
	large
	xl
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

/*
FilterByColor takes a slice of the products, a colour and
returns a slice of pointers to products adhering to the filter
criteria
*/
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}
