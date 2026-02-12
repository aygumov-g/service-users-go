package sso

import (
	"context"

	d_identity "github.com/aygumov-g/service-users-go/internal/domain/identity"
)

type Client interface {
	Me(ctx context.Context, token string) (*d_identity.Identity, error)
}
