package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ageJohn := rand.Intn(100)
	fmt.Println("John is", ageJohn, "years old")
	for i := 0; i < ageJohn; i++ {
		fmt.Println("iteration N°", i)
	}

	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter > 100 {
			break
		}
	}

}
