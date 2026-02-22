package main

import (
	"log"

	"fiber-example/api"
	"fiber-example/api/example"
	"fiber-example/database"
	"fiber-example/server"

	"github.com/joho/godotenv"
)

func main() {
	// not use
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Server initialization
	app := server.Create()

	// Migrations
	database.DB.AutoMigrate(&example.Example{})

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
