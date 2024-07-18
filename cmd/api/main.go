package main

import (
	"context"
	"fmt"
	"github.com/ianbrito/fr-cotacao/internal/infra/api"
	"github.com/ianbrito/fr-cotacao/internal/infra/api/handler"
	"github.com/ianbrito/fr-cotacao/internal/infra/repository"
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
	ctx := context.Background()

	repository.InitDB()
	defer func() {
		err := repository.CloseDB()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("API de Cotações")

	port := os.Getenv("HTTP_PORT")

	server := api.NewWebServer(port)

	quoteHandler := handler.NewQuoteHandler(ctx)
	server.AddHandler("/api/v1/quote", quoteHandler.GetQuote)

	metricsHandler := handler.NewMetricsHandler(ctx)
	server.AddHandler("/api/v1/metrics", metricsHandler.GetMetrics)

	server.Run()
}
