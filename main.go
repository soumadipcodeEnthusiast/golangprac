package main

import (
	"fmt"
	"strconv"
	"strings"
)

var conferenceName = "GO Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

var bookings = make([]map[string]string, 0)

//map are used to store key and vlues according to the keys making easy for the user the see the particular thing

func main() {
	greetUser(conferenceName, conferenceTickets, remainingTickets)
	for {
		firstName, lastName, email, userTickets := getUserinput()
		isvalidemail, isvalidname, isvalidnumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isvalidemail && isvalidname && isvalidnumber {

			bookTickets(remainingTickets, bookings, firstName, lastName, email, userTickets, conferenceName)
			firstNames := getFirstNames(bookings)

			fmt.Printf("these are all tghe booking: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("out conference tickets are booked out.come back next year.")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Println("evil men you made a boy cry")
		} else {
			if !isvalidemail {
				fmt.Println("email address you entered is invalid it should contain @")

			}
			if !isvalidname {
				fmt.Println("your name is too short")
			}
			if !isvalidnumber {
				fmt.Println("either the ticket number you entered is negetive or more than remaining tickets")
			}
		}
	}
	/*city := "London"
	 switch city {
	 case "New York":
		//there will be diffrent logic
	 case "Singapore":
		//there will be diffrent logic
	 case "London":
		//there will be diffrent logic
	 case "Berlin","Tokyo":
		//there will be diffrent logic
	 default:
		fmt.Println("no valid city selected")
	 }*/
}
func greetUser(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("we have totale %v tickets and %v are remaing\n", confTickets, remainingTickets)
}
func getFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)

		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserinput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isvalidname := len(firstName) >= 2 && len(lastName) >= 2
	isvalidemail := strings.Contains(email, "@")
	isvalidnumber := userTickets > 0 && userTickets <= remainingTickets
	return isvalidemail, isvalidname, isvalidnumber
}
func bookTickets(remainingTickets uint, bookings []string, firstName string, lastName string, email string, userTickets uint, conferenceName string) {

	remainingTickets = remainingTickets - userTickets
	//create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("Thankyou %v %v for booking %v tickets. You with recieve a confirmation email at %v\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
