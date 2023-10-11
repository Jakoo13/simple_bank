// This file is to validate our incoming RPC requests

package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 3, 25); err != nil {
		return fmt.Errorf("username error: %w", err)
	}
	if !isValidUsername(username) {
		return fmt.Errorf("username must contain only letters, digits and underscores")
	}
	return nil
}

func ValidateFullName(fullname string) error {
	if err := ValidateString(fullname, 3, 25); err != nil {
		return fmt.Errorf("full name error: %w", err)
	}
	if !isValidFullName(fullname) {
		return fmt.Errorf("fullname must contain only letters or spaces")
	}
	return nil
}

func ValidatePassword(password string) error {
	return ValidateString(password, 6, 25)
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 5, 25); err != nil {
		return fmt.Errorf("email error: %w", err)
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}
