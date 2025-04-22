package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/pimentafm/auction-go/configuration/database/mongodb"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
