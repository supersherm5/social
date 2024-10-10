package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/supersherm5/social/internal/db"
	"github.com/supersherm5/social/internal/storage"
	"github.com/supersherm5/social/internal/utils"
)

func main() {
	// Create a new instance of our application and configuration struct.
	config := config{
		addr: os.Getenv("PORT"),
		db: dbConfig{
			addr:         utils.GetStringEnv("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: utils.GetIntEnv("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: utils.GetIntEnv("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  utils.GetStringEnv("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// Create a new instance of our PostgreSQL database connection.
	pgdb, err := db.NewPG(config.db.addr, config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer pgdb.Close()
	log.Println("Successfully connected to PostgreSQL database")

	// Create a new instance of our storage struct.
	store := storage.NewStorage(pgdb)

	// Create a new instance of our application struct.
	app := &application{
		config: config,
		store:  store,
	}

	// Call the run() method on our application struct.
	if err := app.run(); err != nil {
		log.Fatal(err)
	}

}
