package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	//email validation (Needs to enter a valid email format, containng "@" sign)
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, "mail") && strings.Contains(email, ".com")

	//User Tickets Validation (Needs to enter correct number of tickets (Positive + greater than 0))
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}
