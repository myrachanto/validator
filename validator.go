package validator

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidName     = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidNumber   = regexp.MustCompile(`^[0-9]+$`).MatchString
)

func ValidateString(str string, min, max int) error {
	n := len(str)
	if n < min || n > max {
		return fmt.Errorf("must contain characters between %d - %d ", min, max)
	}
	return nil
}
func ValidateUserName(name string) error {
	if err := ValidateString(name, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(name) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}
func ValidatePhone(num string) error {
	if err := ValidateString(num, 5, 15); err != nil {
		return err
	}
	if !isValidNumber(num) {
		return fmt.Errorf("must contain only digits")
	}
	return nil
}
func ValidateName(name string) error {
	if err := ValidateString(name, 3, 100); err != nil {
		return err
	}
	if !isValidName(name) {
		return fmt.Errorf("must contain only letters, spaces")
	}
	return nil
}
func ValidatePassword(pass string) error {
	return ValidateString(pass, 6, 100)
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 100); err != nil {
		return err
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
