package gopassgen

import (
	"math/rand"
	"time"
)

const (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%^&*()-_=+|[]{};:/?.<>"
	numberSet      = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Options struct {
	Length     int
	HasLower   bool
	HasUpper   bool
	HasSymbols bool
	HasNumbers bool
}

func GeneratePassword(opts Options) (string, error) {
	if opts.Length <= 0 {
		return "", errInvalidLength
	}

	if !opts.HasLower && !opts.HasUpper && !opts.HasSymbols && !opts.HasNumbers {
		return "", errMissingOption
	}

	return createPassword(opts), nil
}

func createPassword(opts Options) string {
	var chars string
	if opts.HasLower {
		chars += lowerCharSet
	}
	if opts.HasUpper {
		chars += upperCharSet
	}
	if opts.HasNumbers {
		chars += numberSet
	}
	if opts.HasSymbols {
		chars += specialCharSet
	}
	return generatePassword(opts.Length, chars)
}

func generatePassword(length int, chars string) string {
	charRunes := []rune(chars)
	passwordRunes := make([]rune, length)
	for i := range passwordRunes {
		passwordRunes[i] = charRunes[rand.Intn(len(charRunes))]
	}
	return string(passwordRunes)
}
