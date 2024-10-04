package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Create a new instance of our application and configuration struct.
	config := config{
		addr: os.Getenv("PORT"),
	}
	app := &application{
		config: config,
	}

	// Call the run() method on our application struct.
	if err := app.run(); err != nil {
		log.Fatal(err)
	}

}
