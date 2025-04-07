package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Token not found", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(auth, "Bearer ")[1]
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
