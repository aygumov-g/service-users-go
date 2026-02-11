package me

import (
	"context"
	"os/user"

	"github.com/aygumov-g/service-users-go/internal/domain/identity"
)

type UserService interface {
	GetOrCreateMe(ctx context.Context, idntt identity.Identity) (*user.User, error)
}

type IdentityService interface {
	Unload(ctx context.Context) (identity.Identity, bool)
}
