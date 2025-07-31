// Package validator
package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/shuvo-paul/email-microservice/internal/models"
)

func IsValidEmail(email string) bool {
	if len(email) == 0 {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(strings.TrimSpace(email))
}

func SanitizeString(input string) string {
	sanitized := strings.ReplaceAll(input, "\x00", "")
	sanitized = strings.TrimSpace(sanitized)
	return sanitized
}

func ValidateEmailRequest(email models.EmailRequest) error {
	if !IsValidEmail(email.To) {
		return errors.New("not a valid email address")
	}

	if SanitizeString(email.Body) == "" {
		return errors.New("body is empty")
	}

	if SanitizeString(email.Subject) == "" {
		return errors.New("subject is empty")
	}
	return nil
}
