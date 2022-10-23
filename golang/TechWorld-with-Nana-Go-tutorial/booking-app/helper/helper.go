package helper

import "strings"

func ValidateUserInput(fName string, lName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// user input validation
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	// isValidCity := city == "Singapore" || city == "London"

	return isValidName, isValidEmail, isValidTickets
}
