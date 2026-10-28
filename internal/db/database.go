package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gatimugabriel/1m-transactions/pkg/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	defer utils.FunctionTimer("db:Connect")()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=public&sslmode=disable", user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(DB)

	if err := DB.PingContext(context.Background()); err != nil {
		log.Fatal("database unreachable", err)
	}

	fmt.Println("connected to database")
}
