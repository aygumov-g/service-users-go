package me

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	users    UserService
	identity IdentityService
}

func NewHandler(users UserService, identity IdentityService) *Handler {
	return &Handler{
		users:    users,
		identity: identity,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	identity, ok := h.identity.Unload(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.users.GetOrCreateMe(r.Context(), identity)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
