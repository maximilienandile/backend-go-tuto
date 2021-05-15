package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c"}
	fmt.Println(mySlice)

	b := make([]uint8, 2)
	fmt.Println(b)
	b[1] = 42
	b[0] = 32
	fmt.Println(b)
	if b[1] == 100 {
		fmt.Println("element at index 1 is equal to 42")
	}
	//b[2] = 45
	b = append(b, 45, 54)
	fmt.Println(b)

	fmt.Println(len(b))
}
