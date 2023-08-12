# GoPassGen

A simple Go package to generate random passwords based on user-defined criteria.

## Usage

```go
import "github.com/yourusername/gopassgen"

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
