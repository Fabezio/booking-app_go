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
	var conferenceTickets uint = 50
 var remainingTickets uint = 50
// remainingTickets--
// lookup := false
/* 
   Données personnelles à entrer
 */
var name string
var ticketsToDeal uint
var email string
var user [] = 
// var entry string
fmt.Printf("Bienvenue sur notre application de réservation %v\n",conferenceName)
fmt.Println("Obtenez vos tickets pour assister à l'événement")
fmt.Println("entrez votre nom")
fmt.Scan(&name)
// 	fmt.Println(&val)
// 	return val
fmt.Printf("Bonjour, %v. Combien serez-vous à assister à la conférence?\n", name)
fmt.Scan(&ticketsToDeal)
fmt.Println("Pour finir, quelle est votre adresse de contact?")
fmt.Scan(&email)
fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel au %v. Passez une excellente journée.\n", name, email)
remainingTickets = remainingTickets - ticketsToDeal

fmt.Printf("Il y a %v tickets et encore %v disponibles\n", conferenceTickets, remainingTickets)
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