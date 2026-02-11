package user

import (
	"context"

	"github.com/aygumov-g/service-users-go/internal/domain/identity"
	"github.com/aygumov-g/service-users-go/internal/domain/user"
)

type Service struct {
	repo Repository
	sso  SSOClient
	clk  Clock
}

func NewService(repo Repository, sso SSOClient) *Service {
	return &Service{
		repo: repo,
		sso:  sso,
	}
}

func (s *Service) GetOrCreateMe(ctx context.Context, idntt identity.Identity) (*user.User, error) {
	u, err := s.repo.GetByID(ctx, idntt.ID)
	if err == nil {
		if u.Login != idntt.Login {
			u.Login = idntt.Login
			u.UpdatedAt = s.clk.Now()

			_ = s.repo.Update(ctx, u)
		}

		return u, nil
	}

	if err != ErrUserNotFound {
		return nil, err
	}

	now := s.clk.Now()
	u = &user.User{
		ID:        idntt.ID,
		Login:     idntt.Login,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
