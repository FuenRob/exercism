package techpalace

import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    response := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return response
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    str_start := ""
    for i := 0; i < numStarsPerLine; i += 1 {
		str_start += "*"
	}
    
    str := str_start + "\n" + welcomeMsg + "\n" + str_start

    return str
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	str := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(str)
}
