package main

import (
	"fmt"
	"github.com/ianbrito/fr-cotacao/internal/infra/api"
	"github.com/ianbrito/fr-cotacao/internal/infra/api/handler"
	"github.com/ianbrito/fr-cotacao/internal/infra/db"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {

	env := os.Getenv("APP_ENV")

	if env != "production" {
		loadDotEnv()
	}

	conn := db.GetConnection()
	defer conn.Close()

	fmt.Println("API de Cotações")

	var port = os.Getenv("HTTP_PORT")

	server := api.NewWebServer(port)

	server.AddHandler("/api/v1/quote", handler.GetQuote)

	server.Run()
}
