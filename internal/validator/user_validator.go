package validator

import (
	"errors"
	"strings"
)

func ValidateUser(name, email string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(email, "@") {
		return errors.New("email is invalid")
	}
	return nil
}
