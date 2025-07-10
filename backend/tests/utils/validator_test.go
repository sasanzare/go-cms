package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sasanzare/go-cms/utils"
)

// TestValidateEmail tests the ValidateEmail function with various email formats.
// It verifies both valid and invalid email patterns to ensure correct validation.
//
// Test Cases:
//   1. Standard email format (test@example.com)
//   2. Email with plus addressing (user+tag@example.com)
//   3. Invalid email (missing @)
//   4. Invalid email (missing domain extension)
//   5. Empty string
//   6. Email with subdomain (user@sub.domain.com)
//
// Note:
//   - Tests cover common edge cases
//   - Each test case includes the email as debug message
func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},         // Valid standard email
		{"user+tag@example.com", true},     // Valid email with plus addressing
		{"invalid.email", false},           // Invalid email (missing @)
		{"another@test", false},            // Invalid email (missing domain extension)
		{"", false},                       // Empty string should be invalid
		{"user@sub.domain.com", true},      // Valid email with subdomain
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, utils.ValidateEmail(tc.email), tc.email)
	}
}

// TestValidatePassword tests the ValidatePassword function with various password scenarios.
// It verifies the password strength requirements including length, character types,
// and special characters.
//
// Test Cases:
//   1. Valid strong password (meets all requirements)
//   2. Too short password
//   3. Password missing letters
//   4. Password missing numbers
//   5. Password missing special characters
//   6. Empty password
//
// Security Note:
//   - Tests minimum security requirements
//   - Each test case includes the password as debug message (avoid sensitive data)
func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"StrongPass123!", true},       // Valid strong password
		{"weak", false},               // Too short (less than 8 characters)
		{"12345678", false},           // Missing letters
		{"abcdefgh", false},           // Missing numbers
		{"ABCDEFGH", false},           // Missing numbers and special characters
		{"Pass123!", true},            // Valid password
		{"", false},                   // Empty password should be invalid
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, utils.ValidatePassword(tc.password), tc.password)
	}
}

// TestSanitizeInput tests the SanitizeInput function with various input scenarios.
// It verifies proper whitespace trimming and removal of potentially harmful characters.
//
// Test Cases:
//   1. Input with surrounding whitespace
//   2. Clean input (no changes needed)
//   3. Whitespace-only input
//   4. Input with HTML tags
//   5. Input with special HTML characters
//   6. Valid email input
//
// Security Note:
//   - Tests basic XSS protection
//   - Verifies preservation of safe content
func TestSanitizeInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"  clean  ", "clean"},                     // Should trim whitespace
		{"no spaces", "no spaces"},                 // No change needed
		{"   ", ""},                                // Whitespace-only should become empty
		{"<script>alert('xss')</script>", "scriptalertxssscript"}, // Should remove HTML tags
		{"user@domain.com", "user@domain.com"},     // Valid email should remain unchanged
		{"a&b<c>d", "abcd"},                       // Should remove special HTML characters
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, utils.SanitizeInput(tc.input), tc.input)
	}
}