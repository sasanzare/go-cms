package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail checks if an email address has a valid format.
//
// Parameters:
//   - email: The email address to validate (string)
//
// Returns:
//   - bool: true if email is valid, false otherwise
//
// Note:
//   - Uses a standard email regex pattern
//   - Doesn't verify if the email actually exists
//
// Example:
//   isValid := ValidateEmail("user@example.com") // returns true
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,63}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// ValidatePassword checks if a password meets strength requirements.
//
// Requirements:
//   - Minimum 8 characters
//   - At least 1 letter
//   - At least 1 number
//   - At least 1 uppercase letter
//   - At least 1 special character (!@#$%^&*)
//
// Parameters:
//   - password: The password to validate (string)
//
// Returns:
//   - bool: true if password meets all requirements, false otherwise
//
// Example:
//   isStrong := ValidatePassword("Passw0rd!") // returns true
func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
	return hasLetter && hasNumber && hasUpper && hasSpecial
}

// SanitizeInput cleans and sanitizes user input by:
//   - Trimming whitespace
//   - Removing potentially harmful characters (<>"'&()/)
//
// Parameters:
//   - input: The string to sanitize (string)
//
// Returns:
//   - string: The sanitized version of the input
//
// Security Note:
//   - This provides basic protection against XSS but shouldn't be
//     the only security measure
//
// Example:
//   clean := SanitizeInput("  <script>alert()</script>  ") // returns "scriptalertscript"
func SanitizeInput(input string) string {
    cleaned := strings.TrimSpace(input)
    cleaned = regexp.MustCompile(`[<>"'&()/]`).ReplaceAllString(cleaned, "")
    return cleaned
}