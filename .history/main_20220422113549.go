package main

import (
	"fmt"
	// "bufio"
)

// func scanner (val void, msg string) void {
// 	fmt.Println(msg)
// 	fmt.Scan(&val)
// 	fmt.Println(&val)
// 	return val
// }
const conferenceName string = "Go Conference"
var conferenceTickets uint = 50
// var remainingTickets uint = 1
var remainingTickets = conferenceTickets
var userName string
var ticketsToDeal uint
var email string
var bookings []string



func getTickets() {
	fmt.Println("==========")
	fmt.Println("Entrez votre nom")
	fmt.Scan(&userName)
	// 	fmt.Println(&val)
	// 	return val
	fmt.Printf("Enchanté, %v.\n", userName)
	fmt.Println("----------")
	fmt.Printf("Il reste %v tickets. Combien serez-vous à participer ?\n", remainingTickets)
	fmt.Scan(&ticketsToDeal)
	fmt.Println("----------")
	fmt.Println("Pour finir, quelle est votre adresse de contact?")
	fmt.Scan(&email)
	fmt.Println("----------")
	fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", userName, email)
	remainingTickets = remainingTickets - ticketsToDeal
	bookings[0] = userName
	if remainingTickets == 0  {
		fmt.Println("Tickets épuisés")

	} else {
		fmt.Printf("Il y a %v tickets disponibles\n", remainingTickets)

	}
	
	
}

func main () {
// remainingTickets--
// lookup := false
/* 
Données personnelles à entrer
*/
// var tickets = [50]
// var entry string
fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
fmt.Println("Obtenez vos tickets pour assister à l'événement")
// for (q:=remainingTickets; q=0; q=ticketsToDeal)
for remainingTickets>0 {
	getTickets()

}
fmt.Printf("dans l'ensemble: %v\n", bookings)
fmt.Printf("En premier: %v\n", bookings[0])
fmt.Printf("c'est mon type: %T\n", bookings)
fmt.Printf("c'est la taille qui compte: %T\n", len(bookings))

// fmt.Println("Montrer le nombre de tickets déjà distribués ? o/n")
// fmt.Scan(&entry)

// lookup = false
// fmt.Println(conferenceName)
// if entry == "o" {
	// 	lookup = true
	// } else if entry == "n" {
	// 	fmt.Println("vous ne verrez pas cette option")
	// }
	// if lookup {
	// 	var soldTickets uint = conferenceTickets - remainingTickets
	// 	fmt.Printf("Tickets distribués jusqu'à maintenant: %v\n",  soldTickets)
	// }

}