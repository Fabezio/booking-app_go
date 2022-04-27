package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	// "strings"
	// "bufio"
)

const conferenceName string = "Go Conference"
const conferenceTokens uint = 50
var RemainingTokens = conferenceTokens
var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTokens uint
	// isUserOpted bool
}

func main () {

	greetUsers()

	for {
		firstName, lastName, email, tokensToDeal := getUserInput()

		isValidName, isValidEmail, isValidTokenNumber := helper.ValidateUserInput( firstName, lastName, email, tokensToDeal, RemainingTokens) 
		if isValidEmail && isValidName && isValidTokenNumber {
			bookToken( firstName , lastName, email, tokensToDeal)

			fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", firstName, email)
			fmt.Println("----------")
		
			firstNames := printFirstNames() 
			fmt.Printf("Participants: %v\n", firstNames)

			if RemainingTokens == 0 {
				fmt.Println("Jetons épuisés. Nous espérons vous voir lors de la prochaine conférence.")
				break
			}
			} else {
				fmt.Println("/!\\")
			if !isValidName {
				fmt.Println("Attention: votre nom ou prénom doit contenir au minimum deux caractères")
				// continue
			}
			if !isValidEmail {
				fmt.Println("Attention: votre adresse email doit contenir le symbole \"@\"")
				// continue
			}
			if !isValidTokenNumber {

				fmt.Println("Le nombre de jetons demandé est trop important. Nous espérons vous voir l'an prochain.")
			}
			// fmt.Println("Euh... Y'a un problème, là!")
			// fmt.Println("----------")
			// continue
		}
	} 


	fmt.Printf("Participants: %v\n", bookings)
}

func greetUsers() {
	fmt.Printf("Bienvenue sur notre application de réservation %v.\n",conferenceName)
	fmt.Printf("Il y a actuellement %v jetons disponibles.\n", conferenceTokens)
	fmt.Println("Obtenez vos jetons pour assister à la représentation.")
}

func printFirstNames() []string {
		firstNames := []string{}
			for _, booking := range bookings {
				// names := strings.Fields(booking)
				firstNames = append(firstNames, booking["firstName"])
			}
			return firstNames
			
}



func getUserInput() (string, string, string, uint) {
	var firstName string
		var lastName string
		var email string
		var tokensToDeal uint

		fmt.Println("==========")
		// getTextInput("entrez votre prénom", firstName)
		fmt.Print("Entrez votre prénom: ")
		fmt.Scan(&firstName)
		fmt.Print("Entrez votre nom de famille:  ")
		fmt.Scan(&lastName)
		fmt.Printf("Quelle est votre adresse de contact?  ")
		fmt.Scan(&email)
		fmt.Printf("Enfin, combien serez-vous à participer ? (%v jetons disponibles)  ", RemainingTokens)
		fmt.Scan(&tokensToDeal)

		return firstName, lastName, email, tokensToDeal
}
// func getTextInput(msg , input string) {
// 	fmt.Println(msg)
// 		fmt.Scan(&input)

// }

func bookToken(firstName string, lastName string, email string, tokensToDeal uint) {
	RemainingTokens = RemainingTokens - tokensToDeal
	// var mySlice [string]
	// var myMap = map[string]string
	var userData = make(map[string]string)
	userData["firstName"]=firstName
	userData["lastName"]=lastName
	userData["email"]=email
	userData["firstName"]=firstName
	userData["tokensToDeal"]=strconv.FormatUint(uint64(tokensToDeal), 10)
	// var fullName = firstName + " " + lastName
	bookings = append(bookings, userData)
	// return bookings, RemainingTokens
	// fmt.Printf("Réservations: %v", bookings)
}