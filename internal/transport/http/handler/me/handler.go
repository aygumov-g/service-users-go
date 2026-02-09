package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aygumov-g/service-users-go/internal/service/user"
	"github.com/aygumov-g/service-users-go/internal/transport/http/middleware"
)

type Handler struct {
	users user.Service
}

func NewHandler(users user.Service) *Handler {
	return &Handler{users: users}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, ok := middleware.TokenFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	u, err := h.users.GetOrCreateMe(r.Context(), token)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(u)
}
