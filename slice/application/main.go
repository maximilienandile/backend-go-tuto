package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 10_000)
	for k := 1; k <= 10_000; k++ {
		mySlice = append(mySlice, k)
	}
	fmt.Printf("Length is %d - Capacity is %d \n", len(mySlice), cap(mySlice))
	fmt.Println(mySlice[5])
}
