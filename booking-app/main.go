package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const totalAmountOfConferenceTickets = 50

var availableTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname string 
	lastname string
	email string
	numberOfTickets uint
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are stil available\n", totalAmountOfConferenceTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstname)
	}

	fmt.Printf("All current bookings details: %v\n", firstNames)
}

func handleUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)
	return firstname, lastname, email, userTickets
}

func handleBooking( userTickets uint, firstname string, lastname string, email string) {
	availableTickets = availableTickets - userTickets

	var userData = UserData {
		firstname,
		lastname,
		email,
		userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)

	fmt.Printf("There are %v remaining tickets for %v\n", availableTickets, conferenceName)

}

func sendTicket(userTickets uint, firstname string, lastname string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastname)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

var wg = sync.WaitGroup{}


func main() {
	greetUsers()

	for availableTickets > 0 && len(bookings) < 50 {

		firstname, lastname, email, userTickets := handleUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstname, lastname, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			handleBooking( userTickets,  firstname, lastname, email)

			wg.Add(1)
			go sendTicket(userTickets, firstname, lastname, email)

			printFirstNames()

			if availableTickets == 0 {
				fmt.Println("All tickets for this events have been sold! Sorry")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The firstname or lastname you entered is invalid")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("The amount of tickets you inputed is invalid or exceeds the amount of tickets we have")
			}
			continue
		}

	}
	wg.Wait()
}
