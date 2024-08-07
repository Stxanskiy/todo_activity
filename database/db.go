package db

import (
	"context"
	"fmt"
	"log"
	"todo_activity/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	cfg := config.LoadConfig()

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err = DB.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}

	log.Println("Connected to database!")
}

func Close() {
	DB.Close()
}
