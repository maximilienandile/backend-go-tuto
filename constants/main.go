package main

import "fmt"

const xlTshirtSizeLabel string = "XL"
const version string = "v1.0.0"

// default : bool
const isOpen = true

// default type is int
const maximumValue = 12

// default type is float64
const vatRat = 19.6

// default type is string
const shopName = "Gopher super store"

const overflowing = 9223372036854775808

func main() {
	fmt.Println(xlTshirtSizeLabel, version)

	const xlLabel = "XL"
	fmt.Println("Size selected:", xlLabel)
}
