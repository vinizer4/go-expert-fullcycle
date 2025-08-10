package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/database/mongodb"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
