package main

import (
	"fmt"
	"strings"
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

// func getTickets() {
// 	var firstName string
// 	var lastName string
// 	var jetonsToDeal uint
// 	var email string
// 	// if addBooking == true {
// 		fmt.Println("==========")
// 		fmt.Print("Entrez votre prénom:\n  ")
// 		fmt.Scan(&firstName)
// 		fmt.Print("Entrez votre nom de famille:\n  ")
// 		fmt.Scan(&lastName)
// 		fmt.Printf("Combien serez-vous à participer ? (%v jetons disponibles)\n  ", remainingTickets)
// 		fmt.Scan(&jetonsToDeal)
// 		// fmt.Println("----------")
// 		fmt.Printf("Pour finir, quelle est votre adresse de contact?\n  ")
// 		fmt.Scan(&email)
// 		// fmt.Println("----------")
// 		fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", firstName, email)
// 		remainingTickets = remainingTickets - jetonsToDeal
// 		var fullName = firstName + " " + lastName
// 		bookings = append(bookings, fullName)
// 		fmt.Println("----------")
		
// 			fmt.Printf("Il reste %v jetons.\n", remainingTickets)
			
// 		//  else
// 		//  {
// 		// 	// fmt.Printf("")
// 		// 	fmt.Scan(&answer)
// 		// 	if answer == "o" {
// 		// 		addBooking = true
// 		// 	} else if answer == "n" {
// 		// 		addBooking = false
// 		// 	} else {
// 		// 		fmt.Println("Veuillez taper 'o' ou 'n' ")
// 		// 	}
// 					// answer = ""
// 		// }
// 	// }
// }

func main () {
	
	fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
	fmt.Println("Obtenez vos jetons pour assister à l'événement")
  // bookedOut := remainingTickets == 0

		
	for remainingTickets > 0 {
		
		var firstName string
	var lastName string
	var jetonsToDeal uint
	var email string
	// if addBooking == true {
		fmt.Println("==========")
		fmt.Print("Entrez votre prénom:\n  ")
		fmt.Scan(&firstName)
		fmt.Print("Entrez votre nom de famille:\n  ")
		fmt.Scan(&lastName)
		fmt.Printf("Combien serez-vous à participer ? (%v jetons disponibles)\n  ", remainingTickets)
		fmt.Scan(&jetonsToDeal)
		// fmt.Println("----------")
		fmt.Printf("Pour finir, quelle est votre adresse de contact?\n  ")
		fmt.Scan(&email)
		// fmt.Println("----------")
		fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", firstName, email)
		remainingTickets = remainingTickets - jetonsToDeal
		var fullName = firstName + " " + lastName
		bookings = append(bookings, fullName)
		fmt.Println("----------")
		
			// fmt.Printf("Il reste %v jetons.\n", remainingTickets)
			
			
			if remainingTickets == 0  {
				fmt.Println("Jetons épuisés. Nous espérons vous voir à la prochaine conférence")
				// break
			}
			} 
			firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		// firstName := names[0]
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("dans l'ensemble: %v\n", firstNames)

	// fmt.Printf("En premier: %v\n", bookings[0])
	// fmt.Printf("c'est mon type: %T\n", bookings)
	// fmt.Printf("c'est la taille qui compte: %v\n", len(bookings))

	// fmt.Println("Montrer le nombre de jetons déjà distribués ? o/n")
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