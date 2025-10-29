package main

import (
	"log"

	app "github.com/kasaderos/notification/internal/app/percolator"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run event percolator: %v", err)
	}
}
