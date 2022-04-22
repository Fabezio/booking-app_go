package main

import( 
	"fmt"
	// "bufio"


)

// func scanner (val void, msg string) void {
// 	fmt.Println(msg)
// 	fmt.Scan(&val)
// 	fmt.Println(&val)
// 	return val
// }

func main () {
	 conferenceName := "Go Conference"
	conferenceTickets := 50
 remainingTickets := 50
remainingTickets--
lookup := false
var name string
var ticketsToDeal number
var entry string
fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
fmt.Printf("Il y a %v tickets et encore %v disponibles\n", conferenceName, remainingTickets)
fmt.Println("Obtenez vos tickets pour assister à l'événement")
fmt.Println("entrez votre nom")
	fmt.Scan(&name)
// 	fmt.Println(&val)
// 	return val
fmt.Printf("Bonjour, %v. Combien serez-vous à assister à la conférence?\n", name)
fmt.Scan(&ticketsToDeal)
fmt.Println("Montrer le nombre de tickets déjà distribués ? o/n")
	fmt.Scan(&entry)

	lookup = false
	// fmt.Println(conferenceName)
	if entry == "o" {
		lookup = true
	} else if entry == "n" {
		fmt.Println("vous ne verrez pas cette option")
	}
	if lookup {
		soldTickets := conferenceTickets - remainingTickets
		fmt.Printf("Tickets distribués jusqu'à maintenant: %v\n",  soldTickets)
	}

}