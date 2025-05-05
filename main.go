package main

import (
	"flag"
	"fmt"
	"log"
	"main/generator"	
)

func main() {
	length := flag.Int("length", 12, "Longitud de la contraseña")
	useUpper := flag.Bool("uppercase", true, "Incluir mayúsculas")
	useLower := flag.Bool("lowercase", true, "Incluir minúsculas")
	useNumbers := flag.Bool("numbers", true, "Incluir números")
	useSymbols := flag.Bool("symbols", true, "Incluir símbolos")

	flag.Parse()

	password, err := generator.PasswordGenerator(*length, *useUpper, *useLower, *useNumbers, *useSymbols)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contraseña generada:", password)
}
