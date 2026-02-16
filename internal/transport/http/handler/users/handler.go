package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/aygumov-g/service-users-go/internal/service/user"
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
	}
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	identity, ok := h.identity.Unload(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	u, err := h.users.GetByIDWithAccess(r.Context(), identity, id)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrForbidden):
			http.Error(w, "forbidden", http.StatusForbidden)
		case errors.Is(err, user.ErrUserNotFound):
			http.Error(w, "user not found", http.StatusForbidden)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}

		return
	}

	var resp userResponse
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp.toResponse(u))
}
