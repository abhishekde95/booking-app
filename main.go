package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTicktes = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTicktes, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets to attend")
}
