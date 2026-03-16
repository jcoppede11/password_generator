package generator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"unicode"
)

// Options defines the parameters for generating a password.
type Options struct {
	Length     int
	UseUpper   bool
	UseLower   bool
	UseNumbers bool
	UseSymbols bool
}

// Generate creates a cryptographically secure random password based on
// the provided options. Returns an error if the configuration is invalid.
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
		return "", errors.New("length must be greater than 0")
	}
	if opts.Length > MaxPasswordLength {
		return "", errors.New("maximum allowed length is 128")
	}
	if charset == "" {
		return "", errors.New("select at least one character type")
	}

	password := make([]byte, opts.Length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error generating randomness: %w", err)
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}

// StrengthScore evaluates the strength of a password and returns one of the
// following categories: "Weak", "Medium", "Strong" or "Very strong".
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
		return "Very strong"
	case score >= 5:
		return "Strong"
	case score >= 3:
		return "Medium"
	default:
		return "Weak"
	}
}
