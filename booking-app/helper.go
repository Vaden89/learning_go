package main

import "strings"

func validateUserInput(firstname string, lastname string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= availableTickets

	return isValidName, isValidEmail, isValidTicketNumber
}