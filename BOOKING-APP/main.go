package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTicket int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			fmt.Printf("Thank You, %v %v for booking %v tickets \n", firstName, lastName, userTickets)
			fmt.Printf("We will send a confirmation email at this adress %v\n", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames()

			fmt.Printf("The first names of our bookings are: %v\n\n", firstNames)

			if remainingTickets == 0 {
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

func greetUsers() {
	fmt.Printf("Welcome to Daddy Awesome's %v Booking App\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n\n")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
