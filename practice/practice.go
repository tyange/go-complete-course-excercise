package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

func main() {
	firstProduct := Product{"first-product", "A Book", "A nice Book", 10.99}
	secondProduct := NewProduct("second-product", "A Carpet", "A beautiful carpet", 99.99)

	firstProduct.printData()
	secondProduct.printData()
}
