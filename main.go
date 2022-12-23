package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go conference"
	const conferenceTicktes int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", conferenceTicktes, remainingTickets)
	fmt.Println("Get your tickets to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)

}
