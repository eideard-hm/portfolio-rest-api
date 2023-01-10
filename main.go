package main

import (
	"fmt"
	"os"

	"github.com/eideard-hm/portfolio-rest-api/src"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Add server configuration
	app := src.Server()

	app.Listen(fmt.Sprintf(":%s", port))
}
