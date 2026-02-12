package user

import (
	"context"
	"errors"

	d_identity "github.com/aygumov-g/service-users-go/internal/domain/identity"
	d_user "github.com/aygumov-g/service-users-go/internal/domain/user"
)

type Service struct {
	repo Repository
	clk  Clock
}

func NewService(repo Repository, clk Clock) *Service {
	return &Service{
		repo: repo,
		clk:  clk,
	}
}

func (s *Service) GetOrCreateMe(ctx context.Context, identity *d_identity.Identity) (*d_user.User, error) {
	u, err := s.repo.GetByID(ctx, identity.ID)
	if err == nil {
		return u, nil
	}

	if !errors.Is(err, ErrUserNotFound) {
		return nil, err
	}

	now := s.clk.Now()

	u = &d_user.User{
		ID:        identity.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.repo.Create(ctx, u)
	if errors.Is(err, ErrUserAlreadyExists) {
		return s.repo.GetByID(ctx, u.ID)
	}

	return u, nil
}

func (s *Service) UpdateMe(ctx context.Context, id int64, input UpdateInput) (*d_user.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.FirstName != nil {
		u.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		u.LastName = *input.LastName
	}

	if input.Bio != nil {
		u.Bio = *input.Bio
	}

	if input.AvatarURL != nil {
		u.AvatarURL = *input.AvatarURL
	}

	u.UpdatedAt = s.clk.Now()

	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) GetByIDWithAccess(ctx context.Context, identity *d_identity.Identity, targetID int64) (*d_user.User, error) {
	if identity.Role != "admin" && identity.ID != targetID {
		return nil, ErrForbidden
	}

	u, err := s.repo.GetByID(ctx, targetID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
