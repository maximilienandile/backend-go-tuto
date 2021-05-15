package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	agePaul := rand.Intn(110)
	ageJohn := rand.Intn(110)
	fmt.Println("Paul", agePaul, "John", ageJohn)
	//if ageJohn == agePaul {
	//	// ageJohn is equal to agePaul
	//	fmt.Println("Paul and John have the same age")
	//} else if ageJohn > agePaul {
	//	fmt.Println("John is older than Paul")
	//} else {
	//	fmt.Println("Paul and John have not the same age")
	//	fmt.Println("Paul is older than John")
	//}
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("default section")
	}

	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("John age + Paul age is equal to 10")
	case 20, 30, 40:
		fmt.Println("Age of John + Paul is equal to 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("Age of John + Paul = 2 times the age of Paul")
	}

	switch {
	case agePaul > ageJohn:
		fmt.Println("Paul is older than John")
	case agePaul < ageJohn:
		fmt.Println("John is older than Paul")
	case agePaul == ageJohn:
		fmt.Println("John and Paul have the same age")
	}

	fmt.Println("End of the program")
}
