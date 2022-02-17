package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTicket, remainingTicket)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("What is your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("What is your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("What is your email address: ")
		fmt.Scan(&email)

		fmt.Printf("What is the number of tickets you would like to buy: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validUserInput(firstName, lastName, email, userTickets, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTicket = remainingTicket - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank You, %v %v for booking %v tickets \n", firstName, lastName, userTickets)
			fmt.Printf("We will send a confirmation email at this adress %v\n", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

			firstNames := getFirstNames(bookings)

			fmt.Printf("The first names of our bookings are: %v\n\n", firstNames)

			if remainingTicket == 0 {
				// end program
				fmt.Printf("Our Conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered needs an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The numbver of ticket you've entered is invalid")
			}

		}

	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to Daddy Awesome's %v Booking App\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n\n")

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validUserInput(firstName string, lastName string, email string, userTickets uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
