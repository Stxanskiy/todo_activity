package handler

import (
	"context"
	"encoding/json"
	"net/http"
	db "todo_activity/database"
	"todo_activity/internal/activity/model"

	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.CreateDate = time.Now()

	_, err := db.DB.Exec(context.Background(), "INSERT INTO \"user\" (username, email, password_hash, create_date) VALUES ($1, $2, $3, $4)",
		user.Username, user.Email, user.PasswordHash, user.CreateDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
