package generator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"unicode"
)

// Options define los parámetros para generar una contraseña.
type Options struct {
	Length     int
	UseUpper   bool
	UseLower   bool
	UseNumbers bool
	UseSymbols bool
}

// Generate crea una contraseña aleatoria criptográficamente segura según
// las opciones proporcionadas. Devuelve un error si la configuración no es válida.
func Generate(opts Options) (string, error) {
	const MaxPasswordLength = 128
	var charset string

	if opts.UseUpper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if opts.UseLower {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if opts.UseNumbers {
		charset += "0123456789"
	}
	if opts.UseSymbols {
		charset += "!@#$%^&*()-_=+[]{}<>?/|"
	}
	if opts.Length <= 0 {
		return "", errors.New("la longitud debe ser mayor a 0")
	}
	if opts.Length > MaxPasswordLength {
		return "", errors.New("la longitud máxima permitida es 128")
	}
	if charset == "" {
		return "", errors.New("seleccioná al menos un tipo de caracter")
	}

	password := make([]byte, opts.Length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error al generar aleatoriedad: %w", err)
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}

// StrengthScore evalúa la fortaleza de una contraseña y devuelve una de las
// siguientes categorías: "Débil", "Media", "Fuerte" o "Muy fuerte".
func StrengthScore(password string) string {
	score := 0

	switch l := len(password); {
	case l >= 16:
		score += 3
	case l >= 12:
		score += 2
	case l >= 8:
		score += 1
	}

	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSymbol {
		score++
	}

	switch {
	case score >= 6:
		return "Muy fuerte"
	case score >= 5:
		return "Fuerte"
	case score >= 3:
		return "Media"
	default:
		return "Débil"
	}
}
