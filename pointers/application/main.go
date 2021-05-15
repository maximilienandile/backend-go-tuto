package main

import "fmt"

type User struct {
	ID      string
	Name    string
	Email   string
	Blocked bool
}

func main() {
	daniel := User{
		ID:    "42",
		Name:  "Daniel",
		Email: "daniel@example.com",
	}
	blockUser(&daniel)
	fmt.Println(daniel)
	userPtr := &daniel
	fmt.Println(userPtr)
}

func blockUser(user *User) {
	user.Blocked = true
}
