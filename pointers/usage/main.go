package main

import "fmt"

type Cart struct {
	NbItems uint
}

func main() {
	val := 42
	increment(&val)
	fmt.Println(val)

	newCart := Cart{NbItems: 10}
	addItem(&newCart)
	fmt.Println(newCart)
}

func addItem(cart *Cart) {
	cart.NbItems++
}

func increment(value *int) {
	*value = *value + 1
}
