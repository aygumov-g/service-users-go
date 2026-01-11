package user

import (
	"context"

	"github.com/aygumov-g/service-users-go/internal/domain/user"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*user.User, error)
	Create(ctx context.Context, u *user.User) error
	Update(ctx context.Context, u *user.User) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}
