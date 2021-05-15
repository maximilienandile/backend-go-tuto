package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(computePrice(10, 10))
	fmt.Println(computePrice(1, 10))
	fmt.Println(computePrice(2, 10.50))

	price, formatted := computePriceV2(15, 124.2)
	fmt.Println(price, formatted)
	printHello()
	printHello()
	printHello()
	printHello()
}

func computePrice(numberOfNights uint, pricePerNight float32) float32 {
	return float32(numberOfNights) * pricePerNight
}
func computePriceV2(numberOfNights uint, pricePerNight float32) (float32, string) {
	totalPrice := float32(numberOfNights) * pricePerNight
	formatted := fmt.Sprintf("%.2f $", totalPrice)
	return totalPrice, formatted
}

func printHello() {
	fmt.Println("Hello")
}
