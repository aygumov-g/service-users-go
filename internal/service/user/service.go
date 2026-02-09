package user

import (
	"context"

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

func (s *Service) GetOrCreateMe(ctx context.Context, token string) (*user.User, error) {
	identify, err := s.sso.Me(ctx, token)
	if err != nil {
		return nil, err
	}

	u, err := s.repo.GetByID(ctx, identify.ID)
	if err == nil {
		if u.Login != identify.Login {
			u.Login = identify.Login
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
		ID:        identify.ID,
		Login:     identify.Login,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
