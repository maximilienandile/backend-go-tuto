package main

import "fmt"

type Product struct {
	ID        string
	Reference string
	Title     string
	Category
}

type Category struct {
	ID   string
	Name string
}

func main() {

	cat := Category{
		ID:   "catId",
		Name: "My category",
	}
	products := []Product{
		{
			ID:        "42",
			Reference: "myRef",
			Title:     "My first Product",
			Category:  cat,
		},
		{
			ID:        "43",
			Reference: "myRef",
			Title:     "My second Product",
			Category:  cat,
		},
		{
			ID:        "44",
			Reference: "myRef",
			Title:     "My third Product",
			Category:  cat,
		},
	}
	for _, product := range products {
		fmt.Printf("Product of id %s belongs to category of id %s (name of the cat: %s) \n", product.ID, product.Category.ID, product.Category.Name)
	}

}
