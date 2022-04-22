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
var name string
var ticketsToDeal uint
var email string

func ask (val any, msg string) any{
	fmt.Println(msg)
	fmt.Scan(&val)
	return val
}
func getEmailAndConfirm(){
	ask(email, "Pour finir, quelle est votre adresse de contact?")
	// fmt.Println("Pour finir, quelle est votre adresse de contact?")
	// fmt.Scan(&email)
	fmt.Printf("Merci, %v. Vous serez notifié sous peu par courriel à l'adresse %v. Passez une excellente journée.\n", name, email)
	remainingTickets = remainingTickets - ticketsToDeal
	
	fmt.Printf("Il y a %v tickets et encore %v disponibles\n", conferenceTickets, remainingTickets)

}

func getTickets() {
	ask(name, "entrez votre nom")
	// fmt.Scan(&name)
	// 	fmt.Println(&val)
	// 	return val
	fmt.Printf("Enchanté, %v.\n", name)
	fmt.Printf("Il reste %v tickets.\n", remainingTickets)

	ask(ticketsToDeal, "Combien serez-vous à participer ?")
	// fmt.Scan(&ticketsToDeal)
	getEmailAndConfirm()
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
getTickets()
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