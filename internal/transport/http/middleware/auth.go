package middleware

import (
	"context"
	"net/http"
	"strings"
)

const ctxKeyToken = "token"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token2 := strings.TrimPrefix(h, "Bearer ")
		ctx := context.WithValue(r.Context(), ctxKeyToken, token2)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TokenFromContext(ctx context.Context) (string, bool) {
	t, ok := ctx.Value(ctxKeyToken).(string)
	return t, ok
}
