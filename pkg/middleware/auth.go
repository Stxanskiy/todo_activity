package middleware

import (
	"context"
	"net/http"
	"strings"
	"todo_activity/pkg/auth"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Добавляем UserID в контекст запроса
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		// Продолжаем выполнение с новым контекстом
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
