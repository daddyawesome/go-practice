package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTicket int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

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
		//wg.Wait()
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
		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending Tickets:\n %v \n to email address %v \n", ticket, email)
	fmt.Println("######################")
	wg.Done()
}
