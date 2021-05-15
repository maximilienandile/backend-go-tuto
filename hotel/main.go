package main

import (
	"fmt"
	"math/rand"
	"time"
)

const hotelName = "Gopher Paris Inn"
const totalRooms = 134

func main() {
	rand.Seed(time.Now().Unix())
	roomsRented := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsRented
	fmt.Println("Hotel: ", hotelName)
	fmt.Printf("Number of rooms: %d\n", totalRooms)
	fmt.Printf("Rooms available: %d\n", roomsAvailable)
	var occupancyRate float32 = float32(roomsRented) / float32(totalRooms) * 100
	fmt.Printf("						Occupancy Rate: %0.0f %% \n", occupancyRate)
	var occupancyLevel string
	if occupancyRate < 30 {
		// low
		occupancyLevel = "Low"
	} else if occupancyRate < 60 {
		// medium
		occupancyLevel = "Medium"
	} else {
		// high
		occupancyLevel = "High"
	}
	fmt.Println("						Occupancy Level:", occupancyLevel)

	if roomsAvailable == 0 {
		fmt.Println("No rooms available for tonight")
	} else {
		fmt.Println("Rooms :")
		for i := 0; i < roomsAvailable; i++ {
			roomNumber := 110 + i
			nbNights := rand.Intn(10) + 1
			nbPeople := rand.Intn(10) + 1
			if nbNights > 1 {
				fmt.Printf("- %d : %d people / %d nights \n", roomNumber, nbPeople, nbNights)
			} else {
				fmt.Printf("- %d : %d people / %d night \n", roomNumber, nbPeople, nbNights)
			}
		}
	}
}
