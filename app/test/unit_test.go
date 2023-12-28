package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/znxxx/Pitchaya_agnos_backend/services"
)

func TestStrongPasswordSteps(t *testing.T) {
	passwordService := services.NewPasswordService()

	password1 := "aA1bB2"
	steps1 := passwordService.StrongPasswordSteps(password1)
	assert.Equal(t, 0, steps1)

	password2 := "aA1"
	steps2 := passwordService.StrongPasswordSteps(password2)
	assert.Equal(t, 3, steps2)

	password3 := "aa1"
	steps3 := passwordService.StrongPasswordSteps(password3)
	assert.Equal(t, 4, steps3)

	shortPassword := "aA1"
	steps4 := passwordService.StrongPasswordSteps(shortPassword)
	assert.Equal(t, 3, steps4)

	longPassword := "aAbB1cCdD2eEfF3gGhH4iI5jJ6"
	steps5 := passwordService.StrongPasswordSteps(longPassword)
	assert.Equal(t, 6, steps5)

	noUppercase := "aa1bb2"
	steps6 := passwordService.StrongPasswordSteps(noUppercase)
	assert.Equal(t, 1, steps6)

	noLowercase := "AA1BB2"
	steps7 := passwordService.StrongPasswordSteps(noLowercase)
	assert.Equal(t, 1, steps7)

	noDigit := "aAbBcCdD"
	steps8 := passwordService.StrongPasswordSteps(noDigit)
	assert.Equal(t, 1, steps8)

	repeatingChars := "11223!!!..."
	steps9 := passwordService.StrongPasswordSteps(repeatingChars)
	assert.Equal(t, 4, steps9)

	exceed20Char := "11AABBCCaaddaabb112233!.!.!."
	steps10 := passwordService.StrongPasswordSteps(exceed20Char)
	assert.Equal(t, 8, steps10)
}
