package auth

import (
	"context"

	"github.com/aygumov-g/service-users-go/internal/domain/identity"
)

type SSOClient interface {
	Me(ctx context.Context, token string) (*identity.Identity, error)
}

type IdentityService interface {
	Upload(ctx context.Context, idntt identity.Identity) context.Context
}
