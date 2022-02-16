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

	fmt.Printf("Welcome to Daddy Awesome's %v Booking App\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here to attend\n\n")

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTicket = remainingTicket - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank You, %v %v for booking %v tickets \n", firstName, lastName, userTickets)
			fmt.Printf("We will send a confirmation email at this adress %v\n", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("Thefirst names of our bookings are: %v\n\n", firstNames)

			if remainingTicket == 0 {
				// end program
				fmt.Printf("Our Conference is booked out. Come back next year")
				break
			}

		} else {
			fmt.Println("Your Input Data is Invalid, Please Try Again!")
		}

	}

}
