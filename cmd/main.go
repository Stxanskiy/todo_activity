package main

import (
	"log"
	"net/http"
	db "todo_activity/database"
	routes "todo_activity/internal/activity/delivery/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db.Connect()
	defer db.Close()

	r := routes.SetupRouter()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
