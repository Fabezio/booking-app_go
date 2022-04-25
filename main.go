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
	for remainingTokens > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var tokensToDeal uint
		var email string

		fmt.Println("==========")

		fmt.Print("Entrez votre prénom:\n  ")
		fmt.Scan(&firstName)
		fmt.Print("Entrez votre nom de famille:\n  ")
		fmt.Scan(&lastName)
		fmt.Printf("Quelle est votre adresse de contact?\n  ")
		fmt.Scan(&email)
		fmt.Printf("Enfin, combien serez-vous à participer ? (%v jetons disponibles)\n  ", remainingTokens)
		fmt.Scan(&tokensToDeal)
		
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		if !isValidName {
			fmt.Println("Attention: votre nom ou prénom doit contenir au minimum deux caractères")
			continue
		}
		if !isValidEmail {
			fmt.Println("Attention: votre adresse email doit contenir le symbole \"@\"")
			continue
		}

		if tokensToDeal <= remainingTokens {
			remainingTokens = remainingTokens - tokensToDeal
			var fullName = firstName + " " + lastName
			bookings = append(bookings, fullName)

		fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", firstName, email)
			fmt.Println("----------")
		
			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Participants: %v\n", firstNames)
			if remainingTokens == 0 {
				fmt.Println("Jetons épuisés. Nous espérons vous voir lors de la prochaine conférence.")
				break
			}
		} else {
			fmt.Println("Le nombre de jetons demandé est trop important. Nous espérons vous voir l'an prochain.")
			fmt.Println("----------")
			// continue
		}


		
	} 
}