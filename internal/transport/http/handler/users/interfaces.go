package users

import (
	"context"

	d_identity "github.com/aygumov-g/service-users-go/internal/domain/identity"
	d_user "github.com/aygumov-g/service-users-go/internal/domain/user"
)

type UserService interface {
	GetByIDWithAccess(ctx context.Context, identity *d_identity.Identity, targetID int64) (*d_user.User, error)
}

type IdentityHTTP interface {
	Unload(ctx context.Context) (*d_identity.Identity, bool)
}
