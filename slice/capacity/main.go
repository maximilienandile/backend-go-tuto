package main

import "fmt"

func main() {
	s := []uint{10, 20, 30, 40}
	fmt.Println("length:", len(s), "capacity:", cap(s))
	s = append(s, 50)
	fmt.Println("length:", len(s), "capacity:", cap(s))
	s = append(s, 60)
	fmt.Println("length:", len(s), "capacity:", cap(s))
	s = append(s, 70)
	fmt.Println("length:", len(s), "capacity:", cap(s))
	s = append(s, 80)
	fmt.Println("length:", len(s), "capacity:", cap(s))
	s = append(s, 90)
	fmt.Println("length:", len(s), "capacity:", cap(s))

	b := make([]int, 0, 10)
	fmt.Println(b)
	fmt.Println("capacity", cap(b), "length", len(b))
	b = append(b, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("capacity", cap(b), "length", len(b))
	fmt.Println(b)
	b = append(b, 12)
	fmt.Println("capacity", cap(b), "length", len(b))

}
