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
var bookings []string
var addBooking = true
var answer = ""

func getTickets() {
	var userName string
	var ticketsToDeal uint
	var email string
	if addBooking == true {
		fmt.Println("==========")
		fmt.Print("Entrez votre nom:\n  ")
		fmt.Scan(&userName)
		fmt.Printf("Combien serez-vous à participer ? (%v tickets disponibles)\n  ", remainingTickets)
		fmt.Scan(&ticketsToDeal)
		// fmt.Println("----------")
		fmt.Printf("Pour finir, quelle est votre adresse de contact?\n  ")
		fmt.Scan(&email)
		// fmt.Println("----------")
		fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", userName, email)
		remainingTickets = remainingTickets - ticketsToDeal
		bookings = append(bookings, userName)
		fmt.Println("----------")
		
			fmt.Printf("Il reste %v tickets. Faire une autre réservation? O/N ", remainingTickets)
		// if remainingTickets == 0  {
		// 	fmt.Println("Tickets épuisés")
		// } else {
		// 	// fmt.Printf("")
		// 	fmt.Scan(&answer)
		// 	if answer == "o" {
		// 		addBooking = true
		// 	} else if answer == "n" {
		// 		addBooking = false
		// 	} else {
		// 		fmt.Println("Veuillez taper 'o' ou 'n' ")
		// 	}
					// answer = ""
		// }
	}
}

func main () {
	
	fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
	fmt.Println("Obtenez vos tickets pour assister à l'événement")

	for remainingTickets > 0 {
		getTickets()
	} 
	fmt.Printf("dans l'ensemble: %v\n", bookings)
	fmt.Printf("En premier: %v\n", bookings[0])
	fmt.Printf("c'est mon type: %T\n", bookings)
	fmt.Printf("c'est la taille qui compte: %v\n", len(bookings))

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