package main

import "fmt"

type User struct {
	Firstname string
	Lastname  string
}

func main() {
	var p *int
	fmt.Println(p)

	age := 42
	p = &age
	fmt.Println(p)

	john := User{
		Firstname: "John",
		Lastname:  "Doe",
	}
	fmt.Println(john)
	pu := &john
	fmt.Println(pu)

	fmt.Println(p)
	fmt.Println(*p)

}
