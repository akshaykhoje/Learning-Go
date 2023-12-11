package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, reaminingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidUserTickets bool = userTickets > 0 && userTickets <= reaminingTickets

	return isValidName, isValidEmail, isValidUserTickets
}
