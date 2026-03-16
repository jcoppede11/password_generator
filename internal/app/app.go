package app

import (
	"flag"
	"fmt"
	"os"

	generator "github.com/jcoppede11/password_generator"
)

// Run parses command-line flags, generates a password and
// writes the result to stdout. Returns an error if something fails.
func Run() error {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	length := flags.Int("length", 12, "Password length")
	useUpper := flags.Bool("uppercase", true, "Include uppercase letters")
	useLower := flags.Bool("lowercase", true, "Include lowercase letters")
	useNumbers := flags.Bool("numbers", true, "Include numbers")
	useSymbols := flags.Bool("symbols", true, "Include symbols")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	password, err := generator.Generate(generator.Options{
		Length:     *length,
		UseUpper:   *useUpper,
		UseLower:   *useLower,
		UseNumbers: *useNumbers,
		UseSymbols: *useSymbols,
	})
	if err != nil {
		return err
	}

	fmt.Println("Generated password:", password)
	fmt.Println("Strength:          ", generator.StrengthScore(password))

	return nil
}
