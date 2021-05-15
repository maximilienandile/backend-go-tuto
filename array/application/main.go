package main

import "fmt"

func main() {
	fibo := [5]uint8{1, 2, 3, 5, 8}
	fmt.Println(fibo)
	for k, _ := range fibo {
		fmt.Println(k)
	}
}
