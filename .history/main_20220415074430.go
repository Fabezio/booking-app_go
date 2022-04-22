package main

import "fmt"

func main () {
	var conferenceName = "Go Conference"
const 	conferenceTickets = 50
var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and remainingTickets %v ones still available\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fmt.Println(conferenceName)

}