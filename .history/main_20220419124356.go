package main

import( 
	"fmt"
	// "bufio"


)

func main () {
	 conferenceName := "Go Conference"
	conferenceTickets := 50
 remainingTickets := 50
remainingTickets--
lookup := false
var entry string
scanner := fmt.Scan(&entry)

	fmt.Printf("Welcome to our %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and remainingTickets %v ones still available\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	lookup = false
	// fmt.Println(conferenceName)
	if entry == "y" {
		lookup = true
	}
	if lookup {
		soldTickets := conferenceTickets - remainingTickets
		fmt.Printf("Delivered tickets: %v\n",  soldTickets)
	}

}