package main

import "fmt"

func main() {
	const emailToSend = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("sending email...")
		emailSent++
		fmt.Println("value of emailSent:", emailSent)
		// emailSent = emailSent + 1
	}

}
