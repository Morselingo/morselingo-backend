package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authentication header format. Expected 'Bearer <token>'", http.StatusUnauthorized)
		}
		tokenString := parts[1]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Token is malformed", http.StatusUnauthorized)
			} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				http.Error(w, "Token is expired or not yet valid", http.StatusUnauthorized)
			} else {
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			}
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyUser, claims.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetUsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(ContextKeyUser).(string)
	return username, ok
}
