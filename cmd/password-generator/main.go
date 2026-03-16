package main

import (
	"log"

	"github.com/jcoppede11/password_generator/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
