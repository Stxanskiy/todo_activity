package user

import (
	"context"
	"encoding/json"
	"net/http"
	db "todo_activity/database"
	"todo_activity/internal/activity/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), "SELECT user_id, username, email, password_hash, create_date FROM person")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.PasswordHash, &user.CreateDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Define other handlers for create, update, delete user
