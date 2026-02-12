package user

import (
	"context"

	d_user "github.com/aygumov-g/service-users-go/internal/domain/user"
)

type Repository interface {
	//Upsert(ctx context.Context, user *user.User) (*user.User, error)
	GetByID(ctx context.Context, id int64) (*d_user.User, error)
	Create(ctx context.Context, user *d_user.User) error
	Update(ctx context.Context, user *d_user.User) error
}
