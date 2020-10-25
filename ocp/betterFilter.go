package main

/*
BetterFilter creates the type. It has no unique fields or
properties, instead it has a function open
to struct of this type.
*/
type BetterFilter struct {
}

/*
Filter takes a list of products and a criteria to use as the filter
returning a slice of pointers to the products fullfiling that
filter type.
*/
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}
