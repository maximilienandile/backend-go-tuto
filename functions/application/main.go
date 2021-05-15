package main

import (
	"fmt"
)

func main() {
	fmt.Println(shippingPrice(9))
	fmt.Println(shippingPrice(12))
	fmt.Println(shippingPrice(100))

}

func shippingPrice(numberKg uint) int {
	if numberKg <= 10 {
		return 10
	} else if numberKg <= 20 {
		return 25
	} else {
		return 50
	}
}
