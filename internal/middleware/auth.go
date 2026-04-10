package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kassshi/golang-practice/internal/auth"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizetionHeader := r.Header.Get("Authorization")
			if authorizetionHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			jwt := strings.TrimPrefix(authorizetionHeader, "Bearer ")
			userID, err := auth.ParseToken(jwt, jwtSecret)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userIDKey, userID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
