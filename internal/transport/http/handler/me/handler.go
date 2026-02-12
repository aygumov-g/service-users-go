package me

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	users    UserService
	identity IdentityHTTP
}

func NewHandler(users UserService, identity IdentityHTTP) *Handler {
	return &Handler{
		users:    users,
		identity: identity,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPatch:
		h.update(w, r)
	}
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
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

	var resp UserResponse
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp.ToResponse(user))
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	identity, ok := h.identity.Unload(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user, err := h.users.UpdateMe(r.Context(), identity.ID, req.toServiceInput())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	var resp UserResponse
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp.ToResponse(user))
}
