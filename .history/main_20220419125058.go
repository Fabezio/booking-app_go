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
fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
fmt.Printf("Il y a %v tickets et encore %v disponibles\n", conferenceName, remainingTickets)
fmt.Println("Obtenez vos tickets pour asister à l'événement")
fmt.Println("Montrer le nombre de tickets déjà distribués ? o/n")
	fmt.Scan(&entry)

	lookup = false
	// fmt.Println(conferenceName)
	if entry == "o" {
		lookup = true
	}
	if lookup {
		soldTickets := conferenceTickets - remainingTickets
		fmt.Printf("Tickets distribués jusqu'à maintenant: %v\n",  soldTickets)
	}

}