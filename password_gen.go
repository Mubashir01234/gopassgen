package gopassgen

import (
	"math/rand"
	"time"
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

type passgen struct {
	opt Options
}

// New creates a new Options for generate passwords.
func New(opt Options) *passgen {
	return &passgen{opt}
}

// GeneratePassword generates a new password from custom options.
func (o *passgen) GeneratePassword() (string, error) {
	if o.opt.Length <= 0 {
		return "", errInvalidLength
	}

	if !o.opt.HasLower && !o.opt.HasUpper && !o.opt.HasSymbols && !o.opt.HasNumbers {
		return "", errMissingOption
	}

	return o.createPassword(), nil
}

// createPassword creates a string according to options.
func (o *passgen) createPassword() string {
	var chars string
	if o.opt.HasLower {
		chars += lowerCharSet
	}
	if o.opt.HasUpper {
		chars += upperCharSet
	}
	if o.opt.HasNumbers {
		chars += numberSet
	}
	if o.opt.HasSymbols {
		chars += specialCharSet
	}
	return o.generatePassword(chars)
}

// generatePassword return generated password.
func (o *passgen) generatePassword(chars string) string {
	charRunes := []rune(chars)
	passwordRunes := make([]rune, o.opt.Length)
	for i := range passwordRunes {
		passwordRunes[i] = charRunes[rand.Intn(len(charRunes))]
	}
	return string(passwordRunes)
}
