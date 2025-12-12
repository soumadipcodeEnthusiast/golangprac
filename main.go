package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "GO Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("the type of conferenceName %T,type of conferenceTickets %T,type of remaining tickets %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("welcome tox", conferenceName, "booking app")
	fmt.Println("we hav this many tickets", 50, "this many ticket are remaining", remainingTickets)
	fmt.Println("you can book your tickets here")
	var bookings []string
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("enter your firstname")
		fmt.Scan(&firstName)
		fmt.Println("enter your lastname")
		fmt.Scan(&lastName)
		fmt.Println("enter your email")
		fmt.Scan(&email)
		fmt.Println("entr no of tickets to buy")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)

			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Thankyou %v %v for booking %v tickets. You with recieve a confirmation email at %v\n", firstName, lastName, email, userTickets)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Printf("these are all tghe booking: %v\n", firstNames)
	}

}
