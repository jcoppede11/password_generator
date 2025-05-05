package generator

import (
	"errors"
	"math/rand"
)

func PasswordGenerator(length int, useUpper, useLower, useNumbers, useSymbols bool) (string, error) {
	const MaxPasswordLength = 128
	var charset string

	if useUpper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useLower {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if useNumbers {
		charset += "0123456789"
	}
	if useSymbols {
		charset += "!@#$%^&*()-_=+[]{}<>?/|"
	}
	if length <= 0 {
		return "", errors.New("la longitud debe ser mayor a 0")
	}
	if length > MaxPasswordLength {
		return "", errors.New("la longitud máxima permitida es 128")
	}	
	if charset == "" {
		return "", errors.New("seleccioná al menos un tipo de caracter")
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password), nil
}
