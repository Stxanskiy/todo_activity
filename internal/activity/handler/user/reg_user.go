package user

import (
	"context"
	"encoding/json"
	"net/http"
	db "todo_activity/database"
	"todo_activity/internal/activity/model"
	"todo_activity/pkg/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.PasswordHash = hashedPassword

	_, err = db.DB.Exec(context.Background(),
		"INSERT INTO person (username, email, password_hash, create_date) VALUES ($1, $2,$3, $4)",
		user.Username, user.Email, user.PasswordHash, user.CreateDate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
