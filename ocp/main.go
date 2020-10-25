package main

import "fmt"

func main() {

	// sep up the demo products

	necklace := Product{
		name:  "necklace",
		color: silver,
		size:  medium,
	}
	apple := Product{
		name:  "Apple",
		color: green,
		size:  small,
	}
	orange := Product{
		name:  "Orange",
		color: orange,
		size:  small,
	}
	lenovo := Product{
		name:  "Lenovo",
		color: black,
		size:  medium,
	}
	macbook := Product{
		name:  "macbook",
		color: silver,
		size:  medium,
	}

	/* create a slice of the products for the principles demo. */
	products := []Product{apple, orange, necklace, lenovo, macbook}

	/* silver products found in the old way (breaks OCP) */
	fmt.Printf("Silver products (old): \n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, silver) {
		fmt.Printf(" - %s is silver \n", v.name)
	}

	/* silver products in the OCP compliant way with specification pattern */
	fmt.Printf("Silver products (new Spec pattern): \n")
	silverSpec := ColorSpecification{silver}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, silverSpec) {
		fmt.Printf(" - %s is silver \n", v.name)
	}

	/* demo the andSpecification */
	fmt.Printf("Medium Silver products (andSpecification example): \n")
	medSpec := SizeSpecification{medium}

	/* combines specs */
	lsSpec := AndSpecification{silverSpec, medSpec}
	for _, v := range bf.Filter(products, lsSpec) {
		fmt.Printf(" - %s is silver and medium sized \n", v.name)
	}

}
