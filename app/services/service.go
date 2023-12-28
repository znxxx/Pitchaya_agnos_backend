package services

import (
	"unicode"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) StrongPasswordSteps(password string) int {
	numSteps := 0
	if len(password) < 6 {
		numSteps += 6 - len(password)
	} else if len(password) >= 20 {
		numSteps += len(password) - 20
	}

	var (
		hasLowercase   bool
		hasUppercase   bool
		hasDigit       bool
		repeatingCount int
	)

	for i, char := range password {
		if unicode.IsLower(char) {
			hasLowercase = true
		} else if unicode.IsUpper(char) {
			hasUppercase = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}

		if i > 1 && char == rune(password[i-1]) && char == rune(password[i-2]) {
			repeatingCount++
		}
	}

	if !hasLowercase {
		numSteps++
	}
	if !hasUppercase {
		numSteps++
	}
	if !hasDigit {
		numSteps++
	}
	if repeatingCount > 0 {
		numSteps += repeatingCount
	}

	return numSteps
}
