package main

import (
	"log"

	app "github.com/kasaderos/notification/internal/app/crawler"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run crawler: %v", err)
	}
}
