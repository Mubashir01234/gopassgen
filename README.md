# GoPassGen
[![Godoc](https://godoc.org/github.com/Mubashir01234/gopassgen?status.svg)](https://godoc.org/github.com/Mubashir01234/gopassgen) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/Mubashir01234/gopassgen/main/LICENSE)

A simple Go package to generate random passwords based on user-defined criteria.

## Usage

```go
import "github.com/Mubashir01234/gopassgen"

opts := gopassgen.Options{
	Length:    12,
	HasLower:  true,
	HasUpper:  true,
	HasSymbols: true,
	HasNumbers: true,
}

password, err := gopassgen.GeneratePassword(opts)
if err != nil {
	fmt.Println("Error generating password:", err)
	return
}
fmt.Println(password)
