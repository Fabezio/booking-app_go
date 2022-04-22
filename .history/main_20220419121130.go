package main

import "fmt"

func main () {
	 conferenceName := "Go Conference"
	conferenceTickets := 50
 remainingTickets := 50
remainingTickets--
// var lookup = false

	fmt.Printf("Welcome to our %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and remainingTickets %v ones still available\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fmt.Println(conferenceName)
	// if lookup {
		soldTickets := conferenceTickets - remainingTickets
		fmt.Printf("Sold tickets: %v.\n",  soldTickets)
	// }

}