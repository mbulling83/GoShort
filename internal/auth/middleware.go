package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userKey contextKey = "user"

// JWTMiddleware validates JWT tokens in incoming requests
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user information to the context
		ctx := context.WithValue(r.Context(), userKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUsernameFromContext retrieves the username from the request context
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(userKey).(string)
	return username, ok
}
