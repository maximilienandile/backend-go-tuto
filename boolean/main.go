package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rooms := 100
	rand.Seed(time.Now().Unix())
	roomsOccupied := rand.Intn(100)
	fmt.Println("rooms:", rooms, "rooms Occupied:", roomsOccupied)

	fmt.Println("Do we have some rooms left ?", rooms > roomsOccupied)
	fmt.Println("Is the hotel full ?", rooms == roomsOccupied)

}
