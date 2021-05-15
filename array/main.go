package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 42
	a[1] = 43
	fmt.Println(a)

	b := [2]string{"FR", "US"}
	fmt.Println(b)

	c := [...]float32{19.6, 20, 32}
	fmt.Println(c)

	fmt.Println(len(c))
	// access an element
	fmt.Println(c[1])

	for k, v := range c {
		fmt.Println(k, v)
	}
	fmt.Println("----")
	for k := 0; k < len(c); k++ {
		fmt.Println(k, c[k])
	}

	var m [2]uint8
	fmt.Println(m)
	// [0 0]
}
