package main

import (
	"log"

	"github.com/gatimugabriel/1m-transactions/internal/db"
	"github.com/gatimugabriel/1m-transactions/internal/transactions"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db.Connect()
	transactions.ProcessTransactions()
}
