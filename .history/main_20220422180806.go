package main

import (
	"fmt"
	"strings"
	// "bufio"
)

const conferenceName string = "Go Conference"
var conferenceTickets uint = 50
var remainingTokens = conferenceTickets
// var userTokens
var bookings []string
// var addBooking = true
// var answer = ""

func main () {
	fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
	fmt.Println("Obtenez vos jetons pour assister à l'événement")
	for remainingTokens > 0 {
		var firstName string
		var lastName string
		var tokensToDeal uint
		var email string

		fmt.Println("==========")

		fmt.Print("Entrez votre prénom:\n  ")
		fmt.Scan(&firstName)
		fmt.Print("Entrez votre nom de famille:\n  ")
		fmt.Scan(&lastName)
		fmt.Printf("Combien serez-vous à participer ? (%v jetons disponibles)\n  ", remainingTokens)
		fmt.Scan(&tokensToDeal)
		fmt.Printf("Pour finir, quelle est votre adresse de contact?\n  ")
		fmt.Scan(&email)

		fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", firstName, email)

		remainingTokens = remainingTokens - tokensToDeal
		if tokensToDeal > remainingTokens {
			fmt.Printf("Le nombre de jetons demandé est trop important. Désolé.")
		}
		var fullName = firstName + " " + lastName
		bookings = append(bookings, fullName)
		fmt.Println("----------")
	
		if remainingTokens == 0  {
			fmt.Println("Jetons épuisés. Nous espérons vous voir à la prochaine conférence")
		}
	} 

	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("Participants: %v\n", firstNames)

}