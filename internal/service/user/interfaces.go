package user

import (
	"context"
	"time"

	"github.com/aygumov-g/service-users-go/internal/domain/user"
	"github.com/aygumov-g/service-users-go/internal/integration/sso"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
}

type SSOClient interface {
	Me(ctx context.Context, token string) (*sso.Identify, error)
}

type Clock interface {
	Now() time.Time
}
