package app

import (
	"flag"
	"fmt"
	"os"

	generator "github.com/jcoppede11/password_generator"
)

// Run parsea los flags de línea de comandos, genera la contraseña y
// escribe el resultado en stdout. Devuelve un error si algo falla.
func Run() error {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	length := flags.Int("length", 12, "Longitud de la contraseña")
	useUpper := flags.Bool("uppercase", true, "Incluir mayúsculas")
	useLower := flags.Bool("lowercase", true, "Incluir minúsculas")
	useNumbers := flags.Bool("numbers", true, "Incluir números")
	useSymbols := flags.Bool("symbols", true, "Incluir símbolos")

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

	fmt.Println("Contraseña generada:", password)
	fmt.Println("Fortaleza:          ", generator.StrengthScore(password))

	return nil
}
