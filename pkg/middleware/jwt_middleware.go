package middleware

import (
	"context"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/auth"
	"net/http"
	"strings"
)

type contextKey string

const (
	ContextUserKey contextKey = "user"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorizatoion header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
