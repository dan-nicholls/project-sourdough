package web

import (
	"context"
	"net/http"
	"strings"
	"log"
	"time"

	"github.com/dan-nicholls/project-sourdough/internal/app"
)


type TokenAuthKey struct {}

func GetTokenFromContext(ctx context.Context) (app.Token, bool) {
	t, ok := ctx.Value(TokenAuthKey{}).(app.Token)
	return t, ok
}

func TokenAuthMiddleware(ts app.TokenStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			const prefix = "Bearer "

			if !strings.HasPrefix(authHeader, prefix) {
				log.Printf("Incorrect formatting of Bearer token")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenID := strings.TrimPrefix(authHeader,prefix)
			token, err := ts.GetToken(tokenID)
			if err != nil {
				log.Printf("Invalid tokenID: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if token.Used || token.ExpiresAt.Before(time.Now()) || token.IssuedAt.After(time.Now()) {
				log.Printf("Invalid token: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), TokenAuthKey{}, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
