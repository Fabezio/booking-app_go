package helper

import (
	"strings"
)




func ValidateUserInput(firstName, lastName, email string, tokensToDeal uint, remainingTokens uint) (bool, bool, bool) {
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTokenNumber := tokensToDeal > 0 && tokensToDeal <= remainingTokens
		return isValidName, isValidEmail, isValidTokenNumber
}