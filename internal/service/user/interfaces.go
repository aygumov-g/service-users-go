package user

import (
	"context"
	"time"

	"github.com/aygumov-g/service-users-go/internal/domain/identity"
	"github.com/aygumov-g/service-users-go/internal/domain/user"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
}

type SSOClient interface {
	Me(ctx context.Context, token string) (*identity.Identity, error)
}

type Clock interface {
	Now() time.Time
}
