package auth

import (
	"net/http"
	"strings"
)

type Middleware struct {
	sso      SSOClient
	identity IdentityService
}

func NewMiddleware(sso SSOClient, identity IdentityService) *Middleware {
	return &Middleware{
		sso:      sso,
		identity: identity,
	}
}

func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(h, "Bearer ")

		idntt, err := m.sso.Me(r.Context(), token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := m.identity.Upload(r.Context(), *idntt)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
